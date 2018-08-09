package main

import (
	"fmt"
    "books/route"
    "books/model"
)

func main() {
	// USEFUL TO ADD DATA FOR DEMONSTRATIONS
    model.InitDatas()

    route.InitRoutes()

	fmt.Println("APP IS READY")
}
