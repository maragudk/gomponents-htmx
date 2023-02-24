// Package htmx provides HTMX attributes and helpers for gomponents.
// See https://htmx.org/
package htmx

import (
	g "github.com/maragudk/gomponents"
)

// Boost to add or remove progressive enhancement for links and forms.
// See https://htmx.org/attributes/hx-boost
func Boost(v string) g.Node {
	return attr("boost", v)
}

// Get from the specified URL.
// See https://htmx.org/attributes/hx-get
func Get(url string) g.Node {
	return attr("get", url)
}

// Post to the specified URL.
// See https://htmx.org/attributes/hx-post
func Post(url string) g.Node {
	return attr("post", url)
}

// PushURL into the browser location bar, creating a new history entry.
// See https://htmx.org/attributes/hx-push-url
func PushURL(v string) g.Node {
	return attr("push-url", v)
}

// Select content to swap in from a response.
// See https://htmx.org/attributes/hx-select
func Select(v string) g.Node {
	return attr("select", v)
}

// SelectOOB content to swap in from a response, out of band (somewhere other than the target).
// See https://htmx.org/attributes/hx-select-oob
func SelectOOB(v string) g.Node {
	return attr("select-oob", v)
}

// Swap controls how content is swapped in.
// See https://htmx.org/attributes/hx-swap
func Swap(v string) g.Node {
	return attr("swap", v)
}

// SwapOOB marks content in a response to be out of band (should swap in somewhere other than the target).
// See https://htmx.org/attributes/hx-swap-oob
func SwapOOB(v string) g.Node {
	return attr("swap-oob", v)
}

// Target specifies the target element to be swapped.
// See https://htmx.org/attributes/hx-target
func Target(v string) g.Node {
	return attr("target", v)
}

// Trigger specifies the event that triggers the request.
// See https://htmx.org/attributes/hx-trigger
func Trigger(v string) g.Node {
	return attr("trigger", v)
}

// Vals adds values to the parameters to submit with the request (JSON-formatted).
// See https://htmx.org/attributes/hx-vals
func Vals(v string) g.Node {
	return attr("vals", v)
}

func attr(name, value string) g.Node {
	return g.Attr("hx-"+name, value)
}
