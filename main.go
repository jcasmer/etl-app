package main

import (
	"etl-app/build/cmd"
	"log"
)

func main() {

	err := cmd.Commands()
	if err != nil {

		log.Fatalln("error unmarshal:", err)
	}
}
