package justwatch_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchNew(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func (response http.ResponseWriter, req *http.Request){
		
	}))
	defer ts.Cl
}
