
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
	go build -tags compile $^

demo/demo: demo/demo.go $(SRC) *.go
	cd demo && go build -v .

demo-dact/demo-dact: demo-dact/demo-dact.go $(SRC) *.go
	cd demo-dact && go build -v .

demos: demo/demo demo-dact/demo-dact
