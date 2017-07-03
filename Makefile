CMD_APISERVER=github.com/danielsussa/simpleCRUD/backend
CMD_TEST=github.com/danielsussa/simpleCRUD/test

builds:
	echo "up to date!"
	go get -v -d ${CMD_APISERVER}
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/simpleCRUD ${CMD_APISERVER}
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/smoke-test ${CMD_TEST}

docker:
	docker build -t simple-crud -f Dockerfile .

remove:
	docker rmi simple-crud -f

stop:
	docker ps -q --filter ancestor="simple-crud" | xargs -r docker stop

tests:
	echo "Testing..."
	go get github.com/danielsussa/go-task-runner
	go-task-runner
