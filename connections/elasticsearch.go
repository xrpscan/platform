package connections

import (
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/xrpscan/platform/config"
)

var EsClient *elasticsearch.Client
var once_es sync.Once

func NewEsClient() {
	once_es.Do(func() {
		config := elasticsearch.Config{
			Addresses: []string{
				config.EnvEsURL(),
			},
			Username:               config.EnvEsUsername(),
			Password:               config.EnvEsPassword(),
			CertificateFingerprint: config.EnvEsFingerprint(),
		}

		es, err := elasticsearch.NewClient(config)
		if err != nil {
			log.Fatalf("Error creating elasticsearch connection: %s\n", err)
		}

		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error getting elasticsearch info: %s\n", err)
		}
		defer res.Body.Close()
		log.Println(res)
		EsClient = es
	})
}

func GetEsClient() *elasticsearch.Client {
	NewEsClient()
	return EsClient
}
