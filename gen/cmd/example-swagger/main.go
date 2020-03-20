package main

import (
	"fmt"

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
	}
	migrations.Migrate(config.PostgresURI, config.MigrationDirectory, config.MigrationDirection)
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
