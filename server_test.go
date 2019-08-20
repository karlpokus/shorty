package shorty

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/karlpokus/shorty/internal/dbmock"
)

var (
	addr         = "localhost:9012"
	db           = make(dbmock.StoreMock)
	shortUrlPath = "/abc123"
)

var testTable = []struct {
	name, method, registerPath, visitPath string
	requestBody                           io.Reader
	fn                                    http.HandlerFunc
	status                                int
	responseBody                          []byte
}{
	{
		"create short url",
		"POST",
		"/",
		"/",
		bytes.NewReader([]byte("www.google.com")),
		create(db, addr),
		200,
		[]byte(addr + shortUrlPath),
	},
	{
		"redirect to long url",
		"GET",
		"/:url",
		shortUrlPath,
		nil,
		follow(db),
		302,
		[]byte(`<a href="/www.google.com">Found</a>.`),
	},
	{
		"list available short urls",
		"GET",
		"/",
		"/",
		nil,
		list(db),
		200,
		[]byte(""),
	},
}

func TestRoutes(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			router := httprouter.New()
			router.Handler(tt.method, tt.registerPath, tt.fn)
			r := httptest.NewRequest(tt.method, tt.visitPath, tt.requestBody)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)

			if res.StatusCode != tt.status {
				t.Errorf("expected %d, got %d", tt.status, res.StatusCode)
			}
			if !bytes.Equal(bytes.TrimSpace(body), tt.responseBody) {
				t.Errorf("expected %s, got %s", tt.responseBody, bytes.TrimSpace(body))
			}
		})
	}
}
