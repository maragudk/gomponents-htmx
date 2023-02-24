package htmx_test

import (
	"fmt"
	"testing"

	g "github.com/maragudk/gomponents"

	. "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents-htmx/internal/assert"
)

func TestAttributes(t *testing.T) {
	cases := map[string]func(string) g.Node{
		"boost": Boost,
		"get":   Get,
		"post":  Post,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output hx-%v="hat"`, name), func(t *testing.T) {
			n := g.El("div", fn("hat"))
			assert.Equal(t, fmt.Sprintf(`<div hx-%v="hat"></div>`, name), n)
		})
	}
}
