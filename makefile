init:
	yarn install
	go mod download
	cp ./.air.toml.example ./.air.toml
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css

build:
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css
	go build -o tmp/main cmd/main.go

run:
	air

tailwind:
	npx tailwindcss -i ./web/static/css/index.css -o ./web/static/css/dist/output.css --watch

tidy:
	go mod tidy
