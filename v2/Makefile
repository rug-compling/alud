
SRC = \
	auxiliary.go \
	depheads.go \
	deplabels.go \
	emptyheads.go \
	enhanced.go \
	misplacedheads.go

%.go : %-in.go ../compile
	rm -f $@ $@.tmp.go
	../compile $< > $@.tmp.go
	gofmt -w $@.tmp.go || ( chmod 444 $@.tmp.go && false )
	mv $@.tmp.go $@
	chmod 444 $@

all: $(SRC)

