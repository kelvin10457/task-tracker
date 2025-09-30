package config

//this file is for reading the environment variables that you need to conect this program to your db
import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
	Params string
}

func Load() (Config, error) {
	_ = godotenv.Load() //this is for loading every environment (from .env) variable and ignoring the error
	cfg := Config{
		User:   os.Getenv("DB_USER"), //you create an instance of a struct config
		Pass:   os.Getenv("DB_PASS"), //with os.Getenv("NAME_OF_VARIABLE") you get the value
		Host:   os.Getenv("DB_HOST"), //of the environment variable
		Port:   os.Getenv("DB_PORT"),
		Name:   os.Getenv("DB_NAME"),
		Params: os.Getenv("DB_PARAMS"),
	}
	//if the length of the missing environment variables is
	//greater than zero you need to check your environment variables

	//we don't check params because it's optional
	missing := []string{}
	if cfg.User == "" {
		missing = append(missing, "DB_USER")
	}
	if cfg.Pass == "" {
		missing = append(missing, "DB_PASS")
	}
	if cfg.Host == "" {
		missing = append(missing, "DB_HOST")
	}
	if cfg.Port == "" {
		missing = append(missing, "DB_PORT")
	}
	if cfg.Name == "" {
		missing = append(missing, "DB_NAME")
	}

	if len(missing) > 0 {
		return cfg, fmt.Errorf("faltan variables: %v", missing)
	}
	return cfg, nil
}

//this functions makes the data source name with data of the struct config that we created before
//the dsn is a string that represents all the info needed for an app to conect with a database

//the format is:
//usermame:password@tcp(host:port)/nameOfDataBase

// the format is like this if we use params:
//
//usermame:password@tcp(host:port)/nameOfDataBase?params
func (config Config) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Pass, config.Host, config.Port, config.Name)
	if config.Params != "" {
		dsn += "?" + config.Params
	}
	return dsn
}
