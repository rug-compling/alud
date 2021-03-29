
all:
	@echo
	@echo Run make with one of these targets:
	@echo
	@echo "  bin       : install alud and alud-dact"
	@echo "  alud      : install alud"
	@echo "  alud-dact : install alud-dact"
	@echo
	@echo '  devel     : only for maintainers'
	@echo

bin: alud alud-dact

alud:
	cd v2/cmd/alud && go install -v .

alud-dact:
	cd v2/cmd/alud-dact && go install -v .

devel:
	make -C v2
