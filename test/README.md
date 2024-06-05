# Test Inventory

#### Fields Description

- Command: the MongoDB CLI command without `mongocli om|cm`
- E2E OM:
    -  Possible values: `('Y'|'N')`
    -  Indicates if an e2e test for the command is present `('Y')` or not present `('N')` for Ops Manager
- E2E CM:
    -  Possible values: `('Y'|'N')`
    -  Indicates if an e2e test for the command is present `('Y')` or not present `('N')` for Cloud Manager
 - OM:
     -  Possible values: `('Y'|'N'|' ')`
     -  Indicates if the command supports `('Y')` or not support `('N'|' ')` Ops Manager
 - CM:
     -  Possible values: `('Y'|'N'|' ')`
     -  Indicates if the command supports `('Y')` or not support `('N'|' ')` Cloud Manager


#### Inventory

| Command                                             | E2E OM | E2E CM | OM | CM |
|:----------------------------------------------------|:------:|:------:|:--:|:--:|
| `alerts config create`                              |   N    |   N    | Y  | Y  |
| `alerts config delete`                              |   N    |   N    | Y  | Y  |
| `alerts config fields type`                         |   N    |   N    | Y  | Y  |
| `alerts config list`                                |   N    |   N    | Y  | Y  |
| `alerts acknowledge`                                |   N    |   Y    | Y  | Y  |
| `alerts unacknowledge`                              |   N    |   Y    | Y  | Y  |
| `alerts list`                                       |   N    |   Y    | Y  | Y  |
| `alerts describe`                                   |   N    |   Y    | Y  | Y  |
| `alerts global list`                                |   N    |        | Y  |    |
| `accessList create`                                 |        |        |    |    |
| `accessList delete`                                 |        |        |    |    |
| `accessList list`                                   |        |        |    |    |
| `backup snapshots create`                           |        |        |    |    |
| `backup snapshots delete`                           |        |        |    |    |
| `backup snapshots describe`                         |        |        |    |    |
| `backup snapshots list`                             |        |        |    |    |
| `backup snapshots watch`                            |        |        |    |    |
| `backup exports buckets create`                     |        |        |    |    |
| `backup exports buckets list`                       |        |        |    |    |
| `backup exports buckets describe`                   |        |        |    |    |
| `backup exports buckets delete`                     |        |        |    |    |
| `backup exports jobs create`                        |        |        |    |    |
| `backup exports jobs list`                          |        |        |    |    |
| `backup exports jobs describe`                      |        |        |    |    |
| `backup exports jobs delete`                        |        |        |    |    |
| `backup exports jobs watch`                         |        |        |    |    |
| `backup restores list`                              |        |        |    |    |
| `backup restores describe`                          |        |        |    |    |
| `backup restores start`                             |        |        |    |    |
| `backup restores watch`                             |        |        |    |    |
| `backup schedule describe`                          |        |        |    |    |
| `backup schedule delete`                            |        |        |    |    |
| `backup schedule update`                            |        |        |    |    |
| `backup compliancepolicy describe`                  |        |        |    |    |
| `backup compliancepolicy enable`                    |        |        |    |    |
| `backup compliancepolicy copyprotection`            |        |        |    |    |
| `backup compliancepolicy setup`                     |        |        |    |    |
| `backup compliancepolicy pointintimerestore enable` |        |        |    |    |
| `backup compliancepolicy policies describe`         |        |        |    |    |
| `backup compliancepolicy encryptionAtRest enable`   |        |        |    |    |
| `backup compliancepolicy encryptionAtRest disable`  |        |        |    |    |
| `cloudProvider aws accessRoles authorize`           |        |        |    |    |
| `cloudProvider aws accessRoles deauthorize`         |        |        |    |    |
| `cloudProvider aws accessRoles create`              |        |        |    |    |
| `cloudProvider accessRoles list`                    |        |        |    |    |
| `cluster connectionString describe`                 |        |        |    |    |
| `cluster index create`                              |        |        |    |    |
| `cluster create`                                    |        |        |    |    |
| `cluster delete`                                    |        |        |    |    |
| `cluster describe`                                  |        |        |    |    |
| `cluster list`                                      |        |        |    |    |
| `cluster start`                                     |        |        |    |    |
| `cluster pause`                                     |        |        |    |    |
| `cluster update`                                    |   Y    |   Y    | Y  | Y  |
| `cluster watch`                                     |        |        |    |    |
| `cluster onlineArchive create`                      |        |        | N  | N  |
| `cluster onlineArchive delete`                      |        |        | N  | N  |
| `cluster onlineArchive describe`                    |        |        | N  | N  |
| `cluster onlineArchive list`                        |        |        | N  | N  |
| `cluster onlineArchive pause`                       |        |        | N  | N  |
| `cluster onlineArchive start`                       |        |        | N  | N  |
| `cluster onlineArchive update`                      |        |        | N  | N  |
| `cluster search index create`                       |        |        | N  | N  |
| `cluster search index delete`                       |        |        | N  | N  |
| `cluster search index describe`                     |        |        | N  | N  |
| `cluster search index list`                         |        |        | N  | N  |
| `cluster search index update`                       |        |        | N  | N  |
| `cluster search nodes list`                         |        |        | N  | N  |
| `cluster search nodes create`                       |        |        | N  | N  |
| `cluster search nodes update`                       |        |        | N  | N  |
| `cluster search nodes delete`                       |        |        | N  | N  |
| `cluster advancedSettings describe`                 |        |        | N  | N  |
| `cluster advancedSettings update`                   |        |        | N  | N  |
| `dbrole create`                                     |        |        |    |    |
| `dbrole delete`                                     |        |        |    |    |
| `dbrole describe`                                   |        |        |    |    |
| `dbrole list`                                       |        |        |    |    |
| `dbrole update`                                     |        |        |    |    |
| `customDns aws describe`                            |        |        |    |    |
| `customDns aws disable`                             |        |        |    |    |
| `customDns aws enable`                              |        |        |    |    |
| `datalake create`                                   |        |        |    |    |
| `datalake delete`                                   |        |        |    |    |
| `datalake describe`                                 |        |        |    |    |
| `datalake list`                                     |        |        |    |    |
| `datalake update`                                   |        |        |    |    |
| `dbuser certs create`                               |        |        |    |    |
| `dbuser certs list`                                 |        |        |    |    |
| `dbuser create`                                     |        |        |    |    |
| `dbuser delete`                                     |        |        |    |    |
| `dbuser describe`                                   |        |        |    |    |
| `dbuser list`                                       |        |        |    |    |
| `dbuser update`                                     |        |        |    |    |
| `integration create DATADOG`                        |        |        |    |    |
| `integration create FLOWDOCK`                       |        |        |    |    |
| `integration create OPS_GENIE`                      |        |        |    |    |
| `integration create PAGER_DUTY`                     |        |        |    |    |
| `integration create VICTOR_OPS`                     |        |        |    |    |
| `integration create WEBHOOK`                        |        |        |    |    |
| `integration create VICTOR_OPS`                     |        |        |    |    |
| `integration create VICTOR_OPS`                     |        |        |    |    |
| `integration delete`                                |        |        |    |    |
| `integration describe`                              |        |        |    |    |
| `integration list`                                  |        |        |    |    |
| `logs download`                                     |        |        |    |    |
| `accesslogs list`                                   |        |        |    |    |
| `maintenanceWindow clear`                           |        |        |    |    |
| `maintenanceWindow defer`                           |        |        |    |    |
| `maintenanceWindow describe`                        |        |        |    |    |
| `maintenanceWindow update`                          |        |        |    |    |
| `metric database describe`                          |        |        |    |    |
| `metric database list`                              |        |        |    |    |
| `metric disk describe`                              |        |        |    |    |
| `metric disk list`                                  |        |        |    |    |
| `metric processes`                                  |        |        |    |    |
| `networking container delete`                       |        |        |    |    |
| `networking container list`                         |        |        |    |    |
| `networking peering create aws`                     |        |        |    |    |
| `networking peering create azure`                   |        |        |    |    |
| `networking peering create gcp`                     |        |        |    |    |
| `networking peering delete`                         |        |        |    |    |
| `networking peering list`                           |        |        |    |    |
| `networking peering watch`                          |        |        |    |    |
| `privateEndpoint aws interface create`              |        |        |    |    |
| `privateEndpoint aws interface delete`              |        |        |    |    |
| `privateEndpoint aws interface describe`            |        |        |    |    |
| `privateEndpoint aws  create`                       |        |        |    |    |
| `privateEndpoint aws  delete`                       |        |        |    |    |
| `privateEndpoint aws  describe`                     |        |        |    |    |
| `privateEndpoint aws  list`                         |        |        |    |    |
| `privateEndpoint aws  watch`                        |        |        |    |    |
| `privateEndpoint azure interface create`            |        |        |    |    |
| `privateEndpoint azure interface delete`            |        |        |    |    |
| `privateEndpoint azure interface describe`          |        |        |    |    |
| `privateEndpoint azure  create`                     |        |        |    |    |
| `privateEndpoint azure  delete`                     |        |        |    |    |
| `privateEndpoint azure  describe`                   |        |        |    |    |
| `privateEndpoint azure  list`                       |        |        |    |    |
| `privateEndpoint azure  watch`                      |        |        |    |    |
| `privateEndpoint gcp create`                        |        |        |    |    |
| `privateEndpoint gcp delete`                        |        |        |    |    |
| `privateEndpoint gcp list`                          |        |        |    |    |
| `privateEndpoint gcp descibe`                       |        |        |    |    |
| `privateEndpoint gcp watch`                         |        |        |    |    |
| `privateEndpoint dataLake aws create`               |        |        |    |    |
| `privateEndpoint dataLake aws list`                 |        |        |    |    |
| `privateEndpoint dataLake aws delete`               |        |        |    |    |
| `privateEndpoint dataLake aws describe`             |        |        |    |    |
| `privateEndpoint interface create`                  |        |        |    |    |
| `privateEndpoint interface delete`                  |        |        |    |    |
| `privateEndpoint interface describe`                |        |        |    |    |
| `privateEndpoint regionalMode describe`             |        |        |    |    |
| `privateEndpoint regionalMode enable`               |        |        |    |    |
| `privateEndpoint regionalMode disable`              |        |        |    |    |
| `process list`                                      |        |        |    |    |
| `quickstart`                                        |        |        |    |    |
| `security customercert create`                      |        |        |    |    |
| `security customercert disable`                     |        |        |    |    |
| `security customercert describe`                    |        |        |    |    |
| `security ldap delete`                              |        |        |    |    |
| `security ldap describe`                            |        |        |    |    |
| `security ldap save`                                |        |        |    |    |
| `security ldap status`                              |        |        |    |    |
| `security ldap verify`                              |        |        |    |    |
| `security ldap watch`                               |        |        |    |    |
| `streams`                                           |        |        |    |    |
| `streams connection`                                |        |        |    |    |
| `streams connection create`                         |        |        |    |    |
| `streams connection delete`                         |        |        |    |    |
| `streams connection describe`                       |        |        |    |    |
| `streams connection list`                           |        |        |    |    |
| `streams connection update`                         |        |        |    |    |
| `streams instance`                                  |        |        |    |    |
| `streams instance create`                           |        |        |    |    |
| `streams instance delete`                           |        |        |    |    |
| `streams instance describe`                         |        |        |    |    |
| `streams instance list`                             |        |        |    |    |
| `streams instance update`                           |        |        |    |    |
| `config`                                            |        |        |    |    |
| `completion`                                        |   Y    |   Y    | Y  | Y  |
| `config delete`                                     |   Y    |   Y    | Y  | Y  |
| `config list`                                       |   Y    |   Y    | Y  | Y  |
| `config describe`                                   |   Y    |   Y    | Y  | Y  |
| `config rename`                                     |   Y    |   Y    | Y  | Y  |
| `config set`                                        |        |        |    |    |
| `event list`                                        |   N    |   Y    | Y  | Y  |
| `iam globalAccessList create`                       |   N    |        | Y  |    |
| `iam globalAccessList delete`                       |   N    |        | Y  |    |
| `iam globalAccessList describe`                     |   N    |        | Y  |    |
| `iam globalAccessList list`                         |   N    |        | Y  |    |
| `iam globalApiKey create`                           |   N    |        | Y  |    |
| `iam globalApiKey delete`                           |   N    |        | Y  |    |
| `iam globalApiKey describe`                         |   N    |        | Y  |    |
| `iam globalApiKey list`                             |   N    |        | Y  |    |
| `iam globalApiKey update`                           |   N    |        | Y  |    |
| `iam orgs apiKey accessList create`                 |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey accessList delete`                 |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey accessList list`                   |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey  create`                           |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey  delete`                           |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey  describe`                         |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey  list`                             |   Y    |   Y    | Y  | Y  |
| `iam orgs apiKey  update`                           |   Y    |   Y    | Y  | Y  |
| `iam orgs users  list`                              |   N    |   N    | Y  | Y  |
| `iam orgs create`                                   |   Y    |   N    | Y  | Y  |
| `iam orgs delete`                                   |   Y    |   N    | Y  | Y  |
| `iam orgs describe`                                 |   Y    |   N    | Y  | Y  |
| `iam orgs list`                                     |   Y    |   N    | Y  | Y  |
| `iam orgs invitations list`                         |   Y    |   N    | N  | N  |
| `iam orgs invitations describe`                     |   Y    |   N    | N  | N  |
| `iam orgs invitations invite`                       |   Y    |   N    | N  | N  |
| `iam orgs invitations update`                       |   Y    |   N    | N  | N  |
| `iam orgs invitations delete`                       |   Y    |   N    | N  | N  |
| `iam project apiKey create`                         |   N    |   Y    | Y  | Y  |
| `iam project apiKey delete`                         |   N    |   Y    | Y  | Y  |
| `iam project apiKey describe`                       |   N    |   N    | Y  | Y  |
| `iam project apiKey list`                           |   N    |   Y    | Y  | Y  |
| `iam project apiKey assign`                         |   N    |   Y    | Y  | Y  |
| `iam project users list`                            |   N    |   Y    | Y  | Y  |
| `iam project users delete`                          |   N    |   N    | Y  | Y  |
| `iam project team add`                              |   N    |   N    | Y  | Y  |
| `iam project team delete`                           |   N    |   N    | Y  | Y  |
| `iam project team list`                             |   N    |   N    | Y  | Y  |
| `iam project team update`                           |   N    |   N    | Y  | Y  |
| `iam project create`                                |   Y    |   Y    | Y  | Y  |
| `iam project delete`                                |   Y    |   N    | Y  | Y  |
| `iam project describe`                              |   Y    |   N    | Y  | Y  |
| `iam project list`                                  |   Y    |   N    | Y  | Y  |
| `iam project invitations list`                      |   Y    |   N    | N  | N  |
| `iam project invitations describe`                  |   Y    |   N    | N  | N  |
| `iam project invitations invite`                    |   Y    |   N    | N  | N  |
| `iam project invitations update`                    |   Y    |   N    | N  | N  |
| `iam project invitations delete`                    |   Y    |   N    | N  | N  |
| `iam team user add`                                 |   N    |   Y    | Y  | Y  |
| `iam team user delete`                              |   N    |   Y    | Y  | Y  |
| `iam team user list`                                |   N    |   Y    | Y  | Y  |
| `iam team create`                                   |   N    |   Y    | Y  | Y  |
| `iam team delete`                                   |   N    |   Y    | Y  | Y  |
| `iam team describe`                                 |   N    |   Y    | Y  | Y  |
| `iam team list`                                     |   N    |   Y    | Y  | Y  |
| `iam user invite`                                   |   N    |   N    | Y  | Y  |
| `iam user delete`                                   |   N    |   N    | Y  | Y  |
| `iam user describe`                                 |   N    |   Y    | Y  | Y  |
| `admin backup blockstore create`                    |   N    |        | N  |    |
| `admin backup blockstore delete`                    |   N    |        | N  |    |
| `admin backup blockstore describe`                  |   N    |        | N  |    |
| `admin backup blockstore list`                      |   N    |        | N  |    |
| `admin backup blockstore update`                    |   N    |        | N  |    |
| `admin backup fileSystem create`                    |   N    |        | N  |    |
| `admin backup fileSystem delete`                    |   N    |        | N  |    |
| `admin backup fileSystem describe`                  |   N    |        | N  |    |
| `admin backup fileSystem list`                      |   N    |        | N  |    |
| `admin backup fileSystem update`                    |   N    |        | N  |    |
| `admin backup oplog create`                         |   N    |        | N  |    |
| `admin backup oplog delete`                         |   N    |        | N  |    |
| `admin backup oplog describe`                       |   N    |        | N  |    |
| `admin backup oplog list`                           |   N    |        | N  |    |
| `admin backup oplog update`                         |   N    |        | N  |    |
| `admin backup s3 create`                            |   N    |        | N  |    |
| `admin backup s3 delete`                            |   N    |        | N  |    |
| `admin backup s3 describe`                          |   N    |        | N  |    |
| `admin backup s3 list`                              |   N    |        | N  |    |
| `admin backup s3 update`                            |   N    |        | N  |    |
| `admin backup sync create`                          |   N    |        | N  |    |
| `admin backup sync delete`                          |   N    |        | N  |    |
| `admin backup sync describe`                        |   N    |        | N  |    |
| `admin backup sync list`                            |   N    |        | N  |    |
| `admin backup sync update`                          |   N    |        | N  |    |
| `agent apikeys create`                              |   N    |   N    | Y  | Y  |
| `agent apikeys delete`                              |   N    |   N    | Y  | Y  |
| `agent apikeys list`                                |   Y    |   Y    | Y  | Y  |
| `agent version list`                                |   Y    |   Y    | Y  | Y  |
| `agent list`                                        |   Y    |   Y    | Y  | Y  |
| `agent upgrade`                                     |   Y    |   Y    | Y  | Y  |
| `automation describe`                               |   N    |   N    | Y  | Y  |
| `automation status`                                 |   N    |   N    | Y  | Y  |
| `automation update`                                 |   N    |   N    | Y  | Y  |
| `automation watch`                                  |   N    |   N    | Y  | Y  |
| `cluster  apply`                                    |   Y    |   Y    | Y  | Y  |
| `cluster  create`                                   |   Y    |   Y    | Y  | Y  |
| `cluster  delete`                                   |   Y    |   Y    | Y  | Y  |
| `cluster  describe`                                 |   Y    |   Y    | Y  | Y  |
| `cluster  list`                                     |   Y    |   Y    | Y  | Y  |
| `cluster  restart`                                  |   Y    |   Y    | Y  | Y  |
| `cluster  shutdown`                                 |   Y    |   Y    | Y  | Y  |
| `cluster  startup`                                  |   N    |   N    | Y  | Y  |
| `cluster  unmanage`                                 |   Y    |   Y    | Y  | Y  |
| `cluster  update`                                   |   N    |   N    | Y  | Y  |
| `cluster  reclaimFreeSpace`                         |   Y    |   Y    | Y  | Y  |
| `cluster  index create`                             |   N    |   N    | Y  | Y  |
| `dbuser create`                                     |   N    |   Y    | Y  | Y  |
| `dbuser delete`                                     |   N    |   Y    | Y  | Y  |
| `dbuser list`                                       |   N    |   Y    | Y  | Y  |
| `diagnose-archive download`                         |   N    |   N    | Y  | Y  |
| `featurePolicy list`                                |   Y    |   Y    | Y  | Y  |
| `featurePolicy update`                              |   Y    |   Y    | Y  | Y  |
| `logs jobs collect`                                 |   N    |   N    | Y  | Y  |
| `logs jobs download`                                |   N    |   N    | Y  | Y  |
| `logs jobs delete`                                  |   N    |   N    | Y  | Y  |
| `logs jobs list`                                    |   N    |   N    | Y  | Y  |
| `logs decrypt`                                      |        |   Y    |    | Y  |
| `logs keyProvider list`                             |        |   Y    |    | Y  |
| `maintenanceWindows create`                         |   Y    |   Y    | Y  | Y  |
| `maintenanceWindows delete`                         |   Y    |   Y    | Y  | Y  |
| `maintenanceWindows list`                           |   Y    |   Y    | Y  | Y  |
| `maintenanceWindows update`                         |   Y    |   Y    | Y  | Y  |
| `database describe`                                 |   N    |   N    | Y  | Y  |
| `database list`                                     |   N    |   N    | Y  | Y  |
| `disk describe`                                     |   N    |   N    | Y  | Y  |
| `disk list`                                         |   N    |   N    | Y  | Y  |
| `process`                                           |   N    |   N    | Y  | Y  |
| `process describe`                                  |   N    |   N    | Y  | Y  |
| `process list`                                      |   N    |   N    | Y  | Y  |
| `security enable`                                   |   Y    |   Y    | Y  | Y  |
| `monitoring enable`                                 |   Y    |   Y    | Y  | Y  |
| `monitoring disable`                                |   Y    |   Y    | Y  | Y  |
| `monitoring stop`                                   |   Y    |   Y    | Y  | Y  |
| `owner create`                                      |   Y    |        | Y  |    |
| `server list`                                       |   Y    |   Y    | Y  | Y  |
| `softwareComponent list`                            |   N    |   N    | Y  | Y  |
| `versionManifest update`                            |   N    |        | Y  |    |
| `serverUsage orgs hosts list`                       |   N    |        | Y  |    |
| `serverUsage orgs serverType get`                   |   N    |        | Y  |    |
| `serverUsage orgs serverType set`                   |   N    |        | Y  |    |
| `serverUsage project hosts list`                    |   N    |        | Y  |    |
| `serverUsage project serverType get`                |   N    |        | Y  |    |
| `serverUsage projet serverType set`                 |   N    |        | Y  |    |
| `serverUsage capture`                               |   N    |        | Y  |    |
| `serverUsage download`                              |   N    |        | Y  |    |
| `performanceAdvisor namespace list`                 |   N    |   N    | Y  | Y  |
| `performanceAdvisor slowQueryLogs list`             |   N    |   N    | Y  | Y  |
| `performanceAdvisor suggestedIndexes list`          |   N    |   N    | Y  | Y  |
| `performanceAdvisor slowOT enable`                  |        |        | N  | N  |
| `performanceAdvisor slowOT disable`                 |        |        | N  | N  |
| `serverless create`                                 |        |        |    |    |
| `serverless delete`                                 |        |        |    |    |
| `serverless describe`                               |        |        |    |    |
| `serverless list`                                   |        |        |    |    |
| `serverless watch`                                  |        |        |    |    |
| `serverless update`                                 |        |        |    |    |
| `serverless backup snapshots list`                  |        |        |    |    |
| `serverless backup snapshots describe`              |        |        |    |    |
| `serverless backup snapshots watch`                 |        |        |    |    |
| `livemigrations link create`                        |        |        | N  | N  |
| `livemigrations link delete`                        |        |        | Y  | Y  |
| `livemigrations validation create`                  |        |        |    |    |
| `livemigrations create`                             |        |        |    |    |
| `livemigrations cutover`                            |        |        | N  | N  |
| `completion bash`                                   |   Y    |   Y    | Y  | Y  |
| `completion zsh`                                    |   Y    |   Y    | Y  | Y  |
| `completion fish`                                   |   Y    |   Y    | Y  | Y  |
| `completion powershell`                             |   Y    |   Y    | Y  | Y  |
| `setup`                                             |        |        |    |    |
| `register`                                          |        |        |    |    |
| `kubernetes config generate`                        |        |        | N  |    |
| `kubernetes config apply`                           |        |        | N  |    |
| `kubernetes operator install`                       |        |        | N  |    |
