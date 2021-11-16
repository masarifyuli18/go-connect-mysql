package main

import (
	"fmt"
	_config "go-connect-mysql/config"
)

func main() {
	fmt.Println("starting web server at http://localhost:8080/")
	_config.Controller()
	_config.SecurityConfig()
}
