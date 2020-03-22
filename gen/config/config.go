package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

//Config is service config structure
type Config struct {
	PostgresURI        string
	Port               int
	TLSCert            flags.Filename
	TLSKey             flags.Filename
	MigrationDirectory string
	MigrationDirection int
}

//Configure reads config and return it  from env
func Configure() Config {
	return Config{
		PostgresURI:        os.Getenv("POSTGRES_URI"), //
		Port:               8080,                      //os .GetEnv("PORT")
		TLSCert:            flags.Filename(os.Getenv("TLS_CERTIFICATE")),
		TLSKey:             flags.Filename(os.Getenv("TLS_PRIVATE_KEY")),
		MigrationDirectory: os.Getenv("MIG_DIRECTORY"), //
		MigrationDirection: 1,                          //
	}
}
