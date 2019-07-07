
SRC = \
	auxiliary.go \
	conllu.go \
	depheads.go \
	deplabels.go \
	deprels.go \
	document.go \
	emptyheads.go \
	enhanced.go \
	features.go \
	fixpunct.go \
	helpers.go \
	main.go \
	misplacedheads.go \
	postags.go \
	xpath.go \
	untokenize.go \



%.go : %-in.go compile
	rm -f $@
	./compile $< | gofmt > $@ || rm $@
	chmod 444 $@

all: unidep

unidep: $(SRC)
	go build -o $@ $^

compile: compile.go
	go build -tags compile $^
