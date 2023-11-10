include .env
export 

all: build run

clean:
	echo "Clen bin path"
	rm -rf bin/

build:
	echo "Build Application"
	go build -tags dev -o bin/cms cmd/api/main.go

run:
	echo "Run Application"
	./bin/cms

compile:
	echo "Compiling for every OS and Platform"
	CGO_ENABLED=0 GOOS=linux go build -o bin/cms cmd/api/main.go
	CGO_ENABLED=0 GOOS=windows go build -o bin/cms.exe cmd/api/main.go

dev:
	air

lint:
	golangci-lint run
