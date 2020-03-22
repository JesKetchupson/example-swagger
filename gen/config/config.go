package config

import (
	"os"
	"strconv"

	"github.com/jessevdk/go-flags"
)

//Config is service config structure
type Config struct {
	MigrationDirection   int
	ReconnectionAttempts int
	Port                 int
	PostgresURI          string
	MigrationDirectory   string
	TLSCert              flags.Filename
	TLSKey               flags.Filename
}

//Configure reads config and return it  from env
func Configure() Config {
	md, err := strconv.Atoi(os.Getenv("MIG_DIRECTION"))
	if err != nil {
		md = 1
	}
	ra, err := strconv.Atoi(os.Getenv("RECONNECTION_ATTEMPTS"))
	if err != nil {
		ra = 5
	}
	port, err := strconv.Atoi(os.Getenv("RECONNECTION_ATTEMPTS"))
	if err != nil {
		port = 8080
	}
	return Config{
		PostgresURI:          os.Getenv("POSTGRES_URI"), //
		Port:                 port,
		TLSCert:              flags.Filename(os.Getenv("TLS_CERTIFICATE")),
		TLSKey:               flags.Filename(os.Getenv("TLS_PRIVATE_KEY")),
		MigrationDirectory:   os.Getenv("MIG_DIRECTORY"),
		MigrationDirection:   md,
		ReconnectionAttempts: ra,
	}
}
