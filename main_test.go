package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	// テスト用サーバーの立ち上げ
	ts := httptest.NewServer(Route())
	defer ts.Close()

	// テスト用サーバーのURLは、ts.URL で取得できる
	res, err := http.Get(ts.URL + "/contact?name=gopher")
	if err != nil {
		t.Fatalf("http.Get faild: %s", err)
	}
	contact, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("read from HTTP Response Body failed: %s", err)
	}
	expected := "Hello, gopher"
	if string(contact) != expected {
		t.Fatalf("response of /contact?name=gopher returns %s, want %s", string(contact), expected)
	}
}

// BenchmarkCalc 足し算する
func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calc(b.N, b.N)
	}
}
