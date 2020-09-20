package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// HandleError agrega información al error del archivo y la linea donde marco el error.
func HandleError(err error) (reterr error) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(1)
		reterr = fmt.Errorf("[%s:%d] %v /", filepath.Base(fn), line, err)
	}
	return
}

// Testigo funcion para imprimir un mensaje del lugar donde se ponga
func Testigo() {
	_, fn, line, _ := runtime.Caller(1)
	fmt.Printf("\r\n%s:%d /\r\n", fn, line)
}
