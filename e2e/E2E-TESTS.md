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

Command                         | E2E Atlas | E2E OM    | E2E CM    | Atlas     | OM    | CM    |
| :---                          | :---:     | :---:     | :---:     | :---:     | :---: | :---: |
| `alerts config create`        | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config delete`        | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config fields type`   | Y         | N         | N         | Y         | Y     | Y     |
| `alerts config list`          | Y         | N         | N         | Y         | Y     | Y     |
| `alerts acknowledge`          | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts unacknowledge`        | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts list`                 | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts describe`             | Y         | N         | Y         | Y         | Y     | Y     |
| `alerts global list`          |           | N         |           |           | Y     |       |
