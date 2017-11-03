package gorreios

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/encoding/charmap"
)

const (
	host = "http://www2.correios.com.br/sistemas/rastreamento/resultado_semcontent.cfm"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

func GetTrackingInfo(code string) (tracking Tracking, err error) {
	data := url.Values{}
	data.Set("Objetos", code)
	resp, err := httpClient.PostForm(host, data)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	utf8Body := charmap.ISO8859_1.NewDecoder().Reader(resp.Body)
	tracking, err = parseTracking(utf8Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}
