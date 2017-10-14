package server

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"strconv"
	"fmt"
)

var connection *sql.DB

type UserDraft struct {
	Email string
	Name string
	IsActive bool
}

type User struct {
	Id int
	Email string
	Name string
	IsActive bool
}

func connect(host string, port int, user string, password string, database string) *sql.DB {
	var connectString string = user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database

	fmt.Println("Try connect to the database: " + connectString)

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database: " + connectString)

	return db
}

func insertUser(userDraft *UserDraft) (*User, error) {
	stmtIns, err := connection.Prepare("INSERT INTO users (email, name, is_active) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("database error")
	}

	defer stmtIns.Close()
	res, err := stmtIns.Exec(userDraft.Email, userDraft.Name, userDraft.IsActive)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("user is not unique")
	}

	id, _ := res.LastInsertId()

	return &User{
		int(id),
		userDraft.Email,
		userDraft.Name,
		userDraft.IsActive,
	}, nil
}

func getUser(email string) (*User, error) {
	stmtOut, err := connection.Prepare("SELECT id, email, name, is_active FROM users WHERE email = ?")
	var user User
	err = stmtOut.QueryRow(email).Scan(&user.Id, &user.Email, &user.Name, &user.IsActive)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("database error")
	}

	return &user, nil
}

func getUserById(id int) (*User, error) {
	stmtOut, err := connection.Prepare("SELECT id, email, name, is_active FROM users WHERE id = ?")
	var user User
	err = stmtOut.QueryRow(id).Scan(&user.Id, &user.Email, &user.Name, &user.IsActive)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("database error")
	}

	return &user, nil
}

func countAllUsers() (uint, error) {
	stmtOut, err := connection.Prepare("SELECT count(id) FROM users")
	var count uint
	err = stmtOut.QueryRow().Scan(&count)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("database error")
	}

	return count, nil
}

func update(user *User) error {
	stmtOut, err := connection.Prepare("UPDATE users SET name=?, email=?, is_active=? WHERE id=?")
	_, err = stmtOut.Exec(user.Name, user.Email, user.IsActive, user.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func updateIsActiveUser(id int, isActive bool) error {
	stmtOut, err := connection.Prepare("UPDATE users SET is_active=? WHERE id=?")
	_, err = stmtOut.Exec(&isActive, &id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getAllUsers(offset uint, limit uint) ([]User, error) {
	rows, err := connection.Query("SELECT id, email, name, is_active  FROM users ORDER BY id LIMIT ?, ?", offset, limit)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("database error")
	}

	var users []User = []User{}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Name, &user.IsActive)

		users = append(users, user)
	}

	return users, nil
}
