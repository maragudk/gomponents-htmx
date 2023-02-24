// Package htmx provides HTMX attributes and helpers for gomponents.
// See https://htmx.org/
package htmx

import (
	g "github.com/maragudk/gomponents"
)

// Boost for links and forms.
// See https://htmx.org/attributes/hx-boost
func Boost(v string) g.Node {
	return attr("boost", v)
}

// Get the url.
// See https://htmx.org/attributes/hx-get
func Get(url string) g.Node {
	return attr("get", url)
}

// Post to the url.
// See https://htmx.org/attributes/hx-post
func Post(url string) g.Node {
	return attr("post", url)
}

func attr(name, value string) g.Node {
	return g.Attr("hx-"+name, value)
}
