package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/woshilapp/netprotector/server/config"
	"github.com/woshilapp/netprotector/server/global"
	"github.com/woshilapp/netprotector/server/handle"
)

func main() {
	banner := `  _   _          _     ____                          _                   _                  
 | \ | |   ___  | |_  |  _ \   _ __    ___     ___  | |_    ___    ___  | |_    ___    _ __ 
 |  \| |  / _ \ | __| | |_) | | '__|  / _ \   / __| | __|  / _ \  / __| | __|  / _ \  | '__|
 | |\  | |  __/ | |_  |  __/  | |    | (_) | | (__  | |_  |  __/ | (__  | |_  | (_) | | |   
 |_| \_|  \___|  \__| |_|     |_|     \___/   \___|  \__|  \___|  \___|  \__|  \___/  |_|   
 `

	log.Println("Hello World NetProtector Server!")
	fmt.Println(banner)
	fmt.Println("Auctor: woshilapp (github.com/woshilapp/netprotector)\n")

	server := http.NewServeMux()

	// web handler
	server.Handle("/", handle.LoggerMiddleware(http.FileServer(http.Dir("./web"))))
	// client file handler
	clientFileHandler := http.StripPrefix("/client/", http.FileServer(http.Dir("./clientfiles")))
	server.Handle("/client/", handle.LoggerMiddleware(clientFileHandler))
	// api handler
	handle.RegsiterHandles(server)

	// load datas
	log.Println("Loading datas...")

	err := config.ReadConfig()
	if err != nil {
		log.Fatalln("ReadConfig error:", err)
		return
	}
	err = config.ReadRules()
	if err != nil {
		log.Fatalln("ReadRules error:", err)
		return
	}
	err = config.ReadUsers()
	if err != nil {
		log.Fatalln("ReadUsers error:", err)
		return
	}

	port := strconv.Itoa(global.Cfg.Port)

	log.Println("Data loaded")

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
	log.Println("We are good to go!")
	for {
		time.Sleep(10 * time.Second)
	}
}
