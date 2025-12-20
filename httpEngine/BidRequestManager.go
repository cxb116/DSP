package httpEngine

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func BidRequestManager(w http.ResponseWriter, req *http.Request) {
	fmt.Println("BidRequest")

	var reader io.Reader = req.Body

	if req.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(req.Body)
		if err != nil {

			http.Error(w, "invalid gzip request body", http.StatusBadRequest)
			return
		}
		defer gz.Close()
		reader = gz
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, "read body failed", http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))

	// 正常返回
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"code":0,"msg":"ok"}`))
}
