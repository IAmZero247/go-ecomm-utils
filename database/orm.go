package database

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type OrmDB struct {
	OrmInstance *gorm.DB
	Database    Database
}

func OpenORMWithDatabase(database Database) (*OrmDB, error) {
	ormDB := OrmDB{}
	if database == nil {
		return nil, errors.New("database object is nil")
	}
	var err error
	conn := database.Get().(*sql.DB)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDB.OrmInstance = gormDB
	ormDB.Database = database
	return &ormDB, nil
}

func OpenORM() (*OrmDB, error) {
	postgersDb, err := OpenPostgresSqlDatabase()
	if err != nil {
		return nil, err
	}
	err = postgersDb.Ping()
	if err != nil {
		return nil, err
	}
	return OpenORMWithDatabase(postgersDb)
}
