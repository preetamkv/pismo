go build ./...
go build ./cmd/service/

docker build -t pismo:latest .
docker run -p 8080:8080 pismo:latest