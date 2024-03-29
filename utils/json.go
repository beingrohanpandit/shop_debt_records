package utils

import (
	"io"

	"github.com/IBM/fp-go/function"
	E "github.com/IBM/fp-go/ioeither"
	J "github.com/IBM/fp-go/json"
)

func ParseJSON[T any](r io.ReadCloser) E.IOEither[error, T] {
	return function.Pipe1(E.TryCatchError(func() ([]byte, error) {
		return io.ReadAll(r)
	}), E.ChainEitherK(J.Unmarshal[T]))
}
