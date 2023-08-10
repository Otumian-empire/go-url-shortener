.PHONY: lampp klampp app tidy

lampp:
	@sudo /opt/lampp/lampp start

klampp:
	@sudo /opt/lampp/lampp stop

app:
	@clear; go run cmd/go-url-shortener/main.go 

tidy:
	@go mod tidy