CMD_APISERVER=github.com/danielsussa/simpleCRUD/backend

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/simpleCRUD ${CMD_APISERVER}
	docker build -t simple-crud -f Dockerfile .