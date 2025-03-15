package main

import (
	"fmt"
	"inventaris/connection"
	"inventaris/routers"
	"net/http"
)

func main() {
	connection.ConnectionDatabase()
	// controllers.CreateUser("Akbar", "Torikal@akbar.com", 21)
	// controllers.GetUser(1)
	// controllers.UpdateUser(1, "Torikal Akbar")
	// controllers.DeleteUser(1)
	routers.SetupRouter()
	fmt.Println("Server sedang berjalan di port :8080")
	http.ListenAndServe(":8080", nil)
}
