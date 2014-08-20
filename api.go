package main

import (
	"os"

	"github.com/camarabook/camarabook-api/service"
)

func main() {
	scamarabook := service.CamarabookService{}
	runner := scamarabook.Run()

	runner.Run(":" + os.Getenv("PORT"))
}
