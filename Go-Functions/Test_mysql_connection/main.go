package main

import (
	"database/sql"
	"fmt"
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
var db *sql.DB

// func init() {
// 	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, endpoint, port, db_name)
// 	var err error
// 	db, err = sql.Open("mysql", args)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// }

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, endpoint, port, db_name)
	db, err := sql.Open("mysql", conn)

	if err != nil {
		fmt.Println("Not A Success!")
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}

//$Env:GOOS = "linux"
//go build main.go
