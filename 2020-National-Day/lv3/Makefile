main: main.go
	go build main.go

.PHONY: test
test:
	make
	printf "%d\n%d %d %d %d %d\n" 5 1 3 2 5 4 | ./main

.PHONY: clean
clean:
	rm main
