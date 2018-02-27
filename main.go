package main

import (
	"github.com/ikropotov/mobilelogs/routes"
)


func main() {

	r := routes.CreateRouter()

	err := r.Run()

	if err != nil {
		panic("Server stopped with err" )
	}

}

