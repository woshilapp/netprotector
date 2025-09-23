package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/woshilapp/netprotector/server/handle"
)

func main() {
	port := strconv.Itoa(8080)

	banner := `  _   _          _     ____                          _                   _                  
 | \ | |   ___  | |_  |  _ \   _ __    ___     ___  | |_    ___    ___  | |_    ___    _ __ 
 |  \| |  / _ \ | __| | |_) | | '__|  / _ \   / __| | __|  / _ \  / __| | __|  / _ \  | '__|
 | |\  | |  __/ | |_  |  __/  | |    | (_) | | (__  | |_  |  __/ | (__  | |_  | (_) | | |   
 |_| \_|  \___|  \__| |_|     |_|     \___/   \___|  \__|  \___|  \___|  \__|  \___/  |_|   
 `

	log.Println("Hello World NetProtector Server!")
	fmt.Println(banner)
	fmt.Println("Auctor: woshilapp (github.com/woshilapp)")

	server := http.NewServeMux()

	// web handler
	server.Handle("/", handle.LoggerMiddleware(http.FileServer(http.Dir("./web"))))
	// client file handler
	clientFileHandler := http.StripPrefix("/client/", http.FileServer(http.Dir("./clientfiles")))
	server.Handle("/client/", handle.LoggerMiddleware(clientFileHandler))
	// api handler
	handle.RegsiterHandles(server)

	var err error
	go func() { // Serve Thread
		err = http.ListenAndServe(":"+port, server)
		if err != nil {
			log.Fatalln("ListenAndServe error:", err)
			return
		}
	}()

	if err != nil {
		return
	}

	log.Println("Server started at :" + port)
	for {
		time.Sleep(10 * time.Second)
	}
}
