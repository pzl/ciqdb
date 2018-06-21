SRCS=prg.go sections.go settings.go head.go entries.go data.go pc2ln.go perms.go symboltable.go except.go devkey.go linktable.go apidb.go symboltable.go

all: parse

parse: $(SRCS)
	go build

run: parse
	./parse samples/crystal-face.prg

bench: prof.png
	viewnior $<

prof.png: cpu.prof
	go tool pprof -png parse.test $< > $@

cpu.prof: parse_test.go $(SRCS)
	go test -bench . -cpuprofile=$@

clean:
	$(RM) -rf parse prof.png cpu.prof parse.test