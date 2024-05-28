package main

import (
	"fmt"
	"github.com/thgbianeck/api-go-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
