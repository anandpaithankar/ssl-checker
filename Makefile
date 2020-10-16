build:
	go build -o bin/ssl_checker src/main.go
clean:
	rm -f bin/ssl_checker

all: clean build