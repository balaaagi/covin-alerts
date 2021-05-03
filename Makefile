clean:
	@echo "Getting Dependencies..."
	go get -d ./...

osx_build: clean
	@echo "Building "
	go build -o bin/main-osx ./main/main.go

freebsd_build: clean
	@echo "Building "
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 ./main/main.go

linux_build: clean
	@echo "Building "
 	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./main/main.go

windows_build: clean
	@echo "Building "
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

run:
	go run main.go