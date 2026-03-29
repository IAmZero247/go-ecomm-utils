package main

import (
	"fmt"

	"github.com/IAmZero247/go-ecomm-utils/config"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
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
}
