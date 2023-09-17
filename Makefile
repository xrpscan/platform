build:
	go build -o platform main.go 
	go build -o backfill cmd/backfill/backfill.go

clean:
	go clean
	rm backfill
