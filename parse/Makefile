SRCS=prg.go sections.go settings.go head.go entries.go data.go pc2ln.go perms.go

all: parse

parse: $(SRCS)
	go build

run: parse
	./parse ~/dev/MonkeyC/samples/crystal-face/bin/crystal-face.prg

bench: prof.png
	viewnior $<

prof.png: cpu.prof
	go tool pprof -png parse.test $< > $@

cpu.prof: parse_test.go $(SRCS)
	go test -bench . -cpuprofile=$@

clean:
	$(RM) -rf parse prof.png cpu.prof parse.test