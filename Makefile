build:
	go build -o bin/ssl_checker cmd/ssl-checker/main.go
clean:
	rm -f bin/ssl_checker

all: clean build
