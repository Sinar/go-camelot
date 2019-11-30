run: build
	@./go-camelot

build:
	@CGO_ENABLED=1 go build .

tools:
	# Assumes below are on the system; not as clean, but what can you do :(
	@brew install opencv tesseract
