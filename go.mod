module github.com/grafana/grafana

go 1.13

require (
	cloud.google.com/go/firestore v1.1.1 // indirect
	cloud.google.com/go/pubsub v1.2.0 // indirect
	firebase.google.com/go v3.12.0+incompatible
	github.com/BurntSushi/toml v0.3.1
	github.com/VividCortex/mysqlerr v0.0.0-20170204212430-6c6b55f8796f
	github.com/aws/aws-sdk-go v1.25.48
	github.com/beevik/etree v1.1.0 // indirect
	github.com/benbjohnson/clock v0.0.0-20161215174838-7dc76406b6d3
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/crewjam/saml v0.0.0-20191031171751-c42136edf9b1
	github.com/davecgh/go-spew v1.1.1
	github.com/denisenkom/go-mssqldb v0.0.0-20190315220205-a8ed825ac853
	github.com/envoyproxy/go-control-plane v0.9.4 // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/fatih/color v1.7.0
	github.com/go-macaron/binding v0.0.0-20190806013118-0b4f37bab25b
	github.com/go-macaron/gzip v0.0.0-20160222043647-cad1c6580a07
	github.com/go-macaron/session v0.0.0-20190805070824-1a3cdc6f5659
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-stack/stack v1.8.0
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.1
	github.com/gobwas/glob v0.2.3
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/google/pprof v0.0.0-20200229191704-1ebb73c60ed3 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/gosimple/slug v1.4.2
	github.com/grafana/grafana-plugin-model v0.0.0-20190930120109-1fc953a61fb4
	github.com/grafana/grafana-plugin-sdk-go v0.30.0
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/go-version v1.1.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/log15 v0.0.0-20180818164646-67afb5ed74ec
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jung-kurt/gofpdf v1.10.1
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/klauspost/compress v1.4.1 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/lib/pq v1.2.0
	github.com/linkedin/goavro/v2 v2.9.7
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.3.0
	github.com/prometheus/client_model v0.1.0
	github.com/prometheus/common v0.7.0
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
	github.com/robfig/cron/v3 v3.0.0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/stretchr/testify v1.4.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/ua-parser/uap-go v0.0.0-20190826212731-daf92ba38329
	github.com/uber/jaeger-client-go v2.20.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.1.1
	github.com/yudai/gojsondiff v1.0.0
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.uber.org/atomic v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys v0.0.0-20200317113312-5766fd39f98d // indirect
	golang.org/x/tools v0.0.0-20200318150045-ba25ddc85566 // indirect
	golang.org/x/tools/gopls v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/api v0.20.0
	google.golang.org/genproto v0.0.0-20200318110522-7735f76e9fa5 // indirect
	google.golang.org/grpc v1.27.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/ini.v1 v1.46.0
	gopkg.in/ldap.v3 v3.0.2
	gopkg.in/macaron.v1 v1.3.4
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/redis.v5 v5.2.9
	gopkg.in/square/go-jose.v2 v2.4.1
	gopkg.in/yaml.v2 v2.2.5
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
