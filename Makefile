tidy:
	go mod tidy

clean:
	rm -f goodies

compile: clean
	go build -o goodies main.go

env:
	cp env.sample .env
