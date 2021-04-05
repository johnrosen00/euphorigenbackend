export GOOS=linux
go clean
go build
docker build -t johnrosen00/eugateway .
go clean