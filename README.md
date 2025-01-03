# AoC 2024

[![Go Tests](https://github.com/kaveenr/aoc24/actions/workflows/tests.yml/badge.svg)](https://github.com/kaveenr/aoc24/actions/workflows/tests.yml)
[![Go Report](https://goreportcard.com/badge/github.com/kaveenr/aoc24)](https://goreportcard.com/report/github.com/kaveenr/aoc24)

## Prerequisites

1. GNU + Darwin/Linux enviroment
2. Go installed
3. Curl / Pandoc installed for scraping (optional)
4. Make `.env` file from `.env.sample` (optional)

## Instructions

```sh
$ make
Available targets:
- install: Install dependencies
- run day=<day_number>: Run the code for a specific day
- test: Run tests
- bench: Run benchmarks
- new year=<year> day=<day_number>: Create from template and scrape
- scrape year=<year> day=<day_number>: Scrape puzzle and input from site
- scrape-all year=<year> day=<day_number>: Scrape from day 1 to specified
- answer year=<year> day=<day_number> part=<part> answer=<answer>: Answer puzzle

Hints: by default <year>,<day> will be set to current EST date
```
