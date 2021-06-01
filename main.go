package main

import (
	"github.com/zosma180/go-poc/router"
)

func main() {
	r := router.Create()
	r.Run() // listen and serve on 0.0.0.0:8080
}
