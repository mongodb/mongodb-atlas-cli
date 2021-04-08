# E2E Test Inventory

#### Fields Description

- Command: the MongoDB CLI command without `mongocli atlas|om|cm` 
- E2E Atlas:    
    -  `Y`: the e2e test for the related command has been implemented for Atlas
    -  `N`: no e2e test for the related command has been implemented for Atlas
- E2E OM:    
    -  `Y`: the e2e test for the related command has been implemented for Ops Manager
    -  `N`: no e2e test for the related command has been implemented for Ops Manager
- E2E CM:    
    -  `Y`: the e2e test for the related command has been implemented for Cloud Manager
    -  `N`: no e2e test for the related command has been implemented for Cloud Manager
 - Atlas:    
     -  `Y`: the related command works with Atlas
     -  `N`: the related command does't work with Atlas (you can leave `E2E Atlas` empty)
 - OM:    
     -  `Y`: the related command works with Ops Manager
     -  `N`: the related command does't work with Ops Manager (you can leave `E2E OM` empty)
 - CM:    
     -  `Y`: the related command works with Cloud Manager
     -  `N`: the related command does't work with Cloud Manager (you can leave `E2E CM` empty)


#### Inventory

Command                                         | E2E Atlas | E2E OM    | E2E CM    | Atlas     | OM    | CM    |
| :---                                          | :---:     | :---:     | :---:     | :---:     | :---: | :---: |
| `alerts config create`                        | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config delete`                        | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config fields type`                   | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config list`                          | Y         | N         | N         | Y         | Y     | Y     |
| `alerts acknowledge`                          | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts unacknowledge`                        | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts list`                                 | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts describe`                             | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts global list`                          |           | N         |           |           | Y     |       |
| `accessList create`                           | Y         |           |           | Y         |       |       |
| `accessList delete`                           | Y         |           |           | Y         |       |       |
| `accessList list`                             | Y         |           |           | Y         |       |       |
| `backup snapshots create`                     | N         |           |           | Y         |       |       |
| `backup snapshots delete`                     | N         |           |           | Y         |       |       |
| `backup snapshots describe`                   | N         |           |           | Y         |       |       |
| `backup snapshots watch`                      | N         |           |           | Y         |       |       |
| `backup restore list`                         | N         |           |           | Y         |       |       |
| `backup restore start`                        | N         |           |           | Y         |       |       |
| `cloudProvider aws accessRoles authorize`     | N         |           |           | Y         |       |       |
| `cloudProvider aws accessRoles deauthorize`   | N         |           |           | Y         |       |       |
| `cloudProvider aws accessRoles create`        | Y         |           |           | Y         |       |       |
| `cloudProvider accessRoles list`              | Y         |           |           | Y         |       |       |
| `cluster connectionString describe`           | Y         |           |           | Y         |       |       |
| `cluster index create`                        | Y         |           |           | Y         |       |       |
| `cluster create`                              | Y         |           |           | Y         |       |       |
| `cluster delete`                              | Y         |           |           | Y         |       |       |
| `cluster describe`                            | Y         |           |           | Y         |       |       |
| `cluster list`                                | Y         |           |           | Y         |       |       |
| `cluster start`                               | N         |           |           | Y         |       |       |
| `cluster pause`                               | N         |           |           | Y         |       |       |
| `cluster update`                              | Y         |           |           | Y         |       |       |
| `cluster watch`                               | Y         |           |           | Y         |       |       |
| `cluster onlineArchive create`                | N         |           |           | Y         |       |       |
| `cluster onlineArchive delete`                | N         |           |           | Y         |       |       |
| `cluster onlineArchive describe`              | N         |           |           | Y         |       |       |
| `cluster onlineArchive list`                  | N         |           |           | Y         |       |       |
| `cluster onlineArchive pause`                 | N         |           |           | Y         |       |       |
| `cluster onlineArchive start`                 | N         |           |           | Y         |       |       |
| `cluster onlineArchive update`                | N         |           |           | Y         |       |       |
| `cluster search index create`                 | N         |           |           | Y         |       |       |
| `cluster search index delete`                 | N         |           |           | Y         |       |       |
| `cluster search index describe`               | N         |           |           | Y         |       |       |
| `cluster search index list`                   | N         |           |           | Y         |       |       |
| `cluster search index update`                 | N         |           |           | Y         |       |       |
| `dbrole create`                               | Y         |           |           | Y         |       |       |
| `dbrole delete`                               | Y         |           |           | Y         |       |       |
| `dbrole describe`                             | Y         |           |           | Y         |       |       |
| `dbrole list`                                 | Y         |           |           | Y         |       |       |
| `dbrole update`                               | Y         |           |           | Y         |       |       |
| `customDns aws describe`                      | Y         |           |           | Y         |       |       |
| `customDns aws disable`                       | Y         |           |           | Y         |       |       |
| `customDns aws enable`                        | Y         |           |           | Y         |       |       |
| `datalake create`                             | Y         |           |           | Y         |       |       |
| `datalake delete`                             | Y         |           |           | Y         |       |       |
| `datalake describe`                           | Y         |           |           | Y         |       |       |
| `datalake list`                               | Y         |           |           | Y         |       |       |
| `datalake update`                             | Y         |           |           | Y         |       |       |
| `dbuser certs create`                         | Y         |           |           | Y         |       |       |
| `dbuser certs list`                           | Y         |           |           | Y         |       |       |
| `dbuser create`                               | Y         |           |           | Y         |       |       |
| `dbuser delete`                               | Y         |           |           | Y         |       |       |
| `dbuser describe`                             | Y         |           |           | Y         |       |       |
| `dbuser list`                                 | Y         |           |           | Y         |       |       |
| `dbuser update`                               | Y         |           |           | Y         |       |       |
| `integration create DATADOG`                  | Y         |           |           | Y         |       |       |
| `integration create FLOWDOCK`                 | Y         |           |           | Y         |       |       |
| `integration create NEW_RELIC`                | Y         |           |           | Y         |       |       |
| `integration create OPS_GENIE`                | Y         |           |           | Y         |       |       |
| `integration create PAGER_DUTY`               | Y         |           |           | Y         |       |       |
| `integration create VICTOR_OPS`               | Y         |           |           | Y         |       |       |
| `integration create WEBHOOK`                  | Y         |           |           | Y         |       |       |
| `integration create VICTOR_OPS`               | Y         |           |           | Y         |       |       |
| `integration create VICTOR_OPS`               | Y         |           |           | Y         |       |       |
| `integration delete`                          | Y         |           |           | Y         |       |       |
| `integration describe`                        | Y         |           |           | Y         |       |       |
| `integration list`                            | Y         |           |           | Y         |       |       |
| `logs download`                               | Y         |           |           | Y         |       |       |
| `maintenanceWindow clear`                     | Y         |           |           | Y         |       |       |
| `maintenanceWindow defer`                     | N         |           |           | Y         |       |       |
| `maintenanceWindow describe`                  | Y         |           |           | Y         |       |       |
| `maintenanceWindow update`                    | Y         |           |           | Y         |       |       |
| `metric database describe`                    | Y         |           |           | Y         |       |       |
| `metric database list`                        | N         |           |           | Y         |       |       |
| `metric disk describe`                        | Y         |           |           | Y         |       |       |
| `metric disk list`                            | N         |           |           | Y         |       |       |
| `metric processes`                            | Y         |           |           | Y         |       |       |
| `networking container delete`                 | N         |           |           | Y         |       |       |
| `networking container list`                   | N         |           |           | Y         |       |       |
| `networking peering create aws`               | N         |           |           | Y         |       |       |
| `networking peering create azure`             | N         |           |           | Y         |       |       |
| `networking peering create gcp`               | N         |           |           | Y         |       |       |
| `networking peering delete`                   | N         |           |           | Y         |       |       |
| `networking peering list`                     | N         |           |           | Y         |       |       |
| `networking peering watch`                    | N         |           |           | Y         |       |       |
| `privateEndpoint aws interface create`        | N         |           |           | Y         |       |       |
| `privateEndpoint aws interface delete`        | N         |           |           | Y         |       |       |
| `privateEndpoint aws interface describe`      | N         |           |           | Y         |       |       |
| `privateEndpoint aws  create`                 | Y         |           |           | Y         |       |       |
| `privateEndpoint aws  delete`                 | Y         |           |           | Y         |       |       |
| `privateEndpoint aws  describe`               | Y         |           |           | Y         |       |       |
| `privateEndpoint aws  list`                   | Y         |           |           | Y         |       |       |
| `privateEndpoint aws  watch`                  | Y         |           |           | Y         |       |       |
| `privateEndpoint azure interface create`      | N         |           |           | Y         |       |       |
| `privateEndpoint azure interface delete`      | N         |           |           | Y         |       |       |
| `privateEndpoint azure interface describe`    | N         |           |           | Y         |       |       |
| `privateEndpoint azure  create`               | Y         |           |           | Y         |       |       |
| `privateEndpoint azure  delete`               | Y         |           |           | Y         |       |       |
| `privateEndpoint azure  describe`             | Y         |           |           | Y         |       |       |
| `privateEndpoint azure  list`                 | Y         |           |           | Y         |       |       |
| `privateEndpoint azure  watch`                | Y         |           |           | Y         |       |       |
| `privateEndpoint interface create`            | N         |           |           | Y         |       |       |
| `privateEndpoint interface delete`            | N         |           |           | Y         |       |       |
| `privateEndpoint interface describe`          | N         |           |           | Y         |       |       |
| `privateEndpoint  create`                     | Y         |           |           | Y         |       |       |
| `privateEndpoint  delete`                     | Y         |           |           | Y         |       |       |
| `privateEndpoint  describe`                   | Y         |           |           | Y         |       |       |
| `privateEndpoint  list`                       | Y         |           |           | Y         |       |       |
| `privateEndpoint  watch`                      | Y         |           |           | Y         |       |       |
| `privateEndpoint  regionalMode describe`      | Y         |           |           | Y         |       |       |
| `privateEndpoint  regionalMode enable`        | Y         |           |           | Y         |       |       |
| `privateEndpoint  regionalMode disable`       | Y         |           |           | Y         |       |       |
| `process list`                                | Y         |           |           | Y         |       |       |
| `quickstart`                                  | Y         |           |           | Y         |       |       
| `security customercert create`                | N         |           |           | Y         |       |       |
| `security customercert disable`               | N         |           |           | Y         |       |       |
| `security customercert describe`              | N         |           |           | Y         |       |       |
| `security ldap delete`                        | Y         |           |           | Y         |       |       |
| `security ldap describe`                      | Y         |           |           | Y         |       |       |
| `security ldap save`                          | Y         |           |           | Y         |       |       |
| `security ldap status`                        | Y         |           |           | Y         |       |       |
| `security ldap verify`                        | Y         |           |           | Y         |       |       |
| `security ldap watch`                         | Y         |           |           | Y         |       |       |
| `config`                                      |           |           |           |           |       |       |
| `config delete`                               |           |           |           |           |       |       |
| `config list`                                 |           |           |           |           |       |       |
| `config describe`                             |           |           |           |           |       |       |
| `config rename`                               |           |           |           |           |       |       |
| `config set`                                  |           |           |           |           |       |       |
| `event list`                                  | Y         | N         |Y          | Y         | Y     | Y     |
| `iam globalAccessList create`                 |           | N         |           |           | N     |       |
| `iam globalAccessList delete`                 |           | N         |           |           | N     |       |
| `iam globalAccessList describe`               |           | N         |           |           | N     |       |
| `iam globalAccessList list`                   |           | N         |           |           | N     |       |
| `iam globalApiKey create`                     |           | N         |           |           | N     |       |
| `iam globalApiKey delete`                     |           | N         |           |           | N     |       |
| `iam globalApiKey describe`                   |           | N         |           |           | N     |       |
| `iam globalApiKey list`                       |           | N         |           |           | N     |       |
| `iam globalApiKey update`                     |           | N         |           |           | N     |       |