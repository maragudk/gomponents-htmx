// Package htmx provides HTMX attributes and helpers for gomponents.
// See https://htmx.org/
package htmx

import (
	"io"

	g "maragu.dev/gomponents"
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

// On handles any event with a script inline.
// See https://htmx.org/attributes/hx-on
func On(name string, v string) g.Node {
	return &rawAttr{name: "on:" + name, value: v}
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

// Confirm shows a confirm() dialog before issuing a request.
// See https://htmx.org/attributes/hx-confirm
func Confirm(v string) g.Node {
	return attr("confirm", v)
}

// Delete will issue a DELETE to the specified URL and swap the HTML into the DOM using a swap strategy.
// See https://htmx.org/attributes/hx-delete
func Delete(v string) g.Node {
	return attr("delete", v)
}

// Disable htmx processing for the given node and any children nodes.
// See https://htmx.org/attributes/hx-disable
func Disable(v string) g.Node {
	return attr("disable", v)
}

// Disable element until htmx request completes.
// See https://htmx.org/attributes/hx-disabled-elt/
func DisabledElt(v string) g.Node {
	return attr("disabled-elt", v)
}

// Disinherit controls and disables automatic attribute inheritance for child nodes.
// See https://htmx.org/attributes/hx-disinherit
func Disinherit(v string) g.Node {
	return attr("disinherit", v)
}

// Encoding changes the request encoding type.
// See https://htmx.org/attributes/hx-encoding
func Encoding(v string) g.Node {
	return attr("encoding", v)
}

// Ext sets extensions to use for this element.
// See https://htmx.org/attributes/hx-ext
func Ext(v string) g.Node {
	return attr("ext", v)
}

// Headers adds to the headers that will be submitted with the request.
// See https://htmx.org/attributes/hx-headers
func Headers(v string) g.Node {
	return attr("headers", v)
}

// History prevents sensitive data being saved to the history cache.
// See https://htmx.org/attributes/hx-history
func History(v string) g.Node {
	return attr("history", v)
}

// HistoryElt sets the element to snapshot and restore during history navigation.
// See https://htmx.org/attributes/hx-history-elt
func HistoryElt(v string) g.Node {
	return attr("history-elt", v)
}

// Include additional data in requests.
// See https://htmx.org/attributes/hx-include
func Include(v string) g.Node {
	return attr("include", v)
}

// Indicator sets the element to put the htmx-request class on during the request.
// See https://htmx.org/attributes/hx-indicator
func Indicator(v string) g.Node {
	return attr("indicator", v)
}

// Params filters the parameters that will be submitted with a request.
// See https://htmx.org/attributes/hx-params
func Params(v string) g.Node {
	return attr("params", v)
}

// Patch issues a PATCH to the specified URL.
// See https://htmx.org/attributes/hx-patch
func Patch(v string) g.Node {
	return attr("patch", v)
}

// Preserve specifies elements to keep unchanged between requests.
// See https://htmx.org/attributes/hx-preserve
func Preserve(v string) g.Node {
	return attr("preserve", v)
}

// Prompt shows a prompt() before submitting a request.
// See https://htmx.org/attributes/hx-prompt
func Prompt(v string) g.Node {
	return attr("prompt", v)
}

// Put issues a PUT to the specified URL.
// See https://htmx.org/attributes/hx-put
func Put(v string) g.Node {
	return attr("put", v)
}

// ReplaceURL replaces the URL in the browser location bar.
// See https://htmx.org/attributes/hx-replace-url
func ReplaceURL(v string) g.Node {
	return attr("replace-url", v)
}

// Request configures various aspects of the request.
// See https://htmx.org/attributes/hx-request
func Request(v string) g.Node {
	return attr("request", v)
}

// Sync controls how requests made by different elements are synchronized.
// See https://htmx.org/attributes/hx-sync
func Sync(v string) g.Node {
	return attr("sync", v)
}

// Validate forces elements to validate themselves before a request.
// See https://htmx.org/attributes/hx-validate
func Validate(v string) g.Node {
	return attr("validate", v)
}

func attr(name, value string) g.Node {
	return g.Attr("hx-"+name, value)
}

// rawAttr is an attribute that doesn't escape its value.
type rawAttr struct {
	name  string
	value string
}

func (r *rawAttr) Render(w io.Writer) error {
	_, err := w.Write([]byte(" hx-" + r.name + `="` + r.value + `"`))
	return err
}

func (r *rawAttr) Type() g.NodeType {
	return g.AttributeType
}
