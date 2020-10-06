main: main.go
	go build main.go

.PHONY: test
test:
	make
	printf "abracadabra\n" | ./main

.PHONY: clean
clean:
	rm main
