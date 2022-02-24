module github.com/jeanmolossi/literate-robot/src/common

go 1.17

require (
	cloud.google.com/go v0.34.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.0
	github.com/go-chi/render v1.0.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.8.1
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.44.0
