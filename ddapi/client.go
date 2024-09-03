package ddapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	cache "dev.azure.com/bbts-lab/DevSecOps/_git/ddconnector/config"
)

type Client struct {
	cache               cache.Cache
	token               string
	httpClient          http.Client
	skipSSLVerification bool
}

// NewClient cria uma nova instância do cliente.
// A opção skipSSLVerification deve ser usada com cautela. Em ambientes de produção, é importante garantir
// que a comunicação seja segura e que os certificados SSL/TLS sejam verificados para evitar ataques de Man-in-the-Middle (MitM).
// Use skipSSLVerification apenas em ambientes de teste ou desenvolvimento onde a comunicação segura não é crítica.
func NewClient(timeout time.Duration, cacheInterval time.Duration, skipSSLVerification bool) Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipSSLVerification, // Atenção: Definir como true pode ser inseguro em produção.
			MinVersion:         tls.VersionTLS13,
		},
	}
	return Client{
		cache: cache.NewCache(cacheInterval),
		token: os.Getenv("DD_API_KEY"),
		httpClient: http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		skipSSLVerification: skipSSLVerification,
	}
}
