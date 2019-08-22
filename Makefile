
all: compile attribmake v2make

compile: compile.go
	go build -tags compile $^

attribmake: attribmake.go
	go build -tags attrib $^

v2make:
	make -C v2
