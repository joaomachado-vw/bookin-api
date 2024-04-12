build:
	mkdir out
	go build -o out/booking-api

run:
	go run main.go

test:
	go test -v ./...

coverage:
	go test --cover -covermode=count .

clean:
	rm -Rf out
