package main

import (
	"RestApiBackend/infrastructure"
	appServer "RestApiBackend/internal/server"
)

func main() {
	app, dbConnection := infrastructure.App()

	appServer.NewServer(app, dbConnection).Run()
}

// https://outcomeschool.com/blog/go-backend-clean-architecture
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
// https://hariesef.medium.com/go-e2e-tutorial-part-1-clean-architecture-and-folder-structure-4ae6c486867c
// https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API
