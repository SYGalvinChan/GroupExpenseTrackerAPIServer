apiserver:
	go build -o build/apiserver cmd/main.go

clean:
	rm -f build/*