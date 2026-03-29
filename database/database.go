package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/IAmZero247/go-ecomm-utils/utils"
)

type DBOptions struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	Protocol        string
	ConnMaxLifeTime time.Duration
	MaxOpenConn     int
	MaxIdleConn     int
	PARAM           string
}
type Database interface {
	Open(options DBOptions)
	Get() interface{}
	Close()
	Ping() error
}

func BuildDbUrl(options DBOptions) (string, error) {
	handleError := func(msg string) (string, error) { return "", errors.New(msg) }
	if utils.IsBlank(options.Username) {
		return handleError("user name cannot be empty")
	}
	if utils.IsBlank(options.Password) {
		return handleError("password cannot be empty")
	}
	if utils.IsBlank(options.Host) {
		return handleError("host name cannot be empty")
	}
	if options.Port <= 0 {
		return handleError("port cannot be 0 or negative")
	}
	if utils.IsBlank(options.Database) {
		return handleError("database cannot be empty")
	}

	var protocol string
	if utils.IsBlank(options.Protocol) {
		protocol = "tcp"
	} else {
		protocol = options.Protocol
	}
	fmt.Println(protocol)
	var param string
	if utils.IsBlank(options.PARAM) {
		param = "parseTime=true"
	} else {
		param = options.PARAM
	}
	fmt.Println(param)
	var dburl string = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d", options.Username,
		options.Password, options.Database, options.Host, options.Port)
	fmt.Println("dbUrl: ", dburl)
	return dburl, nil
}
