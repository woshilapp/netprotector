package global

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

type TimeRules struct {
	Token      string
	Time_Rules []TimeRule
}

type WirelessRules struct {
	Token          string
	Wireless_Rules []WirelessRule
}

type RouteRules struct {
	Token       string
	Route_Rules []RouteRule
}

type Rules struct {
	Route_Protect    bool
	Ethernet_Protect bool
	Wireless_Protect bool
	Time_Rules       []TimeRule
	Route_Rules      []RouteRule
	Wireless_Rules   []WirelessRule
}
