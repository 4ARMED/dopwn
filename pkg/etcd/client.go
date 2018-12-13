// Lifted largely from https://github.com/kelseyhightower/confd/blob/master/backends/etcdv3/client.go

package etcd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// Client is a wrapper around the etcd client
type Client struct {
	client *clientv3.Client
}

// NewEtcdClient returns an *etcdv3.Client with connection to DO etcd server
func NewEtcdClient(endpoint, cert, key, caCert string) (*Client, error) {
	var cli *clientv3.Client
	cfg := clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	if caCert != "" {
		certBytes, err := ioutil.ReadFile(caCert)
		if err != nil {
			return &Client{cli}, err
		}

		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM(certBytes)

		if ok {
			tlsConfig.RootCAs = caCertPool
		}
	}

	if cert != "" && key != "" {
		tlsCert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return &Client{cli}, err
		}
		tlsConfig.Certificates = []tls.Certificate{tlsCert}
	}

	cfg.TLS = tlsConfig

	cli, err := clientv3.New(cfg)
	if err != nil {
		return &Client{cli}, err
	}
	return &Client{cli}, nil

}

// GetValue returns specified keys from etcd
func (c *Client) GetValue(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	kv := clientv3.NewKV(c.client)
	gr, err := kv.Get(ctx, key)
	cancel()
	if err != nil {
		return []byte(gr.Kvs[0].Value), err
	}

	return []byte(gr.Kvs[0].Value), err
}
