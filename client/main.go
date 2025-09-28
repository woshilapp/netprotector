package main

import (
	"fmt"
	"os"

	"github.com/woshilapp/netprotector/client/rule"
	"github.com/woshilapp/netprotector/client/utils"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	ssid, err := utils.GetSSID()
	if err != nil {
		fmt.Println("SSID:", err.Error())
	} else {
		fmt.Println("SSID", ssid)
	}

	apalist, err := utils.GetNetworkInfo()
	if err != nil {
		fmt.Println("E:", err)
	} else {
		for _, apa := range apalist {
			fmt.Println(apa.Name, apa.Type, apa.Gateway)
		}
	}

	rules, err := rule.GetRules()
	if err != nil {
		fmt.Println("RE:", err)
	}
	if rules != nil {
		fmt.Println(rules.Route_Protect, rules.Ethernet_Protect, rules.Wireless_Protect)
	}

	switch os.Args[1] {
	case "list":
		utils.ListRoutes()
	case "add":
		if len(os.Args) < 6 {
			fmt.Println("Usage: route_manager add <destination> <mask> <gateway> <interface> <metric>")
			return
		}
		utils.AddRoute(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: route_manager delete <destination>")
			return
		}
		utils.DeleteRoute(os.Args[2])
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Windows路由表管理工具")
	fmt.Println("用法:")
	fmt.Println("  route_manager list - 列出所有路由")
	fmt.Println("  route_manager add <destination> <mask> <gateway> <interface> <metric> - 添加路由")
	fmt.Println("  route_manager delete <destination> - 删除路由")
	fmt.Println("示例:")
	fmt.Println("  route_manager add 192.168.1.0 255.255.255.0 192.168.0.1 1 1")
	fmt.Println("  route_manager delete 192.168.1.0")
}
