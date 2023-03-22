package config

var Conf Config

type Config struct {
	URL  string
	Port string `env:"PORT" envDefault:"8080"`
}
