all: format compile

format: ftia.go
	gofmt -w ftia.go

compile: ftia.go
	go build -o ftia ftia.go

clean:
	rm -f ftia

