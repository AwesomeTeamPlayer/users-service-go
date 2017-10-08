package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
	"fmt"
	"os"
	"strconv"
)

func main() {

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