GC=go

all:
	mkdir -p bin
	$(GC) build -o bin/bootdata bootdata.go

install:
	sudo cp bin/bootdata /usr/local/bin/
