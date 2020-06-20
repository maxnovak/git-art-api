build-web:
	go build -o bin/generate-art src/cmd/restAPI/main.go

build-console:
	go build -o bin/generate-art src/cmd/console/main.go

run:
	./bin/generate-art