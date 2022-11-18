package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func NewPostgresDb(cfg *Config) (*sqlx.DB, error) {

	//pq.Open("postgres")
	if cfg == nil {
		cfg = new(Config)
		{
			cfg.Host = "localhost"
			cfg.Port = "5434"
			cfg.Username = "postgres"
			cfg.Password = "1qaz2wsx"
			cfg.DbName = "avitodb"
			cfg.SSLMode = "disable"
		}
	}

	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("DB connected successfully")
	return db, nil
}
