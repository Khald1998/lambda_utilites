#test the connection of mysql

"go get -u github.com/go-sql-driver/mysql"
"go get -u github.com/aws/aws-lambda-go"
$Env:GOOS = "linux"
go build main.go