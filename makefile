all: b

b: generate
	go mod tidy
	go build -o ./build

r: generate
	go mod tidy
	go build -o ./build -ldflags="-s -w"

# pass limit when executing from make couse I'm lazy
run: b
	@if [ -n "$(limit)" ]; then \
		./build/skyDriver -limit $(limit); \
	else \
		./build/skyDriver; \
	fi

generate:
	go run github.com/tc-hib/go-winres@latest make
