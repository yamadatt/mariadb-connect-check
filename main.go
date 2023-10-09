// check_mariadb_ver.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8091", nil))
	fmt.Println("Hello World")

}

func handler(writer http.ResponseWriter, _ *http.Request) {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "mysqluser:mysqlpassword@tcp(x.x.x.x:3306)/mysql")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}

	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)

	message := fmt.Sprint("<h1>Sucess!! Connected to:", version, "</h1>")

	fmt.Fprint(writer, message)

}
