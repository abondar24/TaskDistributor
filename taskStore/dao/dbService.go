package dao

import (
	"database/sql"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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

	return &Database{inst: instance}
}

func buildConnectionUri(conf *config.Config) string {
	var uri strings.Builder

	uri.WriteString(conf.Database.Username)
	uri.WriteString(":")
	uri.WriteString(conf.Database.Password)
	uri.WriteString("@")
	uri.WriteString("tcp")
	uri.WriteString("(")
	uri.WriteString(conf.Database.Host)
	uri.WriteString(":")
	uri.WriteString(strconv.Itoa(conf.Database.Port))
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

	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Println(err.Error())
		}
	}(tx)

	return tx, err
}

func (db *Database) CloseTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
