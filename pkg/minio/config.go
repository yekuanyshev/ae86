package minio

import "net"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	UseSSL   bool
	Bucket   string
}

func (c Config) Endpoint() string {
	return net.JoinHostPort(c.Host, c.Port)
}
