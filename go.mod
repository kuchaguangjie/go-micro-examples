module github.com/micro/examples

go 1.13

replace k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190708180123-608cd7da68f7

replace k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190708175518-244289f83105

replace google.golang.org/grpc => google.golang.org/grpc v1.24.0

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/astaxie/beego v1.12.0
	github.com/eknkc/basex v1.0.0 // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/hako/branca v0.0.0-20180808000428-10b799466ada // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/grpc/v2 v2.3.0
	github.com/micro/go-plugins/client/selector/static/v2 v2.3.0
	github.com/micro/go-plugins/config/source/configmap/v2 v2.3.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.3.0
	github.com/micro/go-plugins/registry/etcd/v2 v2.3.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.3.0
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.3.0
	github.com/micro/go-plugins/wrapper/select/shard/v2 v2.3.0
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077
	github.com/micro/micro/v3 v3.0.5-0.20201219085254-c8ea24387d19 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/pquerna/otp v1.2.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/spf13/viper v1.6.3 // indirect
	github.com/vektah/gqlparser v1.2.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.27.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/olivere/elastic.v5 v5.0.83 // indirect
)
