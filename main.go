package main

import (
	"csrf/db"
	"csrf/server"
	"csrf/server/middleware/myJwt"
	"log"
)

var host = "localhost"
var port = "9000"

func main() {
	// init the DB
	err := db.InitDB()
	if err != nil {
		log.Panic(err)
	} else {
		log.Print("Connected!")
	}
	

	// init the JWTs
	jwtErr := myJwt.InitialiseJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(jwtErr)
	}

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
