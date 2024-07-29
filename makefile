# https://www.gnu.org/software/make/manual/html_node/Syntax-of-Makefiles.html

all: b

b:
	go mod tidy
	go build -o ./build

r:
	go mod tidy
	go build -o ./build -ldflags="-s -w"

run: b
	./build/skyDriver

