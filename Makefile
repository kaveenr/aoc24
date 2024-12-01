.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day"
	@echo "- test: Run tests"
	@echo "- bench: Run benchmarks"
	@echo "- scrape year=<year> day=<day_number>: Scrape puzzle and input from site"
	@echo "- scrape-all year=<year> day=<day_number>: Scrape from day 1 to specified"
	@echo "- answer year=<year> day=<day_number> part=<part> answer=<answer>: Answer puzzle"
	@echo ""
	@echo "Hints: by default <year>,<day> will be set to current EST date"

.PHONY: install run test bench

install:
	go mod tidy

run:
	go run ./cmd/day$(day)

test:
	@go test  -cover $(shell go list ./...) -run=Example

bench:
	@go test -benchmem -run=^$ -cover -bench . $(shell go list ./...)

.PHONY: new answer scrape scrape-all
-include .env
year = $(shell env TZ=EST date +"%Y")
day = $(shell env TZ=EST date +"%-d")

answer:
	@curl -s -X 'POST' 'https://adventofcode.com/$(year)/day/$(day)/answer' \
		-b "session=${AOC_SESSION}" \
		-H 'Content-Type: application/x-www-form-urlencoded' \
		--data 'level=$(part)&answer=$(answer)' \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=plain
	@make scrape

scrape:
	@curl -s https://adventofcode.com/$(year)/day/$(day)/input \
		-b "session=${AOC_SESSION}" \
		> day$(day)/input.txt
	@curl -s https://adventofcode.com/$(year)/day/$(day) \
		-b "session=${AOC_SESSION}" \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=markdown \
		> day$(day)/puzzle.md
	@echo "Scraped Day $(day) for $(year)"

scrape-all:
	@for i in $$(seq 1 $(day)); do \
        make scrape day=$$i; \
    done