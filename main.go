package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nfdeveloper/dev_book_webapp/src/router"
)

func main() {
	fmt.Println("Rodando WebApp!")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
