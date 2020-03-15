all: format compile

format: ftia.go
	gofmt -w ftia.go

compile: ftia.go
	go build -o ftia ./...

install: all
	sudo cp ftia /usr/local/bin/
	cp -r .ftia ~/
	[ -f ~/.ftia/known.txt ] || touch ~/.ftia/known.txt
	[ -f ~/.ftia/known_rev.txt ] || touch ~/.ftia/known_rev.txt

uninstall:
	sudo rm /usr/local/bin/ftia
	rm -rf ~/.ftia

clean:
	rm -f ./ftia
