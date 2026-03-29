package main

import (
	"fmt"

	"github.com/IAmZero247/go-ecomm-utils/config"
)

func main() {
	host := config.Default().GetString("db.postgres.host")
	port := config.Default().GetInt("db.postgres.port")
	userName := config.Default().GetString("db.postgres.username")
	password := config.Default().GetString("db.postgres.password")
	dbname := config.Default().GetString("db.postgres.database")
	fmt.Println("1 : ", host)
	fmt.Println("2 : ", port)
	fmt.Println("3 : ", userName)
	fmt.Println("4 : ", password)
	fmt.Println("5 : ", dbname)

	//db, err := OpenDb()
	//if err != nil {
	//	fmt.Printf("Error while connecting  database: %s", err)
	//}
	//db.Database.Ping()
	//db.Database.Close()
}

//func OpenDb() (*database.OrmDB, error) {
//	orm, err := database.OpenORM()
//	if err != nil {
//		return nil, err
//	}
//	return orm, nil
//}
