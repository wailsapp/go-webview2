module updater

replace generator => ./generator

go 1.24.4

require (
	generator v0.0.0
	github.com/go-resty/resty/v2 v2.7.0
)

require golang.org/x/net v0.0.0-20211030010937-7b24c0a3601d // indirect
