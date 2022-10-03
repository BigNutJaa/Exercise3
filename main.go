package main

import (
	"github.com/BigNutJaa/sortDESC"
	"log"
)

func main() {

	result, err := sortDESC.Sorta("bignut")
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(result)
}
