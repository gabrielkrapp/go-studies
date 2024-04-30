package main

import (
	"banco-de-dados/dbconfig"
	"fmt"
	"log"
)

func main() {
	_, erro := dbconfig.DatabaseConnect()
	if erro != nil {
		fmt.Println("Error connecting to database")
		log.Fatal(erro)
	}
}
