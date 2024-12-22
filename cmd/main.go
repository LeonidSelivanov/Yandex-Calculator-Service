package main

import (
	"github.com/LeonidSelivanov/Yandex-Calculator-Service/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
