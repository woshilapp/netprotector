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
}
