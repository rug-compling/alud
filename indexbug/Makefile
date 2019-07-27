
SRC = \
	main.go \
	xpath.go \



%.go : %-in.go ../compile
	rm -f $@
	../compile $< | gofmt > $@ || rm $@
	chmod 444 $@

all: main

main: $(SRC)
	go build -o $@ $^
