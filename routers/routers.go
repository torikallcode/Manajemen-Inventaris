package routers

import (
	"inventaris/controllers"
	"net/http"
)

func SetupRouter() {
	http.HandleFunc("/users/create", controllers.CreateUser)
	http.HandleFunc("/users/all", controllers.GetAllUser)
	http.HandleFunc("/users/get", controllers.GetUserById)
	http.HandleFunc("/users/update", controllers.UpdateUser)
	http.HandleFunc("/users/delete", controllers.DeleteUser)
}
