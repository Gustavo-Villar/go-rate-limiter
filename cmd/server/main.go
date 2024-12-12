package main

import (
	"github.com/gustavo-villar/go-rate-limiter/config"
	"github.com/gustavo-villar/go-rate-limiter/router"
)

func main() {
	config.Init()
	router.Init()
}
