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

	cfg := config.Configure()
	opt, err := pg.ParseURL(cfg.PostgresURI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for err = migrations.Do(cfg.PostgresURI, cfg.MigrationDirectory, cfg.MigrationDirection); err != nil; {
		if err != nil && cfg.ReconnectionAttempts == 0 {
			fmt.Println(err)
			os.Exit(1)
		}
		cfg.ReconnectionAttempts--
		time.Sleep(5 * time.Second)
	}

	restapi.Init(pg.Connect(opt))
	api := operations.NewTaskListAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	// set the port this service will be run on

	server.ConfigureAPI()
	server.TLSPort = cfg.Port
	server.TLSCertificate = cfg.TLSCert
	server.TLSCertificateKey = cfg.TLSKey
	if err = server.Serve(); err != nil {
		panic(err)
	}

}
