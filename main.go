package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"rdstation-crm/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	accessToken := ""

	result, err := handlers.FetchContacts(accessToken, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULTADO:", result)
}
