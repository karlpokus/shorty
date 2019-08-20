package shorty

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func newRouter(addr string) http.Handler {
	db := make(database)
	router := httprouter.New()
	router.Handler("POST", "/", logRequest(create(db, addr)))
	router.Handler("GET", "/:url", logRequest(follow(db)))
	router.Handler("GET", "/", logRequest(list(db)))
	return router
}

func logRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Stdout.Print(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	}
}

func create(store Store, addr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Stderr.Printf("Malformed post body: %s", err)
			http.Error(w, "Bad Request: Malformed post body", 400)
			return
		}
		defer r.Body.Close()
		key, err := store.Add(string(body))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, "%s\n", fmt.Sprintf("%s/%s", addr, key))
	}
}

func follow(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		url, err := store.Find(params[0].Value)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		http.Redirect(w, r, url, 302)
	}
}

func list(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", formatList(store.List()))
	}
}

func formatList(db map[string]string) string {
	var out []string
	for k, v := range db {
		out = append(out, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(out, "\n")
}
