
SRC = main.go unidep.go xpath.go fixpunct.go

%.go : %-in.go compile
	rm -f $@
	./compile $< | gofmt > $@ || rm $@
	chmod 444 $@

all: unidep

unidep: $(SRC)
	go build -o $@ $^

compile: compile.go
	go build $^
