package encdec

import (
	"VinylShop/pkg/web/errors"
	"fmt"
	"net/http"
)

func SendErr(w http.ResponseWriter, error *errors.Error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(error.Code())
	w.Write([]byte(fmt.Sprintf("%d %s", error.Code(), error.Error())))
}
