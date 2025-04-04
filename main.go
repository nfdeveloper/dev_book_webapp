package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nfdeveloper/dev_book_webapp/src/router"
	"github.com/nfdeveloper/dev_book_webapp/src/utils"
)

func main() {
	fmt.Println("Rodando WebApp!")
	utils.CarregarTemplates()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
