build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/simpleCRUD .
	docker build -t simple-crud -f Dockerfile .