package ticketapp

import (
	"errors"
	"os"
	"sync"
)

type Envs struct {
	DbUsername string
	DbPassword string
	DbName     string
	Host       string
}

var once sync.Once
var envs *Envs
var initErr error

func InitEnvs() error {
	var once sync.Once
	once.Do(func() {
		envs, initErr = loadEnvs()
	})
	return initErr
}

func GetEnvs() *Envs {
	return envs
}

func loadEnvs() (*Envs, error) {
	envs := Envs{}

	s := os.Getenv("HOST")
	if s != "" {
		envs.Host = s
	} else {
		return nil, errors.New("fail to get env: HOST")
	}

	s = os.Getenv("DB_USER")
	if s != "" {
		envs.DbUsername = s
	} else {
		return nil, errors.New("fail to get env: DB_USER")
	}

	s = os.Getenv("DB_PASSWORD")
	if s != "" {
		envs.DbPassword = s
	} else {
		return nil, errors.New("fail to get env: DB_PASSWORD")
	}

	s = os.Getenv("DB_NAME")
	if s != "" {
		envs.DbName = s
	} else {
		return nil, errors.New("fail to get env: DB_NAME")
	}

	return &envs, nil
}
