package utils

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procOpenSCManager       = modadvapi32.NewProc("OpenSCManagerW")
	procOpenService         = modadvapi32.NewProc("OpenServiceW")
	procCreateService       = modadvapi32.NewProc("CreateServiceW")
	procDeleteService       = modadvapi32.NewProc("DeleteService")
	procStartService        = modadvapi32.NewProc("StartServiceW")
	procControlService      = modadvapi32.NewProc("ControlService")
	procQueryServiceStatus  = modadvapi32.NewProc("QueryServiceStatus")
	procCloseServiceHandle  = modadvapi32.NewProc("CloseServiceHandle")
	procChangeServiceConfig = modadvapi32.NewProc("ChangeServiceConfigW")
)

type ServiceManager struct {
	handle syscall.Handle
}

type Service struct {
	handle syscall.Handle
}

func NewServiceManager() (*ServiceManager, error) {
	handle, _, err := procOpenSCManager.Call(
		0,
		0,
		uintptr(0xF003F), // SC_MANAGER_ALL_ACCESS
	)
	if handle == 0 {
		return nil, fmt.Errorf("OpenSCManager failed: %v", err)
	}
	return &ServiceManager{handle: syscall.Handle(handle)}, nil
}

func (sm *ServiceManager) Close() error {
	if sm.handle != 0 {
		_, _, err := procCloseServiceHandle.Call(uintptr(sm.handle))
		sm.handle = 0
		if err != syscall.Errno(0) {
			return fmt.Errorf("CloseServiceHandle failed: %v", err)
		}
	}
	return nil
}

func (sm *ServiceManager) OpenService(name string, access uint32) (*Service, error) {
	ptr, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	handle, _, err := procOpenService.Call(
		uintptr(sm.handle),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(access),
	)
	if handle == 0 {
		return nil, fmt.Errorf("OpenService failed: %v", err)
	}
	return &Service{handle: syscall.Handle(handle)}, nil
}

func (sm *ServiceManager) CreateService(name, displayName, binPath string) (*Service, error) {
	n, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	d, err := syscall.UTF16PtrFromString(displayName)
	if err != nil {
		return nil, err
	}
	b, err := syscall.UTF16PtrFromString(binPath)
	if err != nil {
		return nil, err
	}

	handle, _, err := procCreateService.Call(
		uintptr(sm.handle),
		uintptr(unsafe.Pointer(n)),
		uintptr(unsafe.Pointer(d)),
		uintptr(0xF01FF), // SERVICE_ALL_ACCESS
		uintptr(0x10),    // SERVICE_WIN32_OWN_PROCESS
		uintptr(2),       // SERVICE_AUTO_START
		uintptr(1),       // SERVICE_ERROR_NORMAL
		uintptr(unsafe.Pointer(b)),
		0,
		0,
		0,
		0,
		0,
	)
	if handle == 0 {
		return nil, fmt.Errorf("CreateService failed: %v", err)
	}
	return &Service{handle: syscall.Handle(handle)}, nil
}

func (s *Service) Delete() error {
	ret, _, err := procDeleteService.Call(uintptr(s.handle))
	if ret == 0 {
		return fmt.Errorf("DeleteService failed: %v", err)
	}
	return nil
}

func (s *Service) Start() error {
	ret, _, err := procStartService.Call(
		uintptr(s.handle),
		0,
		0,
	)
	if ret == 0 {
		return fmt.Errorf("StartService failed: %v", err)
	}
	return nil
}

func (s *Service) Stop() error {
	var status SERVICE_STATUS
	ret, _, err := procControlService.Call(
		uintptr(s.handle),
		uintptr(1), // SERVICE_CONTROL_STOP
		uintptr(unsafe.Pointer(&status)),
	)
	if ret == 0 {
		return fmt.Errorf("ControlService failed: %v", err)
	}
	return nil
}

func (s *Service) Restart() error {
	if err := s.Stop(); err != nil {
		return err
	}
	return s.Start()
}

func (s *Service) QueryStatus() (string, error) {
	var status SERVICE_STATUS
	ret, _, err := procQueryServiceStatus.Call(
		uintptr(s.handle),
		uintptr(unsafe.Pointer(&status)),
	)
	if ret == 0 {
		return "", fmt.Errorf("QueryServiceStatus failed: %v", err)
	}

	switch status.CurrentState {
	case 1:
		return "STOPPED", nil
	case 2:
		return "START_PENDING", nil
	case 3:
		return "STOP_PENDING", nil
	case 4:
		return "RUNNING", nil
	default:
		return "UNKNOWN", nil
	}
}

func (s *Service) Close() error {
	if s.handle != 0 {
		_, _, err := procCloseServiceHandle.Call(uintptr(s.handle))
		s.handle = 0
		if err != syscall.Errno(0) {
			return fmt.Errorf("CloseServiceHandle failed: %v", err)
		}
	}
	return nil
}

type SERVICE_STATUS struct {
	ServiceType             uint32
	CurrentState            uint32
	ControlsAccepted        uint32
	Win32ExitCode           uint32
	ServiceSpecificExitCode uint32
	CheckPoint              uint32
	WaitHint                uint32
}

func main() {
	sm, err := NewServiceManager()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sm.Close()

	// 示例用法
	service, err := sm.CreateService(
		"MyTestService",
		"My Test Service",
		`C:\path\to\your\service.exe`,
	)
	if err != nil {
		fmt.Println("CreateService error:", err)
		return
	}
	defer service.Close()

	fmt.Println("Service created successfully")

	if err := service.Start(); err != nil {
		fmt.Println("Start error:", err)
		return
	}
	fmt.Println("Service started successfully")

	status, err := service.QueryStatus()
	if err != nil {
		fmt.Println("QueryStatus error:", err)
		return
	}
	fmt.Println("Service status:", status)

	if err := service.Stop(); err != nil {
		fmt.Println("Stop error:", err)
		return
	}
	fmt.Println("Service stopped successfully")

	if err := service.Delete(); err != nil {
		fmt.Println("Delete error:", err)
		return
	}
	fmt.Println("Service deleted successfully")
}
