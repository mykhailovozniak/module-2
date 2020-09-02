package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mykhailovozniak/module2/src/html"
	"log"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request)  {
	_, err := fmt.Fprintf(res, "Hello world")

	if err != nil {
		log.Fatal("Error", err)
	}

	log.Print("Successfully return html - hello")
}

func learningMaterials(res http.ResponseWriter, req *http.Request)  {
	var list = []string{
		"http package",
		"Writing Web Applications",
		"Go by Example: HTTP Servers",
		"Hello world HTTP server example",
		"How I write Go HTTP services after seven years",
		"Gorilla web toolkit",
		"One new Module for testing",
	}

	var ulElem = html.RenderUL(list)
	var document = html.RenderHTML5(ulElem)

	_, err := fmt.Fprintf(res, document)

	if err != nil {
		log.Fatal("Error", err)
	}

	log.Print("Successfully return html- learning-materials")
}

func main()  {
	errLoadEnv := godotenv.Load()

	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/materials", learningMaterials)


	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Fatal("Error during start app")
	}
}
