package kubernetes

import (
	"github.com/weaveworks/scope/report"
	storage_v1 "k8s.io/api/storage/v1"
)

// StorageClass represent kubernetes StorageClass interface
type StorageClass interface {
	Meta
	GetNode(probeID string) report.Node
}

// storageClass represents kubernetes storage classes
type storageClass struct {
	*storage_v1.StorageClass
	Meta
}

type person struct {
	name string
}

// NewStorageClass returns new Storage Class type
func NewStorageClass(p *storage_v1.StorageClass) StorageClass {
	return &storageClass{StorageClass: p, Meta: meta{p.ObjectMeta}}
}

// GetNode returns StorageClass as Node
func (p *storageClass) GetNode(probeID string) report.Node {
	return p.MetaNode(report.MakeStorageClassNodeID(p.UID())).WithLatests(map[string]string{
		report.ControlProbeID: probeID,
		NodeType:              "StorageClass",
		Namespace:             p.GetNamespace(),
	})
}
