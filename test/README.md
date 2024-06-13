# Test Inventory

#### Fields Description

- Command: the MongoDB Atlas CLI command without `atlas`
- E2E Atlas:
    -  Possible values: `('Y'|'N')`
    -  Indicates if an e2e test for the command is present `('Y')` or not present `('N')` for Atlas
 - Atlas:
     -  Possible values: `('Y'|'N'|' ')`
     -  Indicates if the command supports `('Y')` or not support `('N'|' ')` Atlas


#### Inventory

| Command                                                 | E2E Atlas | Atlas
| :-----------------------------------------------------: | :-------: | :---:
| `alerts config create`                                  | Y         | Y
| `alerts config delete`                                  | Y         | Y
| `alerts config fields type`                             | Y         | Y
| `alerts config list`                                    | Y         | Y
| `alerts acknowledge`                                    | Y         | Y
| `alerts unacknowledge`                                  | Y         | Y
| `alerts list`                                           | Y         | Y
| `alerts describe`                                       | Y         | Y
| `alerts global list`                                    |           |
| `accessList create`                                     | Y         | Y
| `accessList delete`                                     | Y         | Y
| `accessList list`                                       | Y         | Y
| `backup snapshots create`                               | Y         | Y
| `backup snapshots delete`                               | Y         | Y
| `backup snapshots describe`                             | Y         | Y
| `backup snapshots list`                                 | Y         | Y
| `backup snapshots watch`                                | Y         | Y
| `backup exports buckets create`                         | Y         | Y
| `backup exports buckets list`                           | Y         | Y
| `backup exports buckets describe`                       | Y         | Y
| `backup exports buckets delete`                         | Y         | Y
| `backup exports jobs create`                            | Y         | Y
| `backup exports jobs list`                              | Y         | Y
| `backup exports jobs describe`                          | Y         | Y
| `backup exports jobs delete`                            | Y         | Y
| `backup exports jobs watch`                             | Y         | Y
| `backup restores list`                                  | Y         | Y
| `backup restores describe`                              | Y         | Y
| `backup restores start`                                 | Y         | Y
| `backup restores watch`                                 | Y         | Y
| `backup schedule describe`                              | Y         | Y
| `backup schedule delete`                                | Y         | Y
| `backup schedule update`                                | Y         | Y
| `backup compliancepolicy describe`                      | Y         | Y
| `backup compliancepolicy enable`                        | Y         | Y
| `backup compliancepolicy copyprotection`                | Y         | Y
| `backup compliancepolicy setup`                         | Y         | Y
| `backup compliancepolicy pointintimerestore enable`     | Y         | Y
| `backup compliancepolicy policies describe`             | Y         | Y
| `backup compliancepolicy encryptionAtRest enable`       | N         | N
| `backup compliancepolicy encryptionAtRest disable`      | N         | N
| `cloudProvider aws accessRoles authorize`               | N         | Y
| `cloudProvider aws accessRoles deauthorize`             | N         | Y
| `cloudProvider aws accessRoles create`                  | Y         | Y
| `cloudProvider accessRoles list`                        | Y         | Y
| `cluster connectionString describe`                     | Y         | Y
| `cluster index create`                                  | Y         | Y
| `cluster create`                                        | Y         | Y
| `cluster delete`                                        | Y         | Y
| `cluster describe`                                      | Y         | Y
| `cluster list`                                          | Y         | Y
| `cluster start`                                         | N         | Y
| `cluster pause`                                         | N         | Y
| `cluster update`                                        | Y         | Y
| `cluster watch`                                         | Y         | Y
| `cluster onlineArchive create`                          | Y         | Y
| `cluster onlineArchive delete`                          | Y         | Y
| `cluster onlineArchive describe`                        | Y         | Y
| `cluster onlineArchive list`                            | Y         | Y
| `cluster onlineArchive pause`                           | Y         | Y
| `cluster onlineArchive start`                           | Y         | Y
| `cluster onlineArchive update`                          | Y         | Y
| `cluster search index create`                           | Y         | Y
| `cluster search index delete`                           | Y         | Y
| `cluster search index describe`                         | Y         | Y
| `cluster search index list`                             | Y         | Y
| `cluster search index update`                           | Y         | Y
| `cluster search nodes list`                             | Y         | Y
| `cluster search nodes create`                           | Y         | Y
| `cluster search nodes update`                           | Y         | Y
| `cluster search nodes delete`                           | Y         | Y
| `cluster advancedSettings describe`                     | Y         | Y
| `cluster advancedSettings update`                       | Y         | Y
| `dbrole create`                                         | Y         | Y
| `dbrole delete`                                         | Y         | Y
| `dbrole describe`                                       | Y         | Y
| `dbrole list`                                           | Y         | Y
| `dbrole update`                                         | Y         | Y
| `customDns aws describe`                                | Y         | Y
| `customDns aws disable`                                 | Y         | Y
| `customDns aws enable`                                  | Y         | Y
| `datalake create`                                       | Y         | Y
| `datalake delete`                                       | Y         | Y
| `datalake describe`                                     | Y         | Y
| `datalake list`                                         | Y         | Y
| `datalake update`                                       | Y         | Y
| `dbuser certs create`                                   | Y         | Y
| `dbuser certs list`                                     | Y         | Y
| `dbuser create`                                         | Y         | Y
| `dbuser delete`                                         | Y         | Y
| `dbuser describe`                                       | Y         | Y
| `dbuser list`                                           | Y         | Y
| `dbuser update`                                         | Y         | Y
| `integration create DATADOG`                            | Y         | Y
| `integration create FLOWDOCK`                           | Y         | Y
| `integration create OPS_GENIE`                          | Y         | Y
| `integration create PAGER_DUTY`                         | Y         | Y
| `integration create VICTOR_OPS`                         | Y         | Y
| `integration create WEBHOOK`                            | Y         | Y
| `integration create VICTOR_OPS`                         | Y         | Y
| `integration create VICTOR_OPS`                         | Y         | Y
| `integration delete`                                    | Y         | Y
| `integration describe`                                  | Y         | Y
| `integration list`                                      | Y         | Y
| `logs download`                                         | Y         | Y
| `accesslogs list`                                       | Y         | Y
| `maintenanceWindow clear`                               | Y         | Y
| `maintenanceWindow defer`                               | N         | Y
| `maintenanceWindow describe`                            | Y         | Y
| `maintenanceWindow update`                              | Y         | Y
| `metric database describe`                              | Y         | Y
| `metric database list`                                  | N         | Y
| `metric disk describe`                                  | Y         | Y
| `metric disk list`                                      | N         | Y
| `metric processes`                                      | Y         | Y
| `networking container delete`                           | N         | Y
| `networking container list`                             | N         | Y
| `networking peering create aws`                         | N         | Y
| `networking peering create azure`                       | N         | Y
| `networking peering create gcp`                         | N         | Y
| `networking peering delete`                             | N         | Y
| `networking peering list`                               | N         | Y
| `networking peering watch`                              | N         | Y
| `privateEndpoint aws interface create`                  | N         | Y
| `privateEndpoint aws interface delete`                  | N         | Y
| `privateEndpoint aws interface describe`                | N         | Y
| `privateEndpoint aws  create`                           | Y         | Y
| `privateEndpoint aws  delete`                           | Y         | Y
| `privateEndpoint aws  describe`                         | Y         | Y
| `privateEndpoint aws  list`                             | Y         | Y
| `privateEndpoint aws  watch`                            | Y         | Y
| `privateEndpoint azure interface create`                | N         | Y
| `privateEndpoint azure interface delete`                | N         | Y
| `privateEndpoint azure interface describe`              | N         | Y
| `privateEndpoint azure  create`                         | Y         | Y
| `privateEndpoint azure  delete`                         | Y         | Y
| `privateEndpoint azure  describe`                       | Y         | Y
| `privateEndpoint azure  list`                           | Y         | Y
| `privateEndpoint azure  watch`                          | Y         | Y
| `privateEndpoint gcp create`                            | Y         | Y
| `privateEndpoint gcp delete`                            | Y         | Y
| `privateEndpoint gcp list`                              | Y         | Y
| `privateEndpoint gcp descibe`                           | Y         | Y
| `privateEndpoint gcp watch`                             | Y         | Y
| `privateEndpoint dataLake aws create`                   | Y         | Y
| `privateEndpoint dataLake aws list`                     | Y         | Y
| `privateEndpoint dataLake aws delete`                   | Y         | Y
| `privateEndpoint dataLake aws describe`                 | Y         | Y
| `privateEndpoint interface create`                      | N         | Y
| `privateEndpoint interface delete`                      | N         | Y
| `privateEndpoint interface describe`                    | N         | Y
| `privateEndpoint regionalMode describe`                 | Y         | Y
| `privateEndpoint regionalMode enable`                   | Y         | Y
| `privateEndpoint regionalMode disable`                  | Y         | Y
| `process list`                                          | Y         | Y
| `quickstart`                                            | Y         | Y
| `security customercert create`                          | N         | Y
| `security customercert disable`                         | N         | Y
| `security customercert describe`                        | N         | Y
| `security ldap delete`                                  | Y         | Y
| `security ldap describe`                                | Y         | Y
| `security ldap save`                                    | Y         | Y
| `security ldap status`                                  | Y         | Y
| `security ldap verify`                                  | Y         | Y
| `security ldap watch`                                   | Y         | Y
| `streams`                                               |           |
| `streams connection`                                    |           |
| `streams connection create`                             | Y         | Y
| `streams connection delete`                             | Y         | Y
| `streams connection describe`                           | Y         | Y
| `streams connection list`                               | Y         | Y
| `streams connection update`                             | Y         | Y
| `streams instance`                                      |           |
| `streams instance create`                               | Y         | Y
| `streams instance delete`                               | Y         | Y
| `streams instance describe`                             | Y         | Y
| `streams instance list`                                 | Y         | Y
| `streams instance update`                               | Y         | Y
| `streams instance log`                                  | Y         | Y
| `config`                                                |           |
| `completion`                                            | Y         | Y
| `config delete`                                         | Y         | Y
| `config list`                                           | Y         | Y
| `config describe`                                       | Y         | Y
| `config rename`                                         | Y         | Y
| `config set`                                            |           |
| `event list`                                            | Y         | Y
| `globalAccessList create`                               |           |
| `globalAccessList delete`                               |           |
| `globalAccessList describe`                             |           |
| `globalAccessList list`                                 |           |
| `globalApiKey create`                                   |           |
| `globalApiKey delete`                                   |           |
| `globalApiKey describe`                                 |           |
| `globalApiKey list`                                     |           |
| `globalApiKey update`                                   |           |
| `orgs apiKey accessList create`                         | Y         | Y
| `orgs apiKey accessList delete`                         | Y         | Y
| `orgs apiKey accessList list`                           | Y         | Y
| `orgs apiKey  create`                                   | Y         | Y
| `orgs apiKey  delete`                                   | Y         | Y
| `orgs apiKey  describe`                                 | Y         | Y
| `orgs apiKey  list`                                     | Y         | Y
| `orgs apiKey  update`                                   | Y         | Y
| `orgs users  list`                                      | N         | Y
| `orgs create`                                           |           |
| `orgs delete`                                           | N         | Y
| `orgs describe`                                         | Y         | Y
| `orgs list`                                             | Y         | Y
| `orgs invitations list`                                 | Y         | Y
| `orgs invitations describe`                             | Y         | Y
| `orgs invitations invite`                               | Y         | Y
| `orgs invitations update`                               | Y         | Y
| `orgs invitations delete`                               | Y         | Y
| `project apiKey create`                                 | Y         | Y
| `project apiKey delete`                                 | Y         | Y
| `project apiKey describe`                               | N         | Y
| `project apiKey list`                                   | Y         | Y
| `project apiKey assign`                                 | Y         | Y
| `project users list`                                    | Y         | Y
| `project users delete`                                  | N         | Y
| `project team add`                                      | N         | Y
| `project team delete`                                   | N         | Y
| `project team list`                                     | N         | Y
| `project team update`                                   | N         | Y
| `project create`                                        | Y         | Y
| `project delete`                                        | N         | Y
| `project describe`                                      | N         | Y
| `project list`                                          | N         | Y
| `project invitations list`                              | Y         | Y
| `project invitations describe`                          | Y         | Y
| `project invitations invite`                            | Y         | Y
| `project invitations update`                            | Y         | Y
| `project invitations delete`                            | Y         | Y
| `team user add`                                         | Y         | Y
| `team user delete`                                      | Y         | Y
| `team user list`                                        | Y         | Y
| `team create`                                           | Y         | Y
| `team delete`                                           | Y         | Y
| `team describe`                                         | Y         | Y
| `team list`                                             | Y         | Y
| `user invite`                                           | N         | Y
| `user delete`                                           | N         | Y
| `user describe`                                         | Y         | Y
| `performanceAdvisor namespace list`                     | Y         | Y
| `performanceAdvisor slowQueryLogs list`                 | Y         | Y
| `performanceAdvisor suggestedIndexes list`              | Y         | Y
| `performanceAdvisor slowOT enable`                      | Y         | Y
| `performanceAdvisor slowOT disable`                     | Y         | Y
| `serverless create`                                     | Y         | Y
| `serverless delete`                                     | Y         | Y
| `serverless describe`                                   | Y         | Y
| `serverless list`                                       | Y         | Y
| `serverless watch`                                      | Y         | Y
| `serverless update`                                     | Y         | Y
| `serverless backup snapshots list`                      | Y         | Y
| `serverless backup snapshots describe`                  | Y         | Y
| `serverless backup snapshots watch`                     | Y         | Y
| `livemigrations link create`                            | Y         | Y
| `livemigrations link delete`                            | Y         | Y
| `livemigrations validation create`                      |           | Y
| `livemigrations create`                                 |           | Y
| `livemigrations cutover`                                | N         | Y
| `completion bash`                                       | Y         | Y
| `completion zsh`                                        | Y         | Y
| `completion fish`                                       | Y         | Y
| `completion powershell`                                 | Y         | Y
| `setup`                                                 | Y         | Y
| `register`                                              | N         | Y
| `kubernetes config generate`                            | Y         | Y
| `kubernetes config apply`                               | Y         | Y
| `kubernetes operator install`                           | Y         | Y
| `federatedAuthentication federationSettings describe`   | Y         | Y
