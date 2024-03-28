//go:build !appengine
// +build !appengine

package tokbox

import (
	"net/http"

	"golang.org/x/net/context"
)

func client(_ context.Context) *http.Client {
	return &http.Client{}
}
