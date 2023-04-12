package main

import (
	"orderManagement/initialize"
)

func main() {
	initialize.Migrate()

	initialize.ConnectDB()

	StartServer()
}
