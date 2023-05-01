init:
	yarn install
	go mod download
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css

build:
	echo "Building..."
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css
	go build -o tmp/main cmd/main.go

run:
	air

tailwind:
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css --watch

tidy:
	go mod tidy
