// module github.com/yuta519/golang-sandbox
module golang_sandbox

go 1.16

require (
	github.com/anaskhan96/soup v1.2.4
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/yuta519/golang-sandbox/scripts/crawler/pkg => ./pkg
