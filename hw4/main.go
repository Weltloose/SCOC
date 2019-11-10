package main

import (
	"github.com/Weltloose/hw4/router"
)

func main() {
	s := router.GetServer()
	s.Run(":8080")
}
