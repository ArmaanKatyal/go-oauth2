run: build
	./target/go-oauth2

build:
	go build -o ./target/go-oauth2

clean:
	rm -f ./target/go-oauth2

test:
	go test -v ./...