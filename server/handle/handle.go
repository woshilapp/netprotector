package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/woshilapp/netprotector/server/global"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, "Request", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func LoggerHandler(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, "Request", r.RequestURI)
		next(w, r)
	})
}

func ValidToken(token string) bool {
	return slices.ContainsFunc(global.Tokens, func(s string) bool { return s == token })
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World NetProtector Server!")

	global.Clients[r.RemoteAddr] = time.Now()
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	var Token global.Token
	err := json.NewDecoder(r.Body).Decode(&Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ValidToken(Token.Token) {
		data := struct {
			Status           int
			Client           int
			Route_Protect    bool
			Wireless_Protect bool
			Ethernet_Protect bool
		}{
			1,
			len(global.Clients),
			global.Rule.Route_Protect,
			global.Rule.Wireless_Protect,
			global.Rule.Ethernet_Protect,
		}

		encoded, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(encoded)
	} else {
		data := struct {
			Status int
		}{
			Status: 0,
		}

		encoded, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(encoded)
	}
}

func rulesHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Route_Protect    bool
		Wireless_Protect bool
		Ethernet_Protect bool
		Route_Rules      []global.RouteRule
		Wireless_Rules   []global.WirelessRule
		Time_Rules       []global.TimeRule
	}{
		global.Rule.Route_Protect,
		global.Rule.Wireless_Protect,
		global.Rule.Ethernet_Protect,
		global.Rule.Route_Rules,
		global.Rule.Wireless_Rules,
		global.Rule.Time_Rules,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}
