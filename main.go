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

	accessToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwaS5yZC5zZXJ2aWNlcyIsInN1YiI6InRsanByZEJUbGtUazZfaXZ3QmJXeDhkWW1HUXd2RnJ6XzRER0V5Y2pLYk1AY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXBwLnJkc3RhdGlvbi5jb20uYnIvYXBpL3YyLyIsImFwcF9uYW1lIjoiQkkgLSBDb250cm9sIEY1IiwiZXhwIjoxNzY2MDAwMDcxLCJpYXQiOjE3NjU5MTM2NzEsInNjb3BlIjoiIn0.DqPd9tmwCXwp6VgffDfC0f82IA9-tauCbyIm8T5ugU_Uz07dDGs49VYq_ds6r7tfqyXnUzi0anEr7va_tT890oQ0JZnP7d13aN6gvracjiQCrvuBaYXEx72Xn4OFnX95ZB7n6nEs1y3U9p46jRPssm1zR9ENDjMi8zKktRKtU3O1kJxtf-g3ShCRS0DAtlaBg34Byrxq8sGhSvJ8Ba_96HAvn0GfUSj8XU8ZLUlXkK0sDE4WIeCIKE_ITw3AHb8qZ8sWeJlX3_L7_P0CBG_wzpdrOmZQafe-9dTk_PvkHJu5zUnCnZ_BDfC0g9fCWUSEJSQu4_dRubV072tawpjVJQ"

	result, err := handlers.FetchContacts(accessToken, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULTADO:", result)
}
