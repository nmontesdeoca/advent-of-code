# Advent of Code

Just my AoC solutions — each day is its own Go module in `2025/day-X`. Nothing fancy, keeps them isolated from each other.

## How it's set up

- Each day gets its own folder with `go.mod`, `main.go`, `main_test.go`, and `input.txt`
- No sharing code between days — if you need a helper, just throw it in that day's `main.go`
- Read from `./input.txt` relative to the day folder (and there's usually `input_test.txt` for tinkering)

## Running stuff

From inside a day folder:

```bash
go run .
go build
go test -bench=. -benchmem
```

Benchmarks just measure how fast `part1` and `part2` run on the actual input.

## Adding a new day

1. Copy a day folder that looks similar to what you need
2. Update `go.mod` to `github.com/nmontesdeoca/advent-of-code/2025/day-X`
3. Adjust the parsing in `main.go` (newline, comma, whatever the puzzle uses)
4. Implement `part1` and `part2` (they return `int`)
5. Copy the same parsing into `main_test.go` and set up the benchmark loop

## 2025 Solutions

| Day                         | Part 1  | Part 2   | Notes                              |
| --------------------------- | ------- | -------- | ---------------------------------- |
| [Day 1](2025/day-1/main.go) | 25.4 µs | 291 µs   | String validation across ranges    |
| [Day 2](2025/day-2/main.go) | 35.3 ms | 148 ms   | Pattern detection in ID strings    |
| [Day 3](2025/day-3/main.go) | 45.4 µs | 128.0 µs | Grid traversal and neighbor checks |

Benchmarked on: 12th Gen Intel(R) Core(TM) i9-12900HX (benchmarks updated: 2025-12-31)

## More info

- Examples: [day 1](2025/day-1/main.go), [day 2](2025/day-2/main.go)
