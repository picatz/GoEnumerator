
make:
	go build

darwin:
	GOOS=darwin go build -o GoEnumerator.Darwin

windows:
	GOOS=windows go build -o GoEnumerator.Windows

clean:
	go clean
