package main

import (
	"bank-microservice/internal/app"
)

func main() {
	a := app.New()
	a.Run("release.toml")
}
