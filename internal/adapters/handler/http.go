package handler

import (
	"net/http"

	"github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/io"
	J "github.com/IBM/fp-go/json"
	"github.com/IBM/fp-go/option"
)

type HTTPResult struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func WriteLeft(w http.ResponseWriter) func(e error) io.IO[int] {
	return func(e error) io.IO[int] {
		return func() int {
			err := either.GetOrElse(func(_ error) []byte {
				return []byte("Error marshalling message")
			})(J.Marshal(HTTPResult{
				StatusCode: http.StatusBadRequest,
				Error:      e.Error(),
			}))

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(err)

			return 0
		}
	}
}

func WriteRight[T any](status int, w http.ResponseWriter, message option.Option[string]) func(v T) io.IO[int] {
	return func(v T) io.IO[int] {
		return func() int {
			w.WriteHeader(status)
			if option.IsSome(message) {
				m, _ := option.Unwrap(message)
				w.Write([]byte(m))
			} else {
				res := either.GetOrElse(func(_ error) []byte {
					return []byte("Error marshalling message")
				})(J.Marshal(v))

				w.Write(res)
			}
			return 0
		}
	}
}
