package gpay

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRoute(t *testing.T) {
	r := mux.NewRouter()
	RegisterRoutes(r, nil)
	server := httptest.NewServer(r)
	defer server.Close()

	t.Run("echo", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/v1/echo")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(readAll(t, resp.Body))
	})

	t.Run("association", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/v1/associateAccount", "", nil)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(readAll(t, resp.Body))
	})
}

func readAll(t *testing.T, r io.Reader) string {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	return string(raw)
}
