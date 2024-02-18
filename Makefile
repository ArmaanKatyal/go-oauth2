run: build
	./target/go-oauth

build:
	go build -o ./target/go-oauth

clean:
	rm -f ./target/go-oauth
