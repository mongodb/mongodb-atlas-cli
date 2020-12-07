package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestNewReplicaSetProcessConfig(t *testing.T) {
	trueValue := true
	omp := &opsmngr.Process{
		Args26: opsmngr.Args26{
			AuditLog: &opsmngr.AuditLog{
				Destination: "file",
				Path:        "/data/audit.log",
				Format:      "JSON",
				Filter:      "{ atype: { $in: [ \"createCollection\", \"dropCollection\" ] } }",
			},
			NET: opsmngr.Net{
				Port: 27017,
				TLS:  &opsmngr.TLS{Mode: "disabled"},
			},
			Replication: &opsmngr.Replication{
				ReplSetName: "myReplicaSet",
			},
			Storage: &opsmngr.Storage{
				DBPath:         "/data/db",
				DirectoryPerDB: &trueValue,
				WiredTiger: &map[string]interface{}{
					"collectionConfig": map[string]interface{}{},
					"engineConfig": map[string]interface{}{
						"cacheSizeGB": 1,
					},
					"indexConfig": map[string]interface{}{},
				},
			},
			SystemLog: opsmngr.SystemLog{
				Destination: "file",
				Path:        "/data/log/mongodb.log",
			},
		},
		AuthSchemaVersion:           5,
		Disabled:                    false,
		FeatureCompatibilityVersion: "4.4",
		Hostname:                    "n1.omansible.int",
		LogRotate: &opsmngr.LogRotate{
			SizeThresholdMB:  1000,
			TimeThresholdHrs: 24,
		},
		ManualMode:  false,
		Name:        "myReplicaSet_1",
		ProcessType: "mongod",
		Version:     "4.4.1-ent",
	}
	omm := opsmngr.Member{
		ID:           0,
		ArbiterOnly:  false,
		BuildIndexes: true,
		Hidden:       false,
		Host:         "myReplicaSet_1",
		Priority:     1,
		SlaveDelay:   0,
		Votes:        1,
	}

	one := 1.0
	zero := 0.0
	falseValue := false
	expected := &ProcessConfig{
		AuditLogPath:        "/data/audit.log",
		AuditLogDestination: "file",
		AuditLogFormat:      "JSON",
		AuditLogFilter:      "{ atype: { $in: [ \"createCollection\", \"dropCollection\" ] } }",
		BuildIndexes:        &trueValue,
		DBPath:              "/data/db",
		DirectoryPerDB:      &trueValue,
		FCVersion:           "4.4",
		Hostname:            "n1.omansible.int",
		LogDestination:      "file",
		LogPath:             "/data/log/mongodb.log",
		Name:                "myReplicaSet_1",
		Port:                27017,
		Priority:            &one,
		ProcessType:         "mongod",
		SlaveDelay:          &zero,
		Version:             "4.4.1-ent",
		Votes:               &one,
		ArbiterOnly:         &falseValue,
		Disabled:            false,
		Hidden:              &falseValue,
		TLS:                 &TLS{Mode: "disabled"},
		WiredTiger: &map[string]interface{}{
			"collectionConfig": map[string]interface{}{},
			"engineConfig": map[string]interface{}{
				"cacheSizeGB": 1,
			},
			"indexConfig": map[string]interface{}{},
		},
	}
	result := newReplicaSetProcessConfig(omm, omp)

	assert.Equal(t, expected, result)
}
