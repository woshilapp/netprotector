package handle

import "net/http"

func RegsiterHandles(server *http.ServeMux) {
	server.HandleFunc("/api/hello", helloHandler)
}
