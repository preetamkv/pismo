# Generate the executible
go build ./...
go build ./cmd/service/

# Generate Docker image
docker build -t pismo:latest .
clear

# Run the container using the image created
docker run -p 8080:8080 pismo:latest