module hello

go 1.15

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.9.0
	github.com/koding/logging v0.0.0-20160720134017-8b5a689ed69b
)
