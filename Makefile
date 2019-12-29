run: build
	@./gocamelot

build:
	@CGO_ENABLED=1 go build -o gocamelot .

tools:
	# Assumes below are on the system; not as clean, but what can you do :(
	@brew install opencv tesseract
