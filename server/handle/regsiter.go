package handle

import "net/http"

func RegsiterHandles(server *http.ServeMux) {
	// hello handler
	server.Handle("/api/hello", LoggerHandler(helloHandler))

	// auth handler
	server.Handle("/api/auth/login", LoggerHandler(loginHandler))
	server.Handle("/api/auth/logout", LoggerHandler(logoutHandler))

	// status handler
	server.Handle("/api/status", LoggerHandler(statusHandler))

	// rules handler
	server.Handle("/api/rules", LoggerHandler(rulesHandler))

	// modify handler
	server.Handle("/api/modify/route-protect", LoggerHandler(routeProtectHandler))
	server.Handle("/api/modify/wireless-protect", LoggerHandler(wirelessProtectHandler))
	server.Handle("/api/modify/ethernet-protect", LoggerHandler(ethernetProtectHandler))
	server.Handle("/api/modify/route-rules", LoggerHandler(routeRulesHandler))
	server.Handle("/api/modify/wireless-rules", LoggerHandler(wirelessRulesHandler))
	server.Handle("/api/modify/time-rules", LoggerHandler(timeRulesHandler))
}
