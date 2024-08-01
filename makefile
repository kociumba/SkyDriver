all: b

b: generate
	go mod tidy
	go build -o ./build

r: generate
	@if [ -z "$(system)" ]; then \
		echo "No system specified"; \
		exit 1; \
	fi; \
	out=build/SkyDriver-$(system); \
	[ "$(system)" = "windows" ] && out=$${out}.exe; \
	go build -o $${out} -ldflags="-s -w"

# pass limit when executing from make couse I'm lazy
run: b
	./build/skyDriver $(if $(limit),-limit $(limit),) $(if $(sell),-sell $(sell),) $(if $(dbg),-dbg,) $(if $(search), -search $(search),) $(if $(skip),-skip,) $(if $(max), -max $(max),)

generate:
	go run github.com/tc-hib/go-winres@latest make
