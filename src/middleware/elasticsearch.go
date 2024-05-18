package middleware

import (
	"crypto/tls"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net"
	"net/http"
	"time"
)

var esClient *elasticsearch.Client

func init() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "foo",
		Password:  "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12}}}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Print(es.Transport.(*elastictransport.Client).URLs())
}

func GetElasticsearchClient() *elasticsearch.Client {
	return esClient
}
