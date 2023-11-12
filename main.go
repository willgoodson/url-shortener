package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	var ctx = context.Background()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	r.Post("/new", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		url := r.Form.Get("url")
		if url == "" {
			panic("No URL")
		}

		unique_slug := uniuri.NewLen(8)
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		err := rdb.Set(ctx, unique_slug, url, 0).Err()
		rdb.Close()
		if err != nil {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		}
		w.Write([]byte("<a id='result' target='_blank' href='" + unique_slug + "'>" + "http://localhost:3000/" + unique_slug + "</a>"))
	})

	r.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		val, err := rdb.Get(ctx, slug).Result()
		rdb.Close()
		if err != nil {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		}

		http.Redirect(w, r, val, http.StatusPermanentRedirect)
	})
	fmt.Print("Starting Server...")
	http.ListenAndServe(":3000", r)
}
