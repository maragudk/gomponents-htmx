package http_test

import (
	"net/http"
	"testing"

	"github.com/maragudk/is"

	hxhttp "github.com/maragudk/gomponents-htmx/http"
)

func TestBoolGetters(t *testing.T) {
	tests := map[string]func(http.Header) bool{
		"Boosted":                 hxhttp.IsBoosted,
		"History-Restore-Request": hxhttp.IsHistoryRestoreRequest,
		"Request":                 hxhttp.IsRequest,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}

			v := fn(headers)
			is.Equal(t, false, v)

			headers.Set("HX-"+name, "true")
			v = fn(headers)
			is.Equal(t, true, v)
		})
	}
}

func TestValueGetters(t *testing.T) {
	tests := map[string]func(http.Header) string{
		"Current-URL":  hxhttp.GetCurrentURL,
		"Prompt":       hxhttp.GetPrompt,
		"Target":       hxhttp.GetTarget,
		"Trigger-Name": hxhttp.GetTriggerName,
		"Trigger":      hxhttp.GetTrigger,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}

			v := fn(headers)
			is.Equal(t, "", v)

			headers.Set("HX-"+name, "foo")
			v = fn(headers)
			is.Equal(t, "foo", v)
		})
	}
}

func TestSetRefresh(t *testing.T) {
	headers := http.Header{}
	hxhttp.SetRefresh(headers)
	is.Equal(t, "true", headers.Get("HX-Refresh"))
}

func TestSetters(t *testing.T) {
	tests := map[string]func(http.Header, string){
		"Location":             hxhttp.SetLocation,
		"Push-Url":             hxhttp.SetPushURL,
		"Redirect":             hxhttp.SetRedirect,
		"Replace-Url":          hxhttp.SetReplaceURL,
		"Reswap":               hxhttp.SetReswap,
		"Retarget":             hxhttp.SetRetarget,
		"Trigger":              hxhttp.SetTrigger,
		"Trigger-After-Settle": hxhttp.SetTriggerAfterSettle,
		"Trigger-After-Swap":   hxhttp.SetTriggerAfterSwap,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			fn(headers, "foo")
			is.Equal(t, "foo", headers.Get("HX-"+name))
		})
	}
}

func ExampleIsBoosted() {
	_ = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hxhttp.IsBoosted(r.Header) {
			// Boosted!
		}
	})
}
