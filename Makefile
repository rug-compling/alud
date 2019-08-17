
SRC = \
	alpino.go \
	alud.go \
	attributes.go \
	auxiliary.go \
	conllu.go \
	depheads.go \
	deplabels.go \
	deprels.go \
	doc.go \
	emptyheads.go \
	enhanced.go \
	features.go \
	fixpunct.go \
	helpers.go \
	misplacedheads.go \
	postags.go \
	types.go \
	untokenize.go \
	xpath.go

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
	./attribmake $< > $@
