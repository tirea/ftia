all: format compile

format: ftia.go
	gofmt -w ftia.go

compile: ftia.go
	go build -o ftia ftia.go

install: all
	cp ftia /usr/local/bin/
	cp -r .ftia ~/

uninstall:
	rm /usr/local/bin/ftia
	rm -rf ~/.ftia

clean:
	rm -f ./ftia
