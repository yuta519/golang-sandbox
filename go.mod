// module github.com/yuta519/golang-sandbox
module golang_sandbox

go 1.17

require (
	github.com/anaskhan96/soup v1.2.4
	github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/text v0.3.0 // indirect
)

replace github.com/yuta519/golang-sandbox/scripts/crawler/pkg => ./pkg
