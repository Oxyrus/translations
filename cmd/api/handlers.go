package main

import (
	"github.com/go-chi/chi/v5"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"

	// Import the internal/translations package, so that its init()
	// function is called.
	_ "github.com/Oxyrus/translations/internal/translations"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	locale := chi.URLParam(r, "locale")

	var lang language.Tag

	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "es-co":
		lang = language.MustParse("es-CO")
	default:
		http.NotFound(w, r)
		return
	}

	p := message.NewPrinter(lang)
	p.Fprintf(w, "Welcome!\n")

	books := 1
	p.Fprintf(w, "%d books available\n", books)
}
