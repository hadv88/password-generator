BINARY_NAME=password-generator
build:
	go build -o $(BINARY_NAME)
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe || true
test:
	go test ./...
run:
	./$(BINARY_NAME) --length=12 --special
release:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)