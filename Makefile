CMD_APISERVER=github.com/danielsussa/simpleCRUD/backend

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/simpleCRUD ${CMD_APISERVER}
	docker build -t simple-crud -f Dockerfile .

docker:
	docker build -t simple-crud -f Dockerfile .

remove:
	docker rmi simple-crud -f

stop:
	docker ps -q --filter ancestor="simple-crud" | xargs -r docker stop

run:
	echo "Running APP!"
	docker ps -q --filter ancestor="simple-crud" | xargs -r docker stop
	docker run -d -p 8080:8081 simple-crud
	curl localhost:8080
	docker ps -q --filter ancestor="simple-crud" | xargs -r docker stop
