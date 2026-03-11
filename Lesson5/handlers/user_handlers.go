package handlers

import (
	"encoding/json"
	"net/http"
	"user-api/auth"
	"user-api/models"
	"user-api/repository"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается. Нужен POST.", http.StatusMethodNotAllowed)
		return
	}
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Плохой запрос.", http.StatusBadRequest)
		return
	}
	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		http.Error(w, "Не удалось хешировать пароль.", http.StatusInternalServerError)
		return
	}
	u.Password = hashedPassword
	id, err := repository.RegisterUser(u)
	if err != nil {
		http.Error(w, "Ошибка сервера.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается. Нужен POST.", http.StatusMethodNotAllowed)
		return
	}
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Плохой запрос.", http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByLogin(input.Login)
	if err != nil {
		http.Error(w, "Пользователь не найден.", http.StatusNotFound)
		return
	}
	if !auth.CheckPassword(input.Password, user.Password){
		http.Error(w, "Неправильный пароль.", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateToker(user.ID)
	if err != nil {
		http.Error(w, "Ошибка сервера.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается. Нужен Get.", http.StatusMethodNotAllowed)
		return
	}
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Ошибка сервера.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}