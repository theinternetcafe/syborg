all:
	mkdir -p bin	
	go build -o bin/main main.go

clean:
	rm -rf bin