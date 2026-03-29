package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/IAmZero247/go-ecomm-utils/config"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	db *sql.DB
}

func OpenPostgresSqlDatabase() (*PostgresDatabase, error) {
	host := config.Default().GetString("db.postgres.host")
	port := config.Default().GetInt("db.postgres.port")
	username := config.Default().GetString("db.postgres.username")
	password := config.Default().GetString("db.postgres.password")
	database := config.Default().GetString("db.postgres.database")
	connMaxLifeTime := config.Default().GetInt("db.postgres.connMaxLifeTime")
	maxOpenConn := config.Default().GetInt("db.postgres.maxOpenConn")
	maxIdleConn := config.Default().GetInt("db.postgres.maxIdleConn")
	param := config.Default().GetString("db.postgres.param")
	options := DBOptions{
		Host:            host,
		Port:            port,
		Username:        username,
		Password:        password,
		Database:        database,
		Protocol:        "tcp",
		ConnMaxLifeTime: time.Duration(connMaxLifeTime),
		MaxOpenConn:     maxOpenConn,
		MaxIdleConn:     maxIdleConn,
		PARAM:           param,
	}
	fmt.Println("DBOptions : ", options)
	postgresDb := &PostgresDatabase{
		nil,
	}
	postgresDb.Open(options)
	return postgresDb, nil
}

func (m *PostgresDatabase) Open(options DBOptions) {
	dsn, err := BuildDbUrl(options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Opening database connection on host : %s port : %d database : %s username : %s",
		options.Host, options.Port, options.Database, options.Username)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Configuring database connection pool with ConnMaxLifeTime : %d MaxOpenConn : %d MaxIdleConn :  %d",
		options.ConnMaxLifeTime, options.MaxOpenConn, options.MaxIdleConn)
	db.SetConnMaxLifetime(options.ConnMaxLifeTime)
	db.SetMaxOpenConns(options.MaxOpenConn)
	db.SetMaxIdleConns(options.MaxIdleConn)
	m.db = db
	fmt.Printf("Database connection opened on host : %s port : %d database : %s username : %s",
		options.Host, options.Port, options.Database, options.Username)
}

func (m *PostgresDatabase) Get() interface{} {
	if m.db == nil {
		panic("Database connection not initiated. Please call Open()")
	}
	return m.db
}

func (m *PostgresDatabase) Close() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			fmt.Printf("Error while closing database connection %s", err)
			return
		}
	}
	fmt.Println("Database connection closed.")
}

func (m *PostgresDatabase) Ping() error {
	if m.db == nil {
		panic("Database connection not initiated. Please call Open()")
	}
	err := m.db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Postgressql database connection succeeded.")
		return nil
	}
}
