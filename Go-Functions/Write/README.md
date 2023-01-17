
# go_in_lambda
lambda function with golang


## make sure you remamber to use those
"go mod init Write" (inside the write folder)

"go run .\main.go" (inside the write folder) or "go run .\Go-Functions\Write\main.go" 

"go get -u github.com/aws/aws-lambda-go" (inside the write folder)

before upload we need to build the code into compiled binary for aws and zip it

$Env:GOOS = "linux"

go build main.go 

