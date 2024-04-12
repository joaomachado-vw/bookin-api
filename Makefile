build:
	mkdir -p out
	go build -o out/booking-api

run:
	go run main.go

test:
	go test -v ./...

coverage:
	go test --cover -covermode=count .

clean:
	rm -Rf out

tools:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.1

check:
	staticcheck ./...


lint:
	golangci-lint run -v ./...