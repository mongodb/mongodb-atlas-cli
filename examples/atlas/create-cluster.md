# Create a Cluster

## Three member replicaset

To create and get the connection string for a cluster you can do:

```bash
atlas clusters create MyCluster \
  --region EASTERN_US \
  --members 3 \
  --tier M30 \
  --provider GCP \
  --mdbVersion 6.0 \
  --diskSizeGB 30 && \
atlas clusters watch MyCluster && \
atlas clusters describe MyCluster -o go-template="{{.SrvAddress}}"
```

## Multi-Cloud replicaset

Using the following json as an input file
```json
{
    "clusterType": "REPLICASET",
    "links": [],
    "name": "multiCloud",
    "replicationSpecs": [
      {
        "numShards": 1,
        "regionConfigs": [
          {
            "analyticsAutoScaling": {
              "autoIndexing": {
                "enabled": false
              },
              "compute": {
                "enabled": true,
                "maxInstanceSize": "M40",
                "minInstanceSize": "M30",
                "scaleDownEnabled": true
              },
              "diskGB": {
                "enabled": true
              }
            },
            "analyticsSpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "autoScaling": {
              "autoIndexing": {
                "enabled": false
              },
              "compute": {
                "enabled": true,
                "maxInstanceSize": "M40",
                "minInstanceSize": "M30",
                "scaleDownEnabled": true
              },
              "diskGB": {
                "enabled": true
              }
            },
            "electableSpecs": {
              "instanceSize": "M30",
              "nodeCount": 3
            },
            "hiddenSecondarySpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "priority": 7,
            "providerName": "AWS",
            "readOnlySpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "regionName": "US_EAST_1"
          },
          {
            "analyticsAutoScaling": {
              "autoIndexing": {
                "enabled": false
              },
              "compute": {
                "enabled": true,
                "maxInstanceSize": "M40",
                "minInstanceSize": "M30",
                "scaleDownEnabled": true
              },
              "diskGB": {
                "enabled": true
              }
            },
            "analyticsSpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "autoScaling": {
              "autoIndexing": {
                "enabled": false
              },
              "compute": {
                "enabled": true,
                "maxInstanceSize": "M40",
                "minInstanceSize": "M30",
                "scaleDownEnabled": true
              },
              "diskGB": {
                "enabled": true
              }
            },
            "electableSpecs": {
              "instanceSize": "M30",
              "nodeCount": 2
            },
            "hiddenSecondarySpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "priority": 6,
            "providerName": "GCP",
            "readOnlySpecs": {
              "instanceSize": "M30",
              "nodeCount": 0
            },
            "regionName": "EASTERN_US"
          }
        ],
        "zoneName": "Zone 1"
      }
    ]
}
```

```bash
atlas clusters create \
  --file cluster.json
```
