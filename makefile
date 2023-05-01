build:
	echo "Building..."
	yarn
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css
	go build -o tmp/main cmd/main.go

run:
	yarn
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css
	air

tidy:
	go mod tidy
