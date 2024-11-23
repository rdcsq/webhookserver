package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"webhookserver/controllers"
	"webhookserver/middleware"
	"webhookserver/structs"
	"webhookserver/utils"
)

func main() {
	structs.InitializeEnv()

	if cli() {
		return
	}

	startApi()
}

func cli() bool {
	if len(os.Args) < 2 {
		return false
	}

	jwtfs := flag.NewFlagSet("jwt", flag.ContinueOnError)
	name := jwtfs.String("name", "", "name of the JWT token")
	time := jwtfs.Int("time", 0, "not before seconds")

	switch os.Args[1] {
	case "jwt":
		err := jwtfs.Parse(os.Args[2:])
		if err != nil {
			return false
		}

		token, err := utils.CreateJwt(*name, *time)
		if err != nil {
			fmt.Println("An error occured creating the JWT token")
			fmt.Println(err)
			return false
		}
		println(token)
	default:
		return false
	}

	return true
}

func startApi() {
	structs.ParseConfig()

	watcher := utils.WatchConfig()
	defer watcher.Close()

	router := http.NewServeMux()
	router.HandleFunc("POST /execute/{id}", controllers.WebhookHandler)

	server := http.Server{
		Addr:    structs.Env.ListeningAddress,
		Handler: middleware.CreateStack(middleware.Logging, middleware.Auth)(router),
	}

	log.Printf("starting server on %s\n", structs.Env.ListeningAddress)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
