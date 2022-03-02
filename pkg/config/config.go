package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	Name 						  string 		 `split_words:"true" default:"Shopping Cart Service"`
	RootUrl                       string         `split_words:"true" default:"checkout"`
	Port                          string         `default:":9000"`
}

func Get() Configuration {
	config := Configuration{}
	envconfig.MustProcess("checkout", &config)
	return config
}