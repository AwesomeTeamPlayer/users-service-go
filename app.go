package main

import (
	"github.com/AwesomeTeamPlayer/users-service-go/server"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Start")
	rand.NewSource(time.Now().UnixNano())
	server.StartServer()
}