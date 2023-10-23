package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type DB struct {
	*sql.DB
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func CfgFromEnv() (DBConfig, error) {
	u, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		return DBConfig{}, err
	}

	host := u.Hostname()
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return DBConfig{}, err
	}
	user := u.User.Username()
	password, _ := u.User.Password()
	name := u.Path[1:]

	return DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}, nil
}

func NewDB(cfg DBConfig) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
