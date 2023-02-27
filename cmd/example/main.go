package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	ghttp "github.com/maragudk/gomponents/http"

	hx "github.com/maragudk/gomponents-htmx"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
)

func main() {
	if err := start(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func start() error {
	now := time.Now()
	mux := http.NewServeMux()
	mux.HandleFunc("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		if r.Method == http.MethodPost && hxhttp.IsBoosted(r.Header) {
			now = time.Now()

			hxhttp.SetPushURL(w.Header(), "/?time="+now.Format(timeFormat))

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

const timeFormat = "15:04:05"

func page(now time.Time) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: now.Format(timeFormat),
		Head: []g.Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org")),
		},
		Body: []g.Node{
			Div(Class("max-w-7xl mx-auto p-4 prose lg:prose-lg xl:prose-xl"),
				H1(g.Text(`gomponents + HTMX`)),
				P(g.Textf(`Time at last full page refresh was %v.`, now.Format(timeFormat))),
				partial(now),
				FormEl(Method("post"), Action("/"), hx.Boost("true"), hx.Target("#partial"), hx.Swap("outerHTML"),
					Button(Type("submit"), g.Text(`Update time`),
						Class("rounded-md border border-transparent bg-orange-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:ring-offset-2"),
					),
				),
			),
		},
	})
}

func partial(now time.Time) g.Node {
	return P(ID("partial"), g.Textf(`Time was last updated at %v.`, now.Format(timeFormat)))
}
