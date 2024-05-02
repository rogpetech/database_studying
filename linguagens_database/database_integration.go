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
		queryDrop := "DROP DATABASE IF EXISTS " + dbName

		db.Exec(queryDrop)

		var query string = "CREATE DATABASE "
		fmt.Println("")
		fmt.Printf("Database %s does not exist. Creating...\n", dbName)
		fmt.Println("")
		if _, err := db.Exec(query + dbName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("")
		fmt.Println("==================")
		fmt.Println("Database created.")
		fmt.Println("==================")
		fmt.Println("")
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
			readData(db)
		case 4:
			updateData(db)
		case 5:
			deleteData(db)
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter columns (e.g, id, name):")
	columns, _ := reader.ReadString('\n')
	columns = strings.TrimSpace(columns)

	fmt.Println("Enter values (e.g., 1, 'Jhon Doe'):")
	values, _ := reader.ReadString('\n')
	values = strings.TrimSpace(values)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, values)

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted.")
}

func readData(db *sql.DB) {
	fmt.Println("Enter table name:")
	var tableName string

	fmt.Scanln(&tableName)

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	columns, err := rows.Columns()

	if err != nil {
		log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, column := range values {
		fmt.Println()
		fmt.Printf("%s: %s\n", columns[i], string(column))
		fmt.Println()
	}
}

func updateData(db *sql.DB) {
	fmt.Println("Enter table name:")
	var tableName string
	fmt.Scanln(&tableName)

	fmt.Println("Enter SET clause (e.g., name = 'Carlos tester' where id=1)")
	reader := bufio.NewReader(os.Stdin)
	setClause, err := reader.ReadString('\n')
	setClause = strings.TrimSpace(setClause)

	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("UPDATE %s SET %s", tableName, setClause)

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("Data updated.")
	fmt.Println("")

}

func deleteData(db *sql.DB) {
	fmt.Println("Enter table name:")
	var tableName string
	fmt.Scanln(&tableName)

	fmt.Println("Enter condition (e.g., id=1)")
	reader := bufio.NewReader(os.Stdin)
	condition, err := reader.ReadString('\n')
	condition = strings.TrimSpace(condition)

	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, condition)
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("Data deleted.")
	fmt.Println("")

}
