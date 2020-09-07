package main

import (
	"sample/internal/database"
	"sample/internal/logs"
)

func main() {
	_ = logs.InitLogger()
	client := database.NewMySQLClient("")

}
