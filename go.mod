module github.com/sensu/sensu-go

go 1.13

replace (
	github.com/sensu/sensu-go/api/core/v2 => ./api/core/v2
	github.com/sensu/sensu-go/api/core/v3 => ./api/core/v3
	github.com/sensu/sensu-go/backend/store/v2 => ./backend/store/v2
	github.com/sensu/sensu-go/types => ./types
)

require (
	github.com/AlecAivazis/survey/v2 v2.2.14
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/StackExchange/wmi v0.0.0-20180725035823-b12b22c5341f // indirect
	github.com/ash2k/stager v0.0.0-20170622123058-6e9c7b0eacd4 // indirect
	github.com/atlassian/gostatsd v0.0.0-20180514010436-af796620006e
	github.com/dave/jennifer v0.0.0-20171207062344-d8bdbdbee4e1
	github.com/docker/docker v0.0.0-20180409082103-cbde00b44273
	github.com/dustin/go-humanize v1.0.0
	github.com/echlebek/crock v1.0.1
	github.com/echlebek/timeproxy v1.0.0
	github.com/emicklei/proto v1.1.0
	github.com/evanphx/json-patch/v5 v5.1.0
	github.com/frankban/quicktest v1.7.2 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-ole/go-ole v1.2.6-0.20210915003542-8b1f7f90f6b1 // indirect
	github.com/go-resty/resty/v2 v2.5.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/graph-gophers/dataloader v0.0.0-20180104184831-78139374585c
	github.com/graphql-go/graphql v0.7.9-0.20191125031726-2e2b648ecbe4
	github.com/gxed/GoEndian v0.0.0-20160916112711-0f5c6873267e // indirect
	github.com/gxed/eventfd v0.0.0-20160916113412-80a92cca79a8 // indirect
	github.com/hashicorp/go-version v1.2.0
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097
	github.com/ipfs/go-log v1.0.4 // indirect
	github.com/jbenet/go-reuseport v0.0.0-20180416043609-15a1cd37f050 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/libp2p/go-reuseport v0.0.0-20180416043609-15a1cd37f050 // indirect
	github.com/libp2p/go-sockaddr v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/mholt/archiver/v3 v3.3.1-0.20191129193105-44285f7ed244
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olekukonko/tablewriter v0.0.5
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.26.0
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/robfig/cron/v3 v3.0.1
	github.com/sensu/lasr v1.2.1
	github.com/sensu/sensu-go/api/core/v2 v2.6.0
	github.com/sensu/sensu-go/api/core/v3 v3.3.0
	github.com/sensu/sensu-go/types v0.3.0
	github.com/shirou/gopsutil v3.21.9-0.20210911021155-ce5729cbcdf6+incompatible
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/willf/pad v0.0.0-20160331131008-b3d780601022
	go.etcd.io/bbolt v1.3.6
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.etcd.io/etcd/server/v3 v3.5.0
	go.etcd.io/etcd/tests/v3 v3.5.0
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/grpc v1.38.0
	gopkg.in/h2non/filetype.v1 v1.0.3
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible // indirect
)
