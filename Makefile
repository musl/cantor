BIN := $(shell basename $(CURDIR))

.PHONY: all clean clobber frontend test

all: clean test $(BIN) frontend run

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
	go build -o ./$(BIN) ./src

test: $(BIN)
	go test ./src 

run: $(BIN)
	./$(BIN)

