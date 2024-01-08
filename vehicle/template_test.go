package vehicle

import (
	"context"
	"testing"

	"github.com/evcc-io/evcc/util/templates"
	"github.com/evcc-io/evcc/util/test"
)

var acceptable = []string{
	"invalid plugin source: ...",
	"missing mqtt broker configuration",
	"received status code 404 (INVALID PARAMS)", // Nissan
	"missing personID",
	"401 Unauthorized",
	"unexpected length",
	"i/o timeout",
	"no such host",
	"network is unreachable",
	"error connecting: Network Error",
	"unexpected status: 401",
	"missing credentials",    // Tesla
	"missing credentials id", // Tronity
	"missing access and/or refresh token, use `evcc token` to create", // Tesla
}

func TestTemplates(t *testing.T) {
	templates.TestClass(t, templates.Vehicle, func(t *testing.T, values map[string]any) {
		t.Helper()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if _, err := NewFromConfig(ctx, "template", values); err != nil && !test.Acceptable(err, acceptable) {
			t.Log(values)
			t.Error(err)
		}
	})
}
