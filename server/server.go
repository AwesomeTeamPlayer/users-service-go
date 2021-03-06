package server

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
	"os"
	"strconv"
	"fmt"
)

type EmailRequest struct {
	Email string
}

type IdRequest struct {
	Id string
}

type UpdateUserNameRequest struct {
	Id string
	Name string
}

type GetAllUsersRequest struct {
	Limit uint
	Page uint
}

type GetAllUsersResponse struct {
	Users []User
	Count uint
}

type App int

func (t *App) AddUser (r *http.Request, userDraft *UserDraft, result *User) error {
	fmt.Println("AddUser")
	user, err := insertUser(userDraft)

	if err != nil {
		return errors.New("database error")
	}

	userCreated(user.Id)

	fmt.Println("User added")
	*result = *user
	return nil
}

func (t *App) GetUserByEmail (r *http.Request, emailRequest *EmailRequest, result *User) error {
	fmt.Println("GetUserByEmail")

	user, err := getUser(emailRequest.Email)
	if err != nil {
		return errors.New("database error")
	}

	if user == nil {
		return errors.New("user does not exist")
	}

	*result = *user
	return nil
}

func (t *App) GetUserById (r *http.Request, idRequest *IdRequest, result *User) error {
	fmt.Println("GetUserById")

	user, err := getUserById(idRequest.Id)
	if err != nil {
		return errors.New("database error")
	}

	if user == nil {
		return errors.New("user does not exist")
	}

	*result = *user
	return nil
}

func (t *App) GetAllUsers (r *http.Request, getAllUsersRequest *GetAllUsersRequest, result *GetAllUsersResponse) error {
	fmt.Println("GetAllUsers")
	limit := uint(getAllUsersRequest.Limit)
	users, err := getAllUsers(uint(getAllUsersRequest.Page * limit), limit)
	if err != nil {
		return errors.New("database error")
	}

	count, err := countAllUsers()
	if err != nil {
		return errors.New("database error")
	}

	*result = GetAllUsersResponse{users, count}
	return nil
}

func (t *App) UpdateName (r *http.Request, request *UpdateUserNameRequest, result *bool) error {
	fmt.Println("UpdateName")
	user, err := getUserById(request.Id)

	if err != nil {
		return errors.New("database error")
	}

	if user == nil {
		return errors.New("user does not exist")
	}

	user.Name = request.Name

	err = update(user)
	if err != nil {
		return errors.New("database error")
	}

	sendUserNameUpdatedEvent(user.Id)
	fmt.Println("Name updated")
	*result = true
	return nil
}

func (t *App) ActiveUser (r *http.Request, idRequest *IdRequest, result *bool) error {
	fmt.Println("ActiveUser")
	user, err := getUserById(idRequest.Id)

	if err != nil {
		return errors.New("database error")
	}

	if user == nil {
		return errors.New("user does not exist")
	}

	if user.IsActive == true {
		*result = false
		return nil
	}

	user.IsActive = true
	err = update(user)
	if err != nil {
		return errors.New("database error")
	}

	sendUserActivatedEvent(user.Id)

	fmt.Println("user activated")
	*result = true
	return nil
}

func (t *App) InactiveUser (r *http.Request, idRequest *IdRequest, result *bool) error {
	fmt.Println("InactiveUser")
	user, err := getUserById(idRequest.Id)
	if err != nil {
		return errors.New("database error")
	}

	if user == nil {
		return errors.New("user does not exist")
	}

	if user.IsActive == false {
		*result = false
		return nil
	}

	user.IsActive = false
	err = update(user)
	if err != nil {
		return errors.New("database error")
	}

	sendUserInactivatedEvent(user.Id)
	fmt.Println("user inactivated")
	*result = true
	return nil
}

func StartServer() {
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	appPort := os.Getenv("APP_PORT")

	connection = connect(
		os.Getenv("MYSQL_HOST"),
		port,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	app := new(App)
	s.RegisterService(app, "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)

	fmt.Println("Server has started on port " + appPort)
	http.ListenAndServe(":" + appPort, r)
}