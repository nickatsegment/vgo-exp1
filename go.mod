module github.com/nickatsegment/vgo-exp1/v2

go 1.13

require (
	github.com/nickatsegment/vgo-exp1/lib/v2 v2.0.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/nickatsegment/vgo-exp1/lib/v2 v2.0.0 => github.com/nickatsegment/vgo-exp1/lib/v2 v2.0.0-rc1.1

replace github.com/nickatsegment/vgo-exp1/lib/v2 => ./lib
