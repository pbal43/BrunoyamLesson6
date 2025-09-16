package internal

import "flag"

type Config struct {
	Host string
	Port int
	DSN  string
	// TODO: Debug bool
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", "127.0.0.1", "Server host")
	flag.IntVar(&config.Port, "port", 8080, "Server port")
	flag.StringVar(&config.DSN, "dsn", "postgres://postgres:postgres@localhost/postgres", "DB CONNECTION STRING")
	flag.Parse()
	return config
}
