
all: unidep

unidep: main.go unidep.go xpath.go fixpunct.go
	go build -o unidep main.go unidep.go xpath.go fixpunct.go

unidep.go: unidep-in.go compile
	rm -f unidep.go
	./compile unidep-in.go | gofmt > unidep.go || rm unidep.go
	chmod 444 unidep.go

compile: compile.go
	go build compile.go
