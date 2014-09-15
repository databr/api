package service

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type DataRender struct {
	r *http.Request
}

func (render DataRender) Render(w http.ResponseWriter, code int, data ...interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if render.r.URL.Query().Get("pretty") == "true" {
		out, err := json.MarshalIndent(data[0], "", "    ")
		if err != nil {
			return err
		}
		w.Write(out)
		return nil
	}

	writer := render.gzip(w)
	return json.NewEncoder(writer).Encode(data[0])
}

func (render DataRender) gzip(w http.ResponseWriter) io.Writer {
	if render.shouldGzipResonse() {
		gz := gzip.NewWriter(w)
		w.Header().Set("Content-Encoding", "gzip")
		defer gz.Close()
		return gz
	} else {
		return w
	}
}

func (render DataRender) shouldGzipResonse() bool {
	h := render.r.Header
	return ENV != "development" &&
		strings.Contains(h.Get("User-Agent"), "Mozilla") &&
		strings.Contains(h.Get("User-Agent"), "Opera") &&
		strings.Contains(h.Get("Accept-Encoding"), "gzip") &&
		render.r.URL.Query().Get("pretty") != "true"
}
