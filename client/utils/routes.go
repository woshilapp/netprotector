package utils

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

const (
	ANY_SIZE = 1
)

type MIB_IPFORWARDROW struct {
	dwForwardDest      uint32
	dwForwardMask      uint32
	dwForwardPolicy    uint32
	dwForwardNextHop   uint32
	dwForwardIfIndex   uint32
	dwForwardType      uint32
	dwForwardProto     uint32
	dwForwardAge       uint32
	dwForwardNextHopAS uint32
	dwForwardMetric1   uint32
	dwForwardMetric2   uint32
	dwForwardMetric3   uint32
	dwForwardMetric4   uint32
	dwForwardMetric5   uint32
}

type MIB_IPFORWARDTABLE struct {
	dwNumEntries uint32
	table        [ANY_SIZE]MIB_IPFORWARDROW
}

var (
	iphlpapi             = syscall.NewLazyDLL("iphlpapi.dll")
	createIpForwardEntry = iphlpapi.NewProc("CreateIpForwardEntry")
	deleteIpForwardEntry = iphlpapi.NewProc("DeleteIpForwardEntry")
	getIpForwardTable    = iphlpapi.NewProc("GetIpForwardTable")
)

func ListRoutes() {
	var pIpForwardTable *MIB_IPFORWARDTABLE
	var dwSize uint32
	var order uint32 = 0

	_, _, _ = getIpForwardTable.Call(
		uintptr(unsafe.Pointer(pIpForwardTable)),
		uintptr(unsafe.Pointer(&dwSize)),
		uintptr(order),
	)

	buffer := make([]byte, dwSize)
	pIpForwardTable = (*MIB_IPFORWARDTABLE)(unsafe.Pointer(&buffer[0]))

	ret, _, _ := getIpForwardTable.Call(
		uintptr(unsafe.Pointer(pIpForwardTable)),
		uintptr(unsafe.Pointer(&dwSize)),
		uintptr(order),
	)

	if ret != 0 {
		fmt.Println("获取路由表失败，错误代码:", ret)
		return
	}

	fmt.Println("目标网络\t\t子网掩码\t\t下一跳\t\t接口索引\t\t跃点数")
	fmt.Println("----------------------------------------------------------------")

	numEntries := int(pIpForwardTable.dwNumEntries)
	rows := (*[1 << 20]MIB_IPFORWARDROW)(unsafe.Pointer(&pIpForwardTable.table))[:numEntries:numEntries]

	for _, row := range rows {
		fmt.Printf("%-15s\t%-15s\t%-15s\t%-10d\t%-8d\n",
			formatIP(row.dwForwardDest),
			formatIP(row.dwForwardMask),
			formatIP(row.dwForwardNextHop),
			row.dwForwardIfIndex,
			row.dwForwardMetric1,
		)
	}
}

func AddRoute(dest, mask, gateway, ifIndex, metric string) {
	row := MIB_IPFORWARDROW{
		dwForwardDest:    ipToUint32(dest),
		dwForwardMask:    ipToUint32(mask),
		dwForwardNextHop: ipToUint32(gateway),
		dwForwardIfIndex: uint32(atoi(ifIndex)),
		dwForwardType:    4, // 最佳路由
		dwForwardProto:   3, // 静态路由
		dwForwardMetric1: uint32(atoi(metric)),
	}

	ret, _, _ := createIpForwardEntry.Call(uintptr(unsafe.Pointer(&row)))
	if ret != 0 {
		fmt.Printf("添加路由失败，错误代码: %d\n", ret)
	} else {
		fmt.Println("路由添加成功")
	}
}

func DeleteRoute(dest string) {
	var pIpForwardTable *MIB_IPFORWARDTABLE
	var dwSize uint32
	var order uint32 = 0

	_, _, _ = getIpForwardTable.Call(
		uintptr(unsafe.Pointer(pIpForwardTable)),
		uintptr(unsafe.Pointer(&dwSize)),
		uintptr(order),
	)

	buffer := make([]byte, dwSize)
	pIpForwardTable = (*MIB_IPFORWARDTABLE)(unsafe.Pointer(&buffer[0]))

	ret, _, _ := getIpForwardTable.Call(
		uintptr(unsafe.Pointer(pIpForwardTable)),
		uintptr(unsafe.Pointer(&dwSize)),
		uintptr(order),
	)

	if ret != 0 {
		fmt.Println("获取路由表失败，错误代码:", ret)
		return
	}

	numEntries := int(pIpForwardTable.dwNumEntries)
	rows := (*[1 << 20]MIB_IPFORWARDROW)(unsafe.Pointer(&pIpForwardTable.table))[:numEntries:numEntries]

	targetIP := ipToUint32(dest)
	for _, row := range rows {
		if row.dwForwardDest == targetIP {
			ret, _, _ := deleteIpForwardEntry.Call(uintptr(unsafe.Pointer(&row)))
			if ret != 0 {
				fmt.Printf("删除路由失败，错误代码: %d\n", ret)
			} else {
				fmt.Println("路由删除成功")
			}
			return
		}
	}

	fmt.Println("未找到匹配的路由")
}

func ipToUint32(ipStr string) uint32 {
	ip := net.ParseIP(ipStr).To4()
	return uint32(ip[0]) | uint32(ip[1])<<8 | uint32(ip[2])<<16 | uint32(ip[3])<<24
}

func formatIP(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip),
		byte(ip>>8),
		byte(ip>>16),
		byte(ip>>24))
}

func atoi(s string) int {
	i := 0
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}
