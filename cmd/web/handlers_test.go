package main

import (
	"net/http"
	"testing"

	"snippetbox.pp.com/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	cody, _, body := ts.get(t, "/ping")

	assert.Equal(t, cody, http.StatusOK)

	assert.Equal(t, body, "OK")
}
