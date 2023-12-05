package main

import (
	"fmt"
	"net/http"
	"chatbox/controllers"
)


func main() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/ws", controllers.HandleConnections)

	go controllers.HandleMessages()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
