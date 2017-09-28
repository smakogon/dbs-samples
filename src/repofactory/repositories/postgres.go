package repositories

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

//
// NewPostgresRepository constructor
//
func NewPostgresRepository() interface{} {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("databases.postgres.host"), viper.GetString("databases.postgres.database"),
		viper.GetString("databases.postgres.user"), viper.GetString("databases.postgres.password"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	duration, _ := time.ParseDuration(viper.GetString("databases.postgres.max_life_conns"))

	db.SetConnMaxLifetime(duration)
	db.SetMaxIdleConns(viper.GetInt("databases.postgres.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("databases.postgres.max_open_conns"))

	return db
}
