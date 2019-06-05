package driver

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

const driver = "postgres"

type config struct {
	Db            string
	Host          string
	Port          int
	User          string
	Password      string
	Params        string
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPass     string
	MongoDatabase string
	SlackToken    string
}

var c config

func ConnectDB() {
	err := envconfig.Process("PG", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Db)
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

}
