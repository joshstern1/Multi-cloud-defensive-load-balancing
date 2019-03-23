package main

import (
	"time"
	"fmt"
)

func createUser(userObj User) User {
	query := fmt.Sprintf(
		"INSERT INTO users (username, password, email) "+
			"VALUES ('%s', '%s', '%s') "+
			"RETURNING id",
		userObj.Username,
		userObj.Password,
		userObj.Email)
	
	s.Set(time.Now().String(), query)

	db := dbLogin()
	defer db.Close()
	row, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	row.Next()
	var newID int
	scanErr := row.Scan(&newID)

	if scanErr != nil {
		panic(scanErr)
	}

	userObj.ID = newID
	return userObj
}

func getUserById(userID int) (User) {
	db := dbLogin()
	defer db.Close()

	rows, err := db.Query(
		"SELECT * FROM users WHERE id = $1",
		userID)

	for rows.Next() {
		var id int
		var db_username string
		var password string
		var email string

		err = rows.Scan(&id, &db_username, &password, &email)
		if err != nil {
			panic(err)
		}

		return User{ID: id, Username: db_username, Email: email}
	}

	return User{}
}

func getUsers() ([]User) {
	db := dbLogin()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	var users []User
	for rows.Next() {
		var id int
		var db_username string
		var password string
		var email string

		err = rows.Scan(&id, &db_username, &password, &email)
		if err != nil {
			panic (err)
		}
		users = append(users,
			User{ID: id, Username: db_username, Email: email})
	}
	return users
}

func getUser(username string, password string) (User) {
	db := dbLogin()
	defer db.Close()

	rows, err := db.Query(
		"SELECT * FROM users WHERE username = $1 AND password = $2",
		username,
		password)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var db_username string
		var password string
		var email string

		err = rows.Scan(&id, &db_username, &password, &email)
		if err != nil {
			panic(err)
		}

		return User{ID: id, Username: db_username, Email: email}
	}

	return User{}
}

func deleteUser(userID int) {
	db := dbLogin()
	defer db.Close()

	_, err := db.Query(
		"DELETE FROM users where users.id = $1",
		userID)

	if err != nil {
		panic(err)
	}
}
