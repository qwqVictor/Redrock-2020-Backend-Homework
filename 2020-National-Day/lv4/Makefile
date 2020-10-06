main: main.go
	go mod tidy
	go build main.go

.PHONY: test
test: 
	./test/test.sh

.PHONY: clean
clean:
	rm main go.sum
