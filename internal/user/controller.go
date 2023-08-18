package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"user-service/internal/controllers"
)

const (
	usersUrl        = "/users"
	userUrl         = "/users/:id"
	userGenerateUrl = "/users-generate"
)

type controller struct {
	service Service
}

func NewController(service Service) controllers.Controller {
	return &controller{service: service}
}

func (c *controller) Register(router *httprouter.Router) {
	router.GET(usersUrl, c.GetList)
	router.GET(userUrl, c.GetUserById)
	router.POST(usersUrl, c.CreateUser)
	router.PUT(userUrl, c.UpdateUser)
	router.DELETE(userUrl, c.DeleteUser)
	router.GET(userGenerateUrl, c.GenerateRandomUser)
}

func (c *controller) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := context.TODO()

	users, err := c.service.GetAllUsers(ctx)

	if err != nil {
		http.Error(w, "Failed to get users", 500)
	}

	w.WriteHeader(200)

	json.NewEncoder(w).Encode(users)
}

func (c *controller) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("GetUserById"))
}

func (c *controller) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)

	r.ParseForm()

	fmt.Println("Create new User with name ", r.FormValue("firstName"))

	w.Write([]byte("Ok"))
}

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("UpdateUser"))
}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("DeleteUser"))
}

func (c *controller) GenerateRandomUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := context.TODO()

	user, err := c.service.GenerateRandomUser(ctx)

	if err != nil {
		http.Error(w, "Failed to generate user", 500)
	}

	w.WriteHeader(200)

	json.NewEncoder(w).Encode(user)
}
