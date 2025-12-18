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

	accessToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwaS5yZC5zZXJ2aWNlcyIsInN1YiI6InRsanByZEJUbGtUazZfaXZ3QmJXeDhkWW1HUXd2RnJ6XzRER0V5Y2pLYk1AY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXBwLnJkc3RhdGlvbi5jb20uYnIvYXBpL3YyLyIsImFwcF9uYW1lIjoiQkkgLSBDb250cm9sIEY1IiwiZXhwIjoxNzY2MDg3OTE4LCJpYXQiOjE3NjYwMDE1MTgsInNjb3BlIjoiIn0.RbJe4H_dhVOCiTIbrfYOS3JROVmhVdTOifUzwelecw7-rKL_BwYgnUEZ6Y9sPkQHBaKhinJyQpTVDPO-ziVKOIQg5IG6gqPK5g3AsxbHtqGhAM0-LcpSlwia8KsCQ9xZCgvo4HOqCkKv6KW_LURSwoTm5yZ2SgaQ6s5nV_UiPKNC3cR24qd8oXUoaNbjlKBTZot6eKbiwpIOm7AsL3uW8N8VS2Qk5p89oHUqMPTJ_y_oEfImq6cPFLbbDWjFDoQov24SX99LoaT-RuNxwiJhrKqC7uhVUckVyKfu7xkB7IIV6Gee9LyWHE4q9reTD478guFq2VsMXpUq-OKRICi8MQ"

	result, err := handlers.FetchContacts(accessToken, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULTADO:", result)
}
