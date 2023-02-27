// Package http provides helpers for working with HTMX HTTP requests.
// See https://htmx.org/reference/#headers
package http

import (
	"net/http"
)

// IsBoosted indicates that the request is via an element using hx-boost.
func IsBoosted(headers http.Header) bool {
	return headers.Get("HX-Boosted") != ""
}

// GetCurrentURL of the browser.
func GetCurrentURL(headers http.Header) string {
	return headers.Get("HX-Current-URL")
}

// IsHistoryRestoreRequest returns whether the request is for history restoration
// after a miss in the local history cache.
func IsHistoryRestoreRequest(headers http.Header) bool {
	return headers.Get("HX-History-Restore-Request") == "true"
}

// GetPrompt from the user response to an hx-prompt.
// See https://htmx.org/attributes/hx-prompt
func GetPrompt(headers http.Header) string {
	return headers.Get("HX-Prompt")
}

// IsRequest returns whether this is a HTMX request.
func IsRequest(headers http.Header) bool {
	return headers.Get("HX-Request") == "true"
}

// GetTarget returns the id of the target element if it exists.
func GetTarget(headers http.Header) string {
	return headers.Get("HX-Target")
}

// GetTriggerName returns the name of the triggered element if it exists.
func GetTriggerName(headers http.Header) string {
	return headers.Get("HX-Trigger-Name")
}

// GetTrigger returns the id of the triggered element if it exists.
func GetTrigger(headers http.Header) string {
	return headers.Get("HX-Trigger")
}

// SetLocation allows you to do a client-side redirect that does not do a full page reload.
// See https://htmx.org/headers/hx-location
func SetLocation(headers http.Header, v string) {
	headers.Set("HX-Location", v)
}

// SetPushURL pushes a new URL into the history stack.
// See https://htmx.org/headers/hx-push-url
func SetPushURL(headers http.Header, v string) {
	headers.Set("HX-Push-Url", v)
}

// SetRedirect can be used to do a client-side redirect to a new location.
func SetRedirect(headers http.Header, v string) {
	headers.Set("HX-Redirect", v)
}

// SetRefresh will make the client side do a a full refresh of the page.
func SetRefresh(headers http.Header) {
	headers.Set("HX-Refresh", "true")
}

// SetReplaceURL replaces the current URL in the location bar.
// See https://htmx.org/headers/hx-replace-url
func SetReplaceURL(headers http.Header, v string) {
	headers.Set("HX-Replace-Url", v)
}

// SetReswap allows you to specify how the response will be swapped.
// See https://htmx.org/attributes/hx-swap
func SetReswap(headers http.Header, v string) {
	headers.Set("HX-Reswap", v)
}

// SetRetarget sets a CSS selector that updates the target of the content update to a different element on the page.
func SetRetarget(headers http.Header, v string) {
	headers.Set("HX-Retarget", v)
}

// SetTrigger allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func SetTrigger(headers http.Header, v string) {
	headers.Set("HX-Trigger", v)
}

// SetTriggerAfterSettle allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func SetTriggerAfterSettle(headers http.Header, v string) {
	headers.Set("HX-Trigger-After-Settle", v)
}

// SetTriggerAfterSwap allows you to trigger client side events.
// See https://htmx.org/headers/hx-trigger
func SetTriggerAfterSwap(headers http.Header, v string) {
	headers.Set("HX-Trigger-After-Swap", v)
}
