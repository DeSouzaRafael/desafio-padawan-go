package main

import (
	"coinConversion/config"
	"coinConversion/controller"
	"log"
	"os"
	"time"
)

func main() {

	time.Sleep(30 * time.Second)

	if db, err := config.DatabaseInit(); err != nil {
		log.Fatal(err)
	} else {
		controller.SetDb(db)
	}

	serverPort := ":" + os.Getenv("PORT")
	webServer(serverPort)
}
