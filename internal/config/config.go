package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	DBStorageMode bool
	DB            *DB
	Listen        *Listen `yaml:"listen"`
	VersionDB     int     `yaml:"db_version"`
}

type DB struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	NameDB     string `yaml:"name_db"`
	UserName   string `yaml:"user_name"`
	DBPassword string `yaml:"password"`
}

type Listen struct {
	Port string `yaml:"port"`
	IP   string `yaml:"ip"`
}

type Service struct {
	Name      string `json:"name"`
	Port      string `json:"port"`
	IP        string `json:"ip"`
	Endpoints []struct {
		URL     string   `json:"url"`
		Methods []string `json:"methods"`
	} `json:"endpoints"`
}

func (db *DB) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		db.UserName, db.DBPassword, db.Host, db.Port, db.NameDB)
}

func GetConfig(dbStorageMode bool) *Config {
	logger := log.Default()
	instance := &Config{DB: &DB{}, DBStorageMode: dbStorageMode}
	if err := cleanenv.ReadConfig("./conf/config.yml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		logger.Print(help)
		logger.Fatal(err)
	}

	if dbStorageMode {
		dbConfigName := "DBConfig"
		if err := cleanenv.ReadConfig(fmt.Sprintf("./conf/db/%s.yml", dbConfigName), instance.DB); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Print(help)
			logger.Fatal(err)
		}
	}

	return instance
}
