# Advent Code Challenge

https://adventofcode.com/ 2019

## Run tests for all packages prefixed with day
`go test -v -run '' $(ls ./ | grep "day" | xargs -I {} -n 1 echo './{}')`
