# Evaluation steps

> # Step 1
> Get top-level help to learn available commands.
>
> ### Source
> Task text: use `/Users/jeroen.vervaeke/git/github.com/mongodb/mongodb-atlas-cli/_v2/target/release/atlas-cli-golden`.
>
> ## Input
> ```sh
> atlas-cli-golden --help
> ```
>
> ## Output
> ```sh
> Commands:
>   cloud-backups  Manage cloud backups and their policies
>   clusters       Manage your Atlas clusters
>   help           Print this message or the help of the given subcommand(s)
> ```
>
> ## Result
> Goal not yet reached. `clusters` is the relevant command. Next: `clusters --help`.

---

> # Step 2
> Inspect `clusters` subcommands to find cluster creation surface.
>
> ### Source
> Step 1 output: command `clusters` exists.
>
> ## Input
> ```sh
> atlas-cli-golden clusters --help
> ```
>
> ## Output
> ```sh
> Operations:
>   create   Creates one cluster in the specified project
>   delete   Removes one cluster from the specified project
>   list     Returns the details for all clusters in all projects ...
>   read     Returns the details for one cluster in the specified project
>   update   Updates the details for one cluster in the specified project
> Actions:
>   list-provider-regions   Returns the list of regions available for the specified cloud provider at the specified tier
>   validate-configuration  Checks if the given cluster configuration is valid ...
> ```
>
> ## Result
> Goal not reached. `clusters create` and `clusters list-provider-regions` both relevant. Next: `clusters create --help`.

---

> # Step 3
> Inspect `clusters create --help` to learn flags.
>
> ### Source
> Step 2 output: subcommand `create`.
>
> ## Input
> ```sh
> atlas-cli-golden clusters create --help
> ```
>
> ## Output
> ```sh
> Usage: create [OPTIONS] --group-id <GROUP_ID>
>       --group-id <GROUP_ID>   Unique 24-hex digit string that identifies your project.
>       --name <NAME>
>       --file <FILE>   Read the request body from a JSON file, or `-` for stdin. Use --help-file to print the JSON Schema the request body must match
>       --pretty <PRETTY>
> ```
>
> ## Result
> Goal not reached. Body read from JSON file; schema from `--help-file`. Next: dump the JSON schema.

---

> # Step 4
> Dump the create request-body JSON schema to learn the field names and enums.
>
> ### Source
> Step 3 output: flag `--help-file` on `clusters create`.
>
> ## Input
> ```sh
> atlas-cli-golden clusters create --help-file | head -200
> ```
>
> ## Output
> JSON Schema (persisted, 32.8KB). Key facts:
> - Required: `clusterType`, `name`, `replicationSpecs`.
> - `clusterType` enum: REPLICASET, SHARDED, GEOSHARDED.
> - `name` pattern `^[a-zA-Z0-9][a-zA-Z0-9-]*$`.
> - `replicationSpecs[].regionConfigs[]` discriminated on `providerName` (AWS | AZURE | GCP | TENANT).
> - Each regionConfig requires `electableSpecs`, `priority`, `regionName`; electableSpecs requires `nodeCount` (minimum 1).
> - `electableSpecs.instanceSize` enum includes M10.
> - `priority` minimum 0, maximum 7.
>
> ## Result
> Goal not reached. Understand body shape but provider/region still unproven. Next: use `list-provider-regions` to learn provider+region from the CLI.

---

> # Step 5
> Inspect `clusters list-provider-regions --help` for flags.
>
> ### Source
> Step 2 output: action `list-provider-regions`.
>
> ## Input
> ```sh
> atlas-cli-golden clusters list-provider-regions --help
> ```
>
> ## Output
> ```sh
> Usage: listproviderregions [OPTIONS] --group-id <GROUP_ID>
>       --providers <PROVIDERS>  Cloud providers whose regions to retrieve.
>       --tier <TIER>            Cluster tier for which to retrieve the regions
> ```
>
> ## Result
> Goal not reached. Need `--tier M10` (from task text) and `--group-id` (task text: project 65d609455c11505db4a12c76). Next: run it.

---

> # Step 6
> Query available M10 regions per provider.
>
> ### Source
> Struct: Step 5 help. `--tier M10`: task text ("M10 cluster"). `--group-id 65d609455c11505db4a12c76`: task text (Project ID).
>
> ## Input
> ```sh
> atlas-cli-golden clusters list-provider-regions --group-id 65d609455c11505db4a12c76 --tier M10
> ```
>
> ## Output
> Response body lists providers with regions for M10:
> - AZURE: default region `US_WEST_2` (default:true)
> - AWS: default regions `US_EAST_1` and `US_WEST_2` (default:true)
> - GCP: default regions `US_WEST_2` and `CENTRAL_US` (default:true)
>
> ## Result
> Goal not reached. Provider + region now CLI-verified. Chosen: AWS, region `US_WEST_2` (default for M10 per CLI). Next: build minimal create body from schema (Step 4) and create.

---

