package dao

import (
	"database/sql"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"time"
)

func ConvertTime(rawTime sql.RawBytes) (time.Time, error) {
	createdAt, err := time.Parse(model.TimeFormat, string(rawTime))
	if err != nil {
		return time.Time{}, err
	}

	return createdAt, nil
}
