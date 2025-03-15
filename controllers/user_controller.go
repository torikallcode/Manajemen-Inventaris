package controllers

import (
	"encoding/json"
	"inventaris/connection"
	"inventaris/models"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateUser(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	connection.DB.Create(user)
	json.NewEncoder(w).Encode(Response{Status: http.StatusCreated, Message: "User Created", Data: user})
}

func GetAllUser(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user []models.User

	connection.DB.Find(user)
	json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Message: "User retrived", Data: user})
}

func GetUserById(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var user models.User
	connection.DB.First(&user, id)

	json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Message: "User Retrieved", Data: user})
}

func UpdateUser(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var user models.User
	connection.DB.First(&user, id)

	json.NewDecoder(r.Body).Decode(&user)
	connection.DB.Save(&user)
	json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Message: "User Update", Data: user})
}

func DeleteUser(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var user models.User
	connection.DB.Delete(&user, id)

	json.NewEncoder(w).Encode(Response{Status: http.StatusOK, Message: "User Deleted"})
}

// func CreateUser(name string, email string, age int) {
// 	user := models.User{
// 		Name:     name,
// 		Email:    email,
// 		Age:      age,
// 		CreateAt: time.Now(),
// 		UpdateAt: time.Now(),
// 	}

// 	result := connection.DB.Create(&user)
// 	if result.Error != nil {
// 		fmt.Println("Gagal menambahkan user", result.Error)
// 		return
// 	}
// 	fmt.Println("User berhasil ditambahkan dengan ID: ", user.ID)
// }

// func GetUser(id int) {
// 	var user models.User
// 	result := connection.DB.First(&user, id)
// 	if result.Error != nil {
// 		fmt.Println("Gagal mengambil User")
// 		return
// 	}
// 	fmt.Printf("ID: %d\nName: %s\nEmail: %s\n", user.ID, user.Name, user.Email)
// }

// func UpdateUser(id int, newName string) {
// 	var user models.User

// 	result := connection.DB.First(&user, id)
// 	if result.Error != nil {
// 		fmt.Println("Gagal mengambil user")
// 		return
// 	}

// 	user.Name = newName
// 	connection.DB.Save(&user)
// }

// func DeleteUser(id int) {
// 	connection.DB.Delete(&models.User{}, id)
// 	fmt.Print("Berhasil Hapus User")
// }
