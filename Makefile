
all: compile v2make

compile: compile.go
	go build -tags compile $^

v2make:
	make -C v2
