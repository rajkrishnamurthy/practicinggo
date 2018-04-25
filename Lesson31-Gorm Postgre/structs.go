package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type WorkerNode struct {
	gorm.Model          `json:"-"`
	CNAPIVersion        string `json:"cnAPIVersion,omitempty" gorm:""`
	CNPath              string `json:"cnPath,omitempty"`
	CNExternalURL       string `json:"cnExternalURL"`
	CNInternalURL       string `json:"cnInternalURL,omitempty"` // tcp://url:managementport
	CNManagementPort    string `json:"cnManagementPort,omitempty"`
	CNManagementService string `json:"cnManagementService,omitempty"`
	CNTaskDirectory     string `json:"cnTaskDirectory,omitempty"`
	CNCommandShell      string `json:"cnCommandShell,omitempty"`
}

// CNTaskPortRange     CNTaskPortRange   `json:"cnTaskPortRange,omitempty"`
// CNTaskPolicy        CNTaskPolicy      `json:"cnTaskPolicy,omitempty"`
// CNNodeTagsArray     []CNTagKV         `json:"cnNodeTagsArray,omitempty"`
// CNNodeTagsMap       map[string]string `json:"-"` // This map representation is derived from CNNodeTagsArray at Runtime based on JSON inputs, if any

type CNTagKV struct {
	gorm.Model
	Key   string `json:"key"`
	Value string `json:"value"`
}
type CNTaskPortRange struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type CNTaskPolicy struct {
	Expiration string `json:"expiration"`
}
