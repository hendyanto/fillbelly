package main
	
import (
	"fmt"
	"os"
)

func DbConnectionString() string {
	db_pass := "S7D6F9S7S1"
	db_name := "go_database"
	if(os.Getenv("GO_ENV") == "test") {
		db_pass = "S7D6F9S7S1"
		db_name = "go_database_test"
	}
	connStr := fmt.Sprintf("user=go_datauser dbname=%s password='%s'", db_name, db_pass)
	return connStr
}