package routing_test

import (
	"testing"
	"github.com/F-Diniz/gophercon/pkg/routing"
	"net/http/httptest"
	"net/http"
)

func TestBaseRouter(t *testing.T) {
	handler := routing.BaseRouter()

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code %d (%d expected)", res.StatusCode, http.StatusOK)
	}
}