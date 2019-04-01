all: format compile

format: ftia.go
	gofmt -w ftia.go

compile: ftia.go
	go build -o ftia ftia.go

install: all
	cp ftia /usr/local/bin/

uninstall:
	rm /usr/local/bin/ftia

clean:
	rm -f ./ftia
