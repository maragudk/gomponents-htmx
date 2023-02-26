package htmx_test

import (
	"fmt"
	"os"
	"testing"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents-htmx/internal/assert"
)

func TestAttributes(t *testing.T) {
	cases := map[string]func(string) g.Node{
		"boost":      hx.Boost,
		"get":        hx.Get,
		"post":       hx.Post,
		"push-url":   hx.PushURL,
		"select":     hx.Select,
		"select-oob": hx.SelectOOB,
		"swap":       hx.Swap,
		"swap-oob":   hx.SwapOOB,
		"target":     hx.Target,
		"trigger":    hx.Trigger,
		"vals":       hx.Vals,
		"delete":     hx.Delete,
		"confirm":    hx.Confirm,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output hx-%v="hat"`, name), func(t *testing.T) {
			n := g.El("div", fn("hat"))
			assert.Equal(t, fmt.Sprintf(`<div hx-%v="hat"></div>`, name), n)
		})
	}
}

func ExampleGet() {
	n := Button(hx.Post("/clicked"), hx.Swap("outerHTML"))
	_ = n.Render(os.Stdout)
	// Output: <button hx-post="/clicked" hx-swap="outerHTML"></button>
}
