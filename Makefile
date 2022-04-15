SHELL = C:\Program Files\Git\bin\bash

all:
	make windows
	make linux
	make image

windows:
	mkdir -p bin
	go build -o bin/syborg-win64.exe main.go

linux:
	docker build . -f Dockerfile.alpinebuild -t syborg-build:v1.0
	MSYS_NO_PATHCONV=1 docker run --rm -v `pwd`:/usr/src/syborg -w /usr/src/syborg syborg-build:v1.0 go build -o bin/syborg-linux-x64 main.go

image:
	make linux
	docker build -f Dockerfile -t syborg:v$(git tag | sort -V | tail -1)-dirty

clean:
	rm -rf bin