package dagger

import (
	"encoding/json"
	"net/http"
)

func addJSONHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

// JSONContent is an Adapter that adds JSON headers to the
// response writer, making sure that the response is actual JSON.
func JSONContent() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			addJSONHeaders(w)
			h.ServeHTTP(w, r)
		})
	}
}

// JSON writes a generic interface type to JSON
func JSON(w http.ResponseWriter, body interface{}) (int, error) {
	data, err := json.Marshal(body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0, err
	}

	written, err := w.Write(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0, err
	}

	addJSONHeaders(w)

	return written, nil
}
