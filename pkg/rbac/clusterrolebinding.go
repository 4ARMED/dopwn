package rbac

import (
	"time"

	"github.com/pborman/uuid"
)

// Metadata redefines a bunch of things from k8s
type Metadata struct {
	CreationTimeStamp string    `json:"creationTimestamp"`
	Name              string    `json:"name"`
	UID               uuid.UUID `json:"uid"`
}

// RoleRef redefines a bunch of things from k8s
type RoleRef struct {
	APIGroup string `json:"apiGroup"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

// Subject redefines a bunch of things from k8s
type Subject struct {
	APIGroup  string `json:"apiGroup"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// ForearmedClusterRoleBinding redefines a bunch of things from k8s
type ForearmedClusterRoleBinding struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   Metadata  `json:"metadata"`
	RoleRef    RoleRef   `json:"roleRef"`
	Subjects   []Subject `json:"subjects"`
}

// GenerateForearmedClusterRoleBinding creates a cluster role binding in JSON format suitable
// for encoding with auger and pushing into etcd
func GenerateForearmedClusterRoleBinding(name string) ForearmedClusterRoleBinding {
	uid := uuid.NewUUID()
	timestamp := time.Now().Format(time.RFC3339)

	crb := ForearmedClusterRoleBinding{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
		Metadata: Metadata{
			CreationTimeStamp: timestamp,
			Name:              name,
			UID:               uid,
		},
		RoleRef: RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "cluster-admin",
		},
		Subjects: []Subject{
			Subject{
				APIGroup:  "rbac.authorization.k8s.io",
				Kind:      "ServiceAccount",
				Name:      "default",
				Namespace: "kube-system",
			},
		},
	}

	return crb
}
