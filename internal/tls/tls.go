package tls

import (
	"crypto/tls"
	"github.com/jhue/misgo/internal/conf"
)

func GetTLSConfig() (*tls.Config, error) {
	config := conf.GetConfig().TLSConfig
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	cert, err := tls.LoadX509KeyPair(config.CrtPath, config.KeyPath)
	if err != nil {
		return nil, err
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	return cfg, nil
}
