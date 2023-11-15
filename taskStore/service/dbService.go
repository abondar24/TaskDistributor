package service

import (
	"database/sql"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	inst *sql.DB
}

type StoreService interface {
	BeginTx() (*sql.Tx, error)

	RollbackTx(tx *sql.Tx)

	CloseTx(tx *sql.Tx) error
}

func InitDatabase(conf *config.Config) *DatabaseService {
	connUri := buildConnectionUri(conf)
	instance, err := sql.Open(conf.Database.Driver, connUri)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &DatabaseService{inst: instance}
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

func (db *DatabaseService) BeginTx() (*sql.Tx, error) {
	tx, err := db.inst.Begin()
	if err != nil {
		log.Println(err.Error())

	}

	return tx, err
}

func (db *DatabaseService) RollbackTx(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func (db *DatabaseService) CloseTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
