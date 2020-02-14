build:
	@docker build -t c1 .
run:
	@docker run --rm --name c1 c1

.PHONY: build run
