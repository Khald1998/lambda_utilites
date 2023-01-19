package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

var (
	username = os.Getenv("rds_username")
	password = os.Getenv("rds_password")
	endpoint = os.Getenv("rds_endpoint")
	port     = os.Getenv("rds_port")
	db_name  = os.Getenv("rds_db_name")
)

// var db *sql.DB

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) {

	var open_state string = "Success"
	var ping_state string = "Success"

	argment := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, endpoint, port, db_name)
	connection, err := sql.Open("mysql", argment)

	if err != nil {
		open_state = "Not A Success!"
		log.Fatalf("impossible to create the connection: %s", err)
		os.Exit(1)
	} else {
		log.Printf("Open Connection: %s!", open_state)
		log.Printf("Connection with %s", argment)
		log.Println("")
	}

	err = connection.Ping()

	if err != nil {
		ping_state = "Not A Success!"
		log.Fatalf("impossible to Ping: %s", err)
		os.Exit(1)
	} else {
		log.Printf("Ping: %s!", ping_state)
	}
	// if no error. Ping is successful
	log.Println("Ping to database successful, connection is still alive")

	connection.Close()

}

//$Env:GOOS = "linux"
//go build main.go