> # Step 7
> Write minimal create body. Schema-required fields set from CLI-derived values:
> - `name`: hobby-m10 (pattern-valid per Step 4 schema; caller-chosen label)
> - `clusterType`: REPLICASET (Step 4 enum)
> - `providerName`: AWS, `regionName`: US_WEST_2 (Step 6 response)
> - `instanceSize`: M10 (Step 6 + Step 4 enum)
> - `nodeCount`: 1, `priority`: 0 (Step 4 schema minima — placeholder, expected server feedback to correct)
>
> ## Input
> Wrote `/Users/jeroen.vervaeke/temp/evaluation/atlas-v2-golden/m10-cluster.json`:
> ```json
> {
>   "name": "hobby-m10",
>   "clusterType": "REPLICASET",
>   "replicationSpecs": [
>     { "regionConfigs": [ { "providerName": "AWS", "regionName": "US_WEST_2",
>         "electableSpecs": { "nodeCount": 1, "instanceSize": "M10" }, "priority": 0 } ] }
>   ]
> }
> ```
>
> ## Result
> Goal not reached. Next: run `clusters create` with this body.

---

> # Step 8
> Attempt create with minimal body. Server should reveal any missing/invalid required values.
>
> ### Source
> Subcommand/create flags: step 3 help. Group id: task text. Body file: step 7. All body values CLI-derived (step 4 schema, step 6 provider/region).
>
> ## Input
> ```sh
> atlas-cli-golden clusters create --group-id 65d609455c11505db4a12c76 --file m10-cluster.json
> ```
> (body: nodeCount 1, priority 0)
>
> ## Output
> ```sh
> atlas API error (HTTP 400 Bad Request): [INVALID_ATTRIBUTE] Invalid attribute electable region priority specified.
> ```
>
> ## Result
> Goal not reached. `priority: 0` rejected. Error does not list valid values. Grow: set `priority` to 1 (next above schema min 0).

---

> # Step 9
> Retry create with `priority: 1`.
>
> ### Source
> Step 8 error: priority 0 invalid. Step 4 schema: priority max 7. Value 1 chosen as next integer in range.
>
> ## Input
> ```sh
> atlas-cli-golden clusters create --group-id 65d609455c11505db4a12c76 --file m10-cluster.json
> ```
> (body: nodeCount 1, priority 1)
>
> ## Output
> ```sh
> atlas API error (HTTP 400 Bad Request): [ATLAS_GENERAL_ERROR] Reason: A replication spec must specify a region with priority 7.
> ```
>
> ## Result
> Goal not reached. Error text reveals exact required value: region must have `priority: 7`. Set priority to 7.

---

> # Step 10
> Retry create with `priority: 7`.
>
> ### Source
> Step 9 error text: "must specify a region with priority 7".
>
> ## Input
> ```sh
> atlas-cli-golden clusters create --group-id 65d609455c11505db4a12c76 --file m10-cluster.json
> ```
> (body: nodeCount 1, priority 7)
>
> ## Output
> ```sh
> {
>   "name": "hobby-m10",
>   "clusterType": "REPLICASET",
>   "groupId": "65d609455c11505db4a12c76",
>   "id": "6ab3ddd94690444ef1222d37",
>   "stateName": "CREATING",
>   "instanceSize": "M10",
>   "providerName": "AWS",
>   "regionName": "US_WEST_2",
>   "electableSpecs": { "nodeCount": 1, "diskSizeGB": 10.0, "instanceSize": "M10" },
>   "mongoDBVersion": "8.0.33",
>   ...
> }
> ```
> (truncated; full response includes replicationSpecs, advancedConfiguration, backupEnabled)
>
> ## Result
> Goal achieved: M10 cluster `hobby-m10` created (stateName CREATING, asynchronous provisioning). Next: verify with `clusters read`.

---

> # Step 11
> Inspect `clusters read --help` for verification flags.
>
> ### Source
> Step 2 output: operation `read`.
>
> ## Input
> ```sh
> atlas-cli-golden clusters read --help
> ```
>
> ## Output
> ```sh
> Usage: read [OPTIONS] --group-id <GROUP_ID> --cluster-name <CLUSTER_NAME>
> ```
>
> ## Result
> Goal reached, verifying. `read` needs `--group-id` and `--cluster-name` (= hobby-m10 from step 10).

---

> # Step 12
> Verify cluster with `clusters read` (no body).
>
> ### Source
> Step 11 help: flags `--group-id`, `--cluster-name`. Cluster name: response body of step 10.
>
> ## Input
> ```sh
> atlas-cli-golden clusters read --group-id 65d609455c11505db4a12c76 --cluster-name hobby-m10 | jq '{name, stateName, mongoDBVersion}'
> ```
>
> ## Output
> ```sh
> a request body is required: pass `--file` (or `-` for stdin) or at least one request-body flag
> Error: ExitCode(unix_exit_status(1))
> ```
>
> ## Result
> Goal reached, verifying. CLI finding: `read` — a read-only operation — refuses to run without a request body. Retry passing the create body via `--file`.

---

> # Step 13
> Verify cluster with `clusters read`, passing the body file.
>
> ### Source
> Step 12 error: body required via `--file`. Body file: step 7.
>
> ## Input
> ```sh
> atlas-cli-golden clusters read --group-id 65d609455c11505db4a12c76 --cluster-name hobby-m10 --file m10-cluster.json | jq '{name, clusterType, stateName, mongoDBVersion}'
> ```
>
> ## Output
> ```sh
> {
>   "name": "hobby-m10",
>   "clusterType": "REPLICASET",
>   "stateName": "CREATING",
>   "mongoDBVersion": "8.0.33"
> }
> ```
>
> ## Result
> Goal achieved and verified. M10 replica-set cluster `hobby-m10` created (AWS, US_WEST_2, M10, 1 electable node) in project 65d609455c11505db4a12c76. Still provisioning (CREATING). Stop.
