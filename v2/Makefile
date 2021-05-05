
SRC = \
	auxiliary.go \
	depheads.go \
	deplabels.go \
	emptyheads.go \
	enhanced.go \
	misplacedheads.go

%.go : %-in.go compile
	rm -f $@ $@.tmp.go
	./compile $< > $@.tmp.go
	gofmt -w $@.tmp.go || ( chmod 444 $@.tmp.go && false )
	mv $@.tmp.go $@
	chmod 444 $@

all: $(SRC)

compile: compile.go
	@if [ "`testXPath --alud`" != "ALUD" ]; then echo INVALID VERSION OF testXPath; exit 1; fi
	go build -tags compile $^
