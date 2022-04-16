// main package, app entry point
package main

import (
	"flag"
	"fmt"
	"go-gin-mysql-boilerplate/lib/config"
	server "go-gin-mysql-boilerplate/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//get env key payment
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//get from your env
	var MODE string = os.Getenv("MODE")

	environment := flag.String("e", MODE, "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	config.Init(*environment)

	//GetConnection
	config.OpenConnection()

	//init
	server.Init()

}
