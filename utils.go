package dagger

import "net/http"

// Adapter is a function that takes a http.Handler and returns another
type Adapter func(http.Handler) http.Handler

// Adapt takes a handler and wraps multiple adapters around it in reverse order.
// Inspired by https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
func Adapt(adapters ...Adapter) Adapter {
	return func(h http.Handler) http.Handler {
		for _, adapter := range adapters {
			h = adapter(h)
		}
		return h
	}
}
