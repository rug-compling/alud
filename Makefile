
SRC = \
	auxiliary.go \
	depheads.go \
	deplabels.go \
	emptyheads.go \
	enhanced.go \
	misplacedheads.go

%.go : %-in.go compile
	rm -f $@
	./compile $< | gofmt > $@ || rm $@
	chmod 444 $@

all: $(SRC)

compile: compile.go
	go build -tags compile $^
