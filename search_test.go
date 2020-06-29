package justwatch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchNew(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			fmt.Fprintln(rw, "[{\"exposed_url_part\":\"es\",\"is_standard_locale\":true,\"full_locale\":\"es_ES\",\"i18n_state\":\"active\",\"iso_3166_2\":\"ES\",\"country\":\"Spain\",\"currency\":\"EUR\",\"currency_name\":\"Euro\",\"country_names\":{\"de\":\"Spanien\",\"es\":\"España\",\"fr\":\"Espagne\",\"hr\":\"Španjolska\",\"it\":\"Spagna\",\"ja\":\"スペイン\",\"nl\":\"Spanje\",\"ru\":\"Испания\"}}]")
		} else {
			fmt.Fprintln(rw, "{\"skip_count\":0,\"days\":[{\"date\":\"2020-06-25\",\"providers\":[{\"total\":1,\"provider_id\":149,\"items\":[{\"jw_entity_id\":\"tm45815\",\"id\":45815,\"title\":\"El bibliotecario: La maldición del cáliz de Judas\",\"full_path\":\"/es/pelicula/el-bibliotecario-la-maldicion-del-caliz-de-judas\",\"poster\":\"/poster/81171186/{profile}\",\"object_type\":\"movie\",\"offers\":[{\"monetization_type\":\"flatrate\",\"provider_id\":149,\"currency\":\"EUR\",\"urls\":{\"standard_web\":\"http://ver.movistarplus.es/ficha/la-maldicion-del-caliz-de-judas?id=863961\"},\"presentation_type\":\"sd\"}]}]}]}],\"page\":0,\"page_size\":0}")
		}
	}))
	defer ts.Close()

	client, err := NewClient(SetURL(ts.URL))
	if err != nil {
		t.Error("Error when initializing client")
	}
	response, err := client.SearchNew(nil)
	if err != nil {
		t.Error("Error when receiving response from server")
	}
	if len(response.Days) != 1 {
		t.Errorf("Incorrect number of Days elements, got: %d, want: %d", len(response.Days), 1)
	}
	if len(response.Days[0].Providers) != 1 {
		t.Errorf("Incorrect number of Providers elements, got: %d, want: %d", len(response.Days[0].Providers), 1)
	}
	if response.Days[0].Providers[0].ProviderID != 149 {
		t.Errorf("Incorrect provider ID, got: %d, want: %d", response.Days[0].Providers[0].ProviderID, 1)
	}
}
