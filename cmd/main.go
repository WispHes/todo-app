package main

import (
	"fmt"
	"github.com/wisphes/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	fmt.Println("hi")
}
