package global

import (
	"sync"
	"time"
)

type TimeRule struct {
	Time_Start  string
	Time_End    string
	Description string
}

type RouteRule struct {
	Network     string
	Mask        string
	Endpoint    string
	Description string
}

type WirelessRule struct {
	SSID        string
	Description string
}

type User struct {
	Username     string
	PasswordHash string
}

type Rules struct {
	Route_Protect    bool
	Ethernet_Protect bool
	Wireless_Protect bool
	Time_Rules       []TimeRule
	Route_Rules      []RouteRule
	Wireless_Rules   []WirelessRule
}

type Config struct {
	Port int
}

var (
	Cfg   *Config = &Config{}
	Rule  *Rules  = &Rules{}
	Users []*User = []*User{}
)

var (
	Clients   map[string]time.Time = map[string]time.Time{}
	Tokens    []string             = []string{}
	TokenLock sync.Mutex           = sync.Mutex{}
)

func CleanClients() {
	for {
		time.Sleep(10 * time.Second)
		for k, v := range Clients {
			if time.Since(v) > time.Second*20 {
				delete(Clients, k)
			}
		}
	}
}
