package main

import (
	"inventaris/connection"
	"inventaris/controllers"
)

func main() {
	connection.ConnectionDatabase()
	// controllers.CreateUser("Akbar", "Torikal@akbar.com", 21)
	// controllers.GetUser(1)
	// controllers.UpdateUser(1, "Torikal Akbar")
	controllers.DeleteUser(1)
}
