// gvm instal go1.19.2
// gvm use go1.19.2 --default
// go get -u github.com/go-sql-driver/mysql

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Criar uma interface para CRUD e select se comunicando com o banco de dados Go => Database mysql
// Via CLI COMMAND-LINE eu posso cadastrar um database, tabelas, selecionar registros

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Print("Enter database name: ")
	var dbName string
	fmt.Scanln(&dbName)

	if !checkDatabase(db, dbName) {
		var query string = "CREATE DATABASE "
		fmt.Printf("Database %s does not exist. Creating...\n", dbName)
		if _, err := db.Exec(query + dbName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database created.")
	}

	db.Exec("USE " + dbName)

	for {
		fmt.Println("1. Create table")
		fmt.Println("2. Insert data")
		fmt.Println("3. Read data")
		fmt.Println("4. Update data")
		fmt.Println("5. Delete data")
		fmt.Println("6. Exit")
		fmt.Println("Choose an option: ")
		var option int

		fmt.Scanln(&option)

		switch option {
		case 1:
			createTable(db)
		case 2:
			insertData(db)
		case 3:
			// readData(db)
		case 4:
			// updateData(db)
		case 5:
			// deleteData(db)
		case 6:
			fmt.Println("Existing program")
			os.Exit(0)
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func checkDatabase(db *sql.DB, dbName string) bool {
	var temp string
	query := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME " + dbName
	err := db.QueryRow(query).Scan(&temp)
	if err != nil {
		return false
	}

	return false
}

func createTable(db *sql.DB) {
	fmt.Printf("Enter table name:")
	var tableName string
	fmt.Scanln(&tableName)

	fmt.Println("Enter columns and types (e.g., id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255)):")
	reader := bufio.NewReader(os.Stdin)
	columnDetails, _ := reader.ReadString('\n')
	columnDetails = strings.TrimSpace(columnDetails)

	query := fmt.Sprintf("CREATE TABLE %s (%s) ", tableName, columnDetails)
	fmt.Println(query)
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created.")
}

func insertData(db *sql.DB) {
	fmt.Println("Enter table name ")
	var tableName string
	fmt.Scanln(&tableName)

	var columns string
	fmt.Println("Enter columns (e.g, id, name):")
	fmt.Scanln(&columns)

	var values string
	fmt.Println("Enter values (e.g., 1, 'Jhon Doe'):")
	fmt.Scanln(&values)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, values)
	fmt.Println("Query: %s", query)
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted.")
}
