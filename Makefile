osx_build:
	go build -o bin/main-osx ./main/main.go

freebsd_build:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 ./main/main.go

linux_build:
 	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./main/main.go

windows_build:
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

run:
	go run main.go