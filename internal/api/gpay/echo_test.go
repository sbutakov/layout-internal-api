package gpay

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	server := httptest.NewServer(echoServer())
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	t.Log(string(raw))
}
