package configs

import (
	pg "VinylShop/pkg/drivers/postgressql"
	"github.com/joho/godotenv"
	"os"
)

type ApConfig struct {
	Addr string
	Db   *pg.PGCFG
}

func LoadAppConfig() *ApConfig {
	if err := godotenv.Load(); err != nil {
		return nil
	}
	return &ApConfig{
		Addr: os.Getenv("ADDR"),
		Db: &pg.PGCFG{
			Username: os.Getenv("DBUSER"),
			Password: os.Getenv("DBPASS"),
			Host:     os.Getenv("DBHOST"),
			Port:     os.Getenv("DBPORT"),
			Database: os.Getenv("DBNAME"),
			Type:     os.Getenv("DB"),
		},
	}
}
