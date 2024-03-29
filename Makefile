BINARY_NAME=celeritasApp.exe

## build: builds all binaries
build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo Celeritas built!

run: build
	@echo Staring Celeritas...
	@.\tmp\${BINARY_NAME} &
	@echo Celeritas started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

start: run
    
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Celeritas

restart: stop start
