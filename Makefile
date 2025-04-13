default: clean gosix test image

clean:
	rm -f gosix

gosix:
	GOOS=linux go build -o gosix main.go

image: gosix
	docker build -t mackrorysd/gosix .

test:
	go test -cover \
		github.com/mackrorysd/gosix/utilities \
		github.com/mackrorysd/gosix/shell \
		github.com/mackrorysd/gosix

run:
	docker run -it mackrorysd/gosix
