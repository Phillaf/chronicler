module main.go

go 1.15

replace github.com/phillaf/chronicler/finnhub => ./finnhub

replace github.com/phillaf/chronicler/normalizer => ./normalizer

replace github.com/phillaf/chronicler/scalespace => ./scalespace

replace github.com/phillaf/chronicler/pyramid => ./pyramid

replace github.com/phillaf/chronicler/extrema => ./extrema

require (
	github.com/phillaf/chronicler/extrema v0.0.0-00010101000000-000000000000
	github.com/phillaf/chronicler/finnhub v0.0.0-00010101000000-000000000000
	github.com/phillaf/chronicler/normalizer v0.0.0-00010101000000-000000000000
	github.com/phillaf/chronicler/pyramid v0.0.0-00010101000000-000000000000 // indirect
	github.com/phillaf/chronicler/scalespace v0.0.0-00010101000000-000000000000
)
