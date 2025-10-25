package internal

import (
	"cmp"
	"flag"
	"os"
)

const (
	defDNS            = "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	defHost           = "0.0.0.0"
	defPort           = 8080
	defMigrationsPath = "migrations"
)

type Config struct {
	Host        string
	Port        int
	DSN         string
	MigratePath string
	Debug       bool
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", defHost, "Server host")
	flag.IntVar(&config.Port, "port", defPort, "Server port")
	flag.StringVar(&config.DSN, "dsn", defDNS, "DB CONNECTION STRING")
	flag.StringVar(&config.MigratePath, "migrate-path", defMigrationsPath, "Path to migrations folder")
	flag.BoolVar(&config.Debug, "debug", false, "Debug mode (уровень логирования)")
	flag.Parse()

	config.DSN = cmp.Or(os.Getenv("DB_DNS"), defDNS)
	config.MigratePath = cmp.Or(os.Getenv("MIGRATE_PATH"), defMigrationsPath)

	return config
}
