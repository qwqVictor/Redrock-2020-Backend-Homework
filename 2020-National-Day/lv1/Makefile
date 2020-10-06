main: main.go
	go build main.go

.PHONY: test
test:
	make
	echo '3 + 2' | ./main
	echo '3 - 2' | ./main
	echo '3 * 2' | ./main
	echo '3 / 2' | ./main
	echo '10 / 2' | ./main
	echo '1 / 0' | ./main


.PHONY: clean
clean:
	rm main
