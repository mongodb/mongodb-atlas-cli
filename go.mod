module github.com/mongodb/mongodb-atlas-cli/atlascli

go 1.23.1

require (
	cloud.google.com/go/kms v1.20.1
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.16.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.0
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys v1.3.0
	github.com/Masterminds/semver/v3 v3.3.1
	github.com/Netflix/go-expect v0.0.0-20220104043353-73e0943537d2
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/aws/aws-sdk-go-v2 v1.32.6
	github.com/aws/aws-sdk-go-v2/config v1.28.6
	github.com/aws/aws-sdk-go-v2/credentials v1.17.47
	github.com/aws/aws-sdk-go-v2/service/kms v1.37.7
	github.com/briandowns/spinner v1.23.1
	github.com/creack/pty v1.1.24
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/evergreen-ci/shrub v0.0.0-20241002130202-d2f0f3752474
	github.com/go-test/deep v1.1.1
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/golang/mock v1.6.0
	github.com/google/go-github/v61 v61.0.0
	github.com/google/uuid v1.6.0
	github.com/hinshun/vt10x v0.0.0-20220301184237-5011da428d02
	github.com/klauspost/compress v1.17.11
	github.com/mattn/go-isatty v0.0.20
	github.com/mongodb-forks/digest v1.1.0
	github.com/mongodb-labs/cobra2snooty v0.18.2
	github.com/mongodb/mongodb-atlas-kubernetes/v2 v2.5.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/shirou/gopsutil/v4 v4.24.11
	github.com/spf13/afero v1.11.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.10.0
	github.com/tangzero/inflector v1.0.0
	go.mongodb.org/atlas v0.37.0
	go.mongodb.org/atlas-sdk/v20240530005 v20240530005.0.0
	go.mongodb.org/atlas-sdk/v20241113001 v20241113001.0.0
	go.mongodb.org/mongo-driver v1.17.1
	go.opentelemetry.io/otel v1.32.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0
	go.opentelemetry.io/otel/sdk v1.29.0
	go.opentelemetry.io/otel/trace v1.32.0
	golang.org/x/exp v0.0.0-20241004190924-225e2abe05e6
	golang.org/x/mod v0.22.0
	golang.org/x/sys v0.27.0
	golang.org/x/tools v0.27.0
	google.golang.org/api v0.209.0
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.2
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.31.3
	k8s.io/apiextensions-apiserver v0.31.3
	k8s.io/apimachinery v0.31.3
	k8s.io/apiserver v0.31.3
	k8s.io/client-go v0.31.3
	sigs.k8s.io/controller-runtime v0.19.3
	sigs.k8s.io/kind v0.25.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.10.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.5 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	cloud.google.com/go/iam v1.2.2 // indirect
	cloud.google.com/go/longrunning v0.6.2 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.1.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.3.1 // indirect
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/PaesslerAG/gval v1.0.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/alessio/shellescape v1.4.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.2 // indirect
	github.com/aws/smithy-go v1.22.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github/v52 v52.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/safetext v0.0.0-20220905092116-b49f7bc46da2 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.mongodb.org/atlas-sdk/v20231115008 v20231115008.5.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.29.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sync v0.9.0 // indirect
	golang.org/x/term v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	golang.org/x/time v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20241113202542-65e8d215514f // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241113202542-65e8d215514f // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
