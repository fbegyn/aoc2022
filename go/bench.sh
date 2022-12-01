#!/usr/bin/env bash

days=()

for file in ./bin/*; do
    day=$(basename $file)
    days+=("${file} ../inputs/${day}/input.txt")
done

echo $days

hyperfine --warmup 500 --export-csv aoc2022.csv "${days[@]}"
