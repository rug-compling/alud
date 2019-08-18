
SRC = \
	attributes.go \
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
	go build -tags compile $^

attribmake: attribmake.go
	go build -tags attrib $^

attributes.go: attributes.txt attribmake
	rm -f $@
	./attribmake $< > $@
	chmod 444 $@
