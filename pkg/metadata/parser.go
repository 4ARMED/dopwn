package metadata

import (
	"fmt"

	"github.com/kubicorn/kubicorn/pkg/logger"
	yaml "gopkg.in/yaml.v2"
)

// Metadata stores the Kubernetes-related YAML
type Metadata struct {
	EtcdCa             string `yaml:"k8saas_etcd_ca"`
	EtcdKey            string `yaml:"k8saas_etcd_key"`
	EtcdCert           string `yaml:"k8saas_etcd_cert"`
	KubeMasterHostname string `yaml:"k8saas_master_domain_name"`
}

// ParseMetadata reads in a byte array of metadata from Digital Ocean's metadata service
// and returns a struct containing the etcd credentials
func ParseMetadata(in []byte, out *Metadata) (err error) {
	logger.Debug("ParseMetadata called")
	err = yaml.Unmarshal([]byte(in), &out)
	if err != nil {
		return fmt.Errorf("unable to parse metadata: %v", err)
	}

	return nil
}
