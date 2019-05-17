SRCS=$(wildcard *.go cmd/*/*.go)
TARGET=bin/ciqdb
BLDDIR=build


all: $(TARGET)

$(TARGET): $(SRCS)
	go generate ./...
	go build -o $@ ./cmd/ciqdb

run: $(TARGET)
	./bin/ciqdb samples/crystal-face.prg

bench: $(BLDDIR)/prof.png
	viewnior $<

$(BLDDIR)/prof.png: $(BLDDIR)/cpu.prof
	go tool pprof -png ciqdb.test $< > $@

$(BLDDIR)/cpu.prof: $(SRCS) $(BLDDIR)
	go test -bench . -cpuprofile=$@


$(BLDDIR):
	mkdir -p $@

clean:
	$(RM) -rf bin/* ciqdb.test $(BLDDIR)/* ciq/*_string.go