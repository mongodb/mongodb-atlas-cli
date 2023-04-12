# Atlas CLI Generator

:warning: PREVIEW :warning: this tools is still in preview and it might be not suitable for all use cases.

## Usage

Open [spec.yaml](spec.yaml) and add to your needs later run `make gen-code`.

### Adding Stores

Add a new entry like this:

```yaml
  - base_file_name: data_lake_pipeline # this will be internal/store/data_lake_pipeline.go
    template: store # only template available for stores
    creator: # possible versions include `creator`, `updater`, `describer`, `lister` and `deleter`
      name: PipelineCreator # interface name
      method: CreatePipeline # interface func
      sdk_method: DataLakePipelinesApi.CreatePipeline # API from SDKv2 to call
      arg_type: 'atlasv2.IngestionPipeline' # Argument type
      arg_name: IngestionPipeline # Argument method to call on SDK
      return_type: '*atlasv2.IngestionPipeline' # Type the SDK returns
```

### Adding Commands

Add a new entry like this:

```yaml
  - command_path: atlas dataLake pipeline # parent command path
    package_name: pipeline # package name for the folder
    description: Data Lake Pipelines. # description of the command
    template: parent # template can be 'parent', 'list', 'describe', 'create', 'update' or 'delete'
    sub_commands: # parent commands normally have subcommands
      - command_path: atlas dataLake pipeline create # child command path
        package_name: pipeline # package name for the folder (should match parent command)
        store_name: PipelineCreator # name of the store (interface name)
        store_method: CreatePipeline # store method (interface func)
        description: Creates a new Data Lake Pipeline. To use this resource, the requesting API Key must have the Project Owner role. # description of the command
        template: create # template to be used by the command
        output_template: Pipeline {{ .Name }} created. # template to be used by the command (used by templates 'list', 'describe', 'create' and 'update')
        id_name: pipelineName # id passed to the command (used by templates 'describe', 'delete', 'create' and 'update')
        id_description: Name of the pipeline # description of id passed to the command (used by templates 'describe', 'delete', 'create' and 'update')
        request_type: 'atlasv2.IngestionPipeline' # type used to compose the entity passed to the SDK (used by templates 'create' and 'update')
```
