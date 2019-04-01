all: format compile

format: ftia.go
	gofmt -w ftia.go

compile: ftia.go
	go build -o ftia ftia.go

install: all
	sudo cp ftia /usr/local/bin/
	cp -r .ftia ~/
	touch ~/.ftia/known.txt

uninstall:
	sudo rm /usr/local/bin/ftia
	rm -rf ~/.ftia

clean:
	rm -f ./ftia
