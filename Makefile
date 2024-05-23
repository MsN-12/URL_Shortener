run:
	@go run main.go
css:
	@tailwindcss -i ./view/css/app.css -o ./public/styles.css --watch
templ:
	@templ generate --watch --proxy=http://localhost:3000