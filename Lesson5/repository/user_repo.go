package repository

import (
	"user-api/db"
	"user-api/models"
)

// func CreateUser(u models.User) (int64, error) {
// 	res, err := db.DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", u.Name, u.Age)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return res.LastInsertId()
// }

func RegisterUser(u models.User) (int64, error) {
	res, err := db.DB.Exec(`INSERT INTO users (name, age, login, password)
		VALUES (?, ?, ?, ?)`, u.Name, u.Age, u.Login, u.Password)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func GetUserByLogin(login string) (models.User, error) {
	var u models.User
	err := db.DB.QueryRow("SELECT * FROM users WHERE login = ?",
		login,).Scan(&u.ID, &u.Name, &u.Age, &u.Login, &u.Password)
	return u, err
}

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, name, age, login FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Login)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}