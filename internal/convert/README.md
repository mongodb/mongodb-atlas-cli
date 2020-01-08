# Convert

This package aims at converting a CLI version of the automation config to the format required by the C/OM API

Some examples of the expected CLI format are as follow

## JSON

```json
{
  "name": "cluster_3",
  "version": "4.2.2",
  "feature_compatibility_version": "4.2",
  "processes": [
    {
      "hostname": "host0",
      "db_path": "/data/cluster_3/rs1",
      "log_path": "/data/cluster_3/rs1/mongodb.log",
      "priority": 1,
      "votes": 1,
      "port": 30010
    },
    {
      "hostname": "host1",
      "db_path": "/data/cluster_3/rs2",
      "log_path": "/data/cluster_3/rs2/mongodb.log",
      "priority": 1,
      "votes": 1,
      "port": 30020
    },
    {
      "hostname": "host2",
      "db_path": "/data/cluster_3/rs3",
      "log_path": "/data/cluster_3/rs3/mongodb.log",
      "priority": 1,
      "votes": 1,
      "port": 30030
    }
  ]
}
```

## YAML

```yaml
name: "cluster_2"
version: 4.2.2
feature_compatibility_version: 4.2
processes:
  - hostname: host0
    db_path: /data/cluster_2/rs1
    log_path: /data/cluster_2/rs1/mongodb.log
    priority: 1
    votes: 1
    port: 29010
  - hostname: host1
    db_path: /data/cluster_2/rs2
    log_path: /data/cluster_2/rs2/mongodb.log
    priority: 1
    votes: 1
    port: 29020
  - hostname: host2
    db_path: /data/cluster_2/rs3
    log_path: /data/cluster_2/rs3/mongodb.log
    priority: 1
    votes: 1
    port: 29030
```

**Note:** Current implementation of the mapping assumes all processes will use the same binary version and feature compatibility version
