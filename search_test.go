package justwatch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchNew(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintln(rw, "{\"skip_count\":0,\"days\":[{\"date\":\"2020-06-25\",\"providers\":[{\"total\":1,\"provider_id\":149,\"items\":[{\"jw_entity_id\":\"tm45815\",\"id\":45815,\"title\":\"El bibliotecario: La maldición del cáliz de Judas\",\"full_path\":\"/es/pelicula/el-bibliotecario-la-maldicion-del-caliz-de-judas\",\"poster\":\"/poster/81171186/{profile}\",\"object_type\":\"movie\",\"offers\":[{\"monetization_type\":\"flatrate\",\"provider_id\":149,\"currency\":\"EUR\",\"urls\":{\"standard_web\":\"http://ver.movistarplus.es/ficha/la-maldicion-del-caliz-de-judas?id=863961\"},\"presentation_type\":\"sd\"}]}]}]}],\"page\":0,\"page_size\":0}")
		if err != nil {
			t.Error("Error when creating the mock http server")
		}
	}))
	defer ts.Close()

	client, err := NewClient()
	if err != nil {
		t.Error("Error when initializing client")
	}
}
