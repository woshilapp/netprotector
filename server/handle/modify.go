package handle

import (
	"encoding/json"
	"net/http"

	"github.com/woshilapp/netprotector/server/config"
	"github.com/woshilapp/netprotector/server/global"
)

func validAndReturn(w http.ResponseWriter, r *http.Request, token string) bool {
	if !ValidToken(token) {
		data := struct {
			Status int
		}{
			Status: 0,
		}

		encoded, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}

		w.Write(encoded)
		return false
	}

	data := struct {
		Status int
	}{
		Status: 1,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	w.Write(encoded)

	return true
}

func routeProtectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.RouteProtect
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Route_Protect = arg.Route_Protect
		// write
		config.WriteRules()
	}
}

func wirelessProtectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.WirelessProtect
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Wireless_Protect = arg.Wireless_Protect
		// write
		config.WriteRules()
	}
}

func ethernetProtectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.EthernetProtect
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Ethernet_Protect = arg.Ethernet_Protect
		// write
		config.WriteRules()
	}
}

func timeRulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.TimeRules
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Time_Rules = arg.Time_Rules
		// write rules
		config.WriteRules()
	}
}

func wirelessRulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.WirelessRules
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Wireless_Rules = arg.Wireless_Rules
		// write rules
		config.WriteRules()
	}
}

func routeRulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var arg global.RouteRules
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validAndReturn(w, r, arg.Token) {
			return
		}

		global.Rule.Route_Rules = arg.Route_Rules
		// write rules
		config.WriteRules()
	}
}
