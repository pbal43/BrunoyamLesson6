package internal

import "flag"

type Config struct {
	Host string
	Port int
	// TODO: DB connection string
	// TODO: Debug bool
}

func ReadConfig() Config {
	var config Config
	flag.StringVar(&config.Host, "host", "127.0.0.1", "Server host")
	flag.IntVar(&config.Port, "port", 8080, "Server port")
	flag.Parse()
	return config
}
