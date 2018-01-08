BIN := $(shell basename $(CURDIR))

.PHONY: all clean clobber frontend test

all: clean $(BIN) frontend test run

clean:
	rm -f $(BIN)
	rm -fr build

clobber: clean
	rm -fr node_modules

node_modules:
	npm i

frontend: node_modules
	webpack

$(BIN): 
	go build .

test: $(BIN)
	go test -v .

run: $(BIN)
	./$(BIN)

