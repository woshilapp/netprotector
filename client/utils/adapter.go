package utils

import (
	"errors"
	"unsafe"
)

const (
	MAX_ADAPTER_DESCRIPTION_LENGTH = 128
	MAX_ADAPTER_NAME_LENGTH        = 256
	MAX_ADAPTER_ADDRESS_LENGTH     = 8
	IF_TYPE_ETHERNET_CSMACD        = 6
	IF_TYPE_IEEE80211              = 71
)

type IP_ADAPTER_INFO struct {
	Next                *IP_ADAPTER_INFO
	ComboIndex          uint32
	AdapterName         [MAX_ADAPTER_NAME_LENGTH + 4]byte
	Description         [MAX_ADAPTER_DESCRIPTION_LENGTH + 4]byte
	AddressLength       uint32
	Address             [MAX_ADAPTER_ADDRESS_LENGTH]byte
	Index               uint32
	Type                uint32
	DhcpEnabled         uint32
	CurrentIpAddress    *byte
	IpAddressList       IP_ADDR_STRING
	GatewayList         IP_ADDR_STRING
	DhcpServer          IP_ADDR_STRING
	HaveWins            bool
	PrimaryWinsServer   IP_ADDR_STRING
	SecondaryWinsServer IP_ADDR_STRING
	LeaseObtained       int64
	LeaseExpires        int64
}

type IP_ADDR_STRING struct {
	Next      *IP_ADDR_STRING
	IpAddress [16]byte
	IpMask    [16]byte
	Context   uint32
}

type MyAdapter struct {
	Name    string
	Type    string
	Gateway string
}

var (
	// iphlpapi            = syscall.NewLazyDLL("iphlpapi.dll")
	procGetAdaptersInfo = iphlpapi.NewProc("GetAdaptersInfo")
)

func byteSliceToString(b []byte) string {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

func GetNetworkInfo() ([]MyAdapter, error) {
	var adapterInfo *IP_ADAPTER_INFO
	var size uint32 = 0

	// 第一次调用获取所需缓冲区大小
	procGetAdaptersInfo.Call(uintptr(unsafe.Pointer(adapterInfo)), uintptr(unsafe.Pointer(&size)))

	buffer := make([]byte, size)
	adapterInfo = (*IP_ADAPTER_INFO)(unsafe.Pointer(&buffer[0]))

	// 第二次调用获取实际数据
	ret, _, _ := procGetAdaptersInfo.Call(
		uintptr(unsafe.Pointer(adapterInfo)),
		uintptr(unsafe.Pointer(&size)),
	)
	if ret != 0 {
		return nil, errors.New("GetAdaptersInfo failed")
	}

	apalist := []MyAdapter{}

	// 遍历所有网卡
	for info := adapterInfo; info != nil; info = info.Next {
		adapterDesc := byteSliceToString(info.Description[:])
		adapterType := "Wired"
		if info.Type == IF_TYPE_IEEE80211 {
			adapterType = "Wireless"
		}

		gateway := byteSliceToString(info.GatewayList.IpAddress[:])
		if gateway == "" {
			gateway = "None"
		}

		apalist = append(apalist, MyAdapter{
			Name:    adapterDesc,
			Type:    adapterType,
			Gateway: gateway,
		})

		// fmt.Printf("网卡名称: %s\n", adapterDesc)
		// fmt.Printf("网卡类型: %s\n", adapterType)
		// fmt.Printf("默认网关: %s\n", gateway)
		// fmt.Println("-----------------------------")
	}

	return apalist, nil
}
