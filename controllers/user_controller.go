package controllers

import (
	"fmt"
	"inventaris/connection"
	"inventaris/models"
	"time"
)

func CreateUser(name string, email string, age int) {
	user := models.User{
		Name:     name,
		Email:    email,
		Age:      age,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	result := connection.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("Gagal menambahkan user", result.Error)
		return
	}
	fmt.Println("User berhasil ditambahkan dengan ID: ", user.ID)
}

func GetUser(id int) {
	var user models.User
	result := connection.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println("Gagal mengambil User")
		return
	}
	fmt.Printf("ID: %d\nName: %s\nEmail: %s\n", user.ID, user.Name, user.Email)
}

func UpdateUser(id int, newName string) {
	var user models.User

	result := connection.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println("Gagal mengambil user")
		return
	}

	user.Name = newName
	connection.DB.Save(&user)
}

func DeleteUser(id int) {
	connection.DB.Delete(&models.User{}, id)
	fmt.Print("Berhasil Hapus User")
}