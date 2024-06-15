package encdec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, entity interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(entity); err != nil {
		return fmt.Errorf("WriteJSON: %v", err)
	}
	return nil
}
func ReadJSON[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}
