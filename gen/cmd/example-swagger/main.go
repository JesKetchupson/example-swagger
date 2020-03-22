package main

import (
	"fmt"
	"os"
	"time"

	"gen/config"
	"gen/migrations"
	"gen/restapi"
	"gen/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-pg/pg"
)

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	config := config.Configure()
	opt, err := pg.ParseURL(config.PostgresURI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for attempts := 0; err != nil; err = migrations.Do(config.PostgresURI, config.MigrationDirectory, config.MigrationDirection) {
		attempts++
		if err != nil && attempts == 3 {
			fmt.Println(err)
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}

	restapi.Init(pg.Connect(opt))
	api := operations.NewTaskListAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	// set the port this service will be run on

	server.ConfigureAPI()
	server.TLSPort = config.Port
	server.TLSCertificate = config.TLSCert
	server.TLSCertificateKey = config.TLSKey
	if err = server.Serve(); err != nil {
		panic(err)
	}

}
