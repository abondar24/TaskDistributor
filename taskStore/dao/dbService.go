package dao

import (
	"database/sql"
	"errors"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	inst *sql.DB
}

func InitDatabase(conf *config.Config) *Database {
	connUri := buildConnectionUri(conf)
	instance, err := sql.Open(conf.Database.Driver, connUri)
	if err != nil {
		log.Fatalln(err.Error())
	}

	driver, err := mysql.WithInstance(instance, &mysql.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+conf.Database.MigrationPath, // Path to your migration files
		conf.Database.Database, driver)
	if err != nil {
		log.Fatalln(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}

	return &Database{inst: instance}
}

func buildConnectionUri(conf *config.Config) string {
	var uri strings.Builder

	uri.WriteString(conf.Database.Username)
	uri.WriteString(":")
	uri.WriteString(conf.Database.Password)
	uri.WriteString(conf.Broker.Password)
	uri.WriteString("@")
	uri.WriteString("tcp")
	uri.WriteString("(")
	uri.WriteString(conf.Broker.Host)
	uri.WriteString(":")
	uri.WriteString(strconv.Itoa(conf.Broker.Port))
	uri.WriteString(")")
	uri.WriteString("/")
	uri.WriteString(conf.Database.Database)
	uri.WriteString("?charset=utf8")

	return uri.String()
}

func (db *Database) BeginTx() (*sql.Tx, error) {
	tx, err := db.inst.Begin()
	if err != nil {
		log.Println(err.Error())

	}

	return tx, err
}
