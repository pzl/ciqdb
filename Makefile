SRCS=prg.go sections.go settings.go head.go entries.go data.go pc2ln.go perms.go symboltable.go except.go devkey.go linktable.go apidb.go symboltable.go

all: ciqdb

ciqdb: $(SRCS)
	go generate
	go build

run: ciqdb
	./ciqdb samples/crystal-face.prg

bench: prof.png
	viewnior $<

prof.png: cpu.prof
	go tool pprof -png ciqdb.test $< > $@

cpu.prof: parse_test.go $(SRCS)
	go test -bench . -cpuprofile=$@

clean:
	$(RM) -rf ciqdb prof.png cpu.prof ciqdb.test *_string.go