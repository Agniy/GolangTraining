module example.com/usequote

go 1.21

// Блок require появляется автоматически после `go get` или `go mod tidy`.
// Здесь мы зависим от классического учебного пакета rsc.io/quote.
require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
