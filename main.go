package main

import (
	"fmt"
	_config "go-connect-mysql/config"
)

func main() {
	_config.Oauth2Service()
	fmt.Println("starting web server at http://localhost:8080/")
	_config.Controller()
}
