# gomponents-htmx

<img src="logo.png" alt="Logo" width="300" align="right">

[![GoDoc](https://pkg.go.dev/badge/maragu.dev/gomponents-htmx)](https://pkg.go.dev/maragu.dev/gomponents-htmx)
[![CI](https://github.com/maragudk/gomponents-htmx/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/gomponents-htmx/actions/workflows/ci.yml)

[HTMX](https://htmx.org) attributes and helpers for [gomponents](https://www.gomponents.com).

Made with ✨sparkles✨ by [maragu](https://www.maragu.dev/).

Does your company depend on this project? [Contact me at markus@maragu.dk](mailto:markus@maragu.dk?Subject=Supporting%20your%20project) to discuss options for a one-time or recurring invoice to ensure its continued thriving.

## Usage

```shell
go get maragu.dev/gomponents-htmx
```

```go
package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
	. "maragu.dev/gomponents/http"

	hx "maragu.dev/gomponents-htmx"
	hxhttp "maragu.dev/gomponents-htmx/http"
)

func main() {
	if err := start(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func start() error {
	now := time.Now()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		if r.Method == http.MethodPost && hxhttp.IsBoosted(r.Header) {
			now = time.Now()

			hxhttp.SetPushURL(w.Header(), "/?time="+now.Format(timeOnly))

			return partial(now), nil
		}
		return page(now), nil
	}))

	log.Println("Starting on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

const timeOnly = "15:04:05"

func page(now time.Time) Node {
	return HTML5(HTML5Props{
		Title: now.Format(timeOnly),

		Head: []Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org")),
		},

		Body: []Node{
			Div(Class("max-w-7xl mx-auto p-4 prose lg:prose-lg xl:prose-xl"),
				H1(Text(`gomponents + HTMX`)),

				P(Textf(`Time at last full page refresh was %v.`, now.Format(timeOnly))),

				partial(now),

				Form(Method("post"), Action("/"),
					hx.Boost("true"), hx.Target("#partial"), hx.Swap("outerHTML"),

					Button(Type("submit"), Text(`Update time`),
						Class("rounded-md border border-transparent bg-orange-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:ring-offset-2"),
					),
				),
			),
		},
	})
}

func partial(now time.Time) Node {
	return P(ID("partial"), Textf(`Time was last updated at %v.`, now.Format(timeOnly)))
}
```
