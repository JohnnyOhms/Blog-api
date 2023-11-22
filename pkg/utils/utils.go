package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, i interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		log.Fatal(json.Unmarshal([]byte(body), i))
	}
}
