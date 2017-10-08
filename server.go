package main

import (
	"net/http"
	"errors"
)

type EmailRequest struct {
	Email string
}

type IdRequest struct {
	Id uint
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
	user, err := insertUser(userDraft)

	if err != nil {
		return errors.New("database error")
	}
	*result = *user
	return nil
}

func (t *App) GetUserByEmail (r *http.Request, emailRequest *EmailRequest, result *User) error {
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

func (t *App) ActiveUser (r *http.Request, idRequest *IdRequest, result *bool) error {
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

	err = updateIsActiveUser(idRequest.Id, true)
	if err != nil {
		return errors.New("database error")
	}

	*result = true
	return nil
}

func (t *App) InactiveUser (r *http.Request, idRequest *IdRequest, result *bool) error {
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

	err = updateIsActiveUser(idRequest.Id, false)
	if err != nil {
		return errors.New("database error")
	}

	*result = true
	return nil
}
