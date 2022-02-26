FROM golang:1.17

RUN go get github.com/cespare/reflex
RUN go get github.com/deepmap/oapi-codegen/pkg/runtime
COPY reflex.conf /

COPY /start.sh /

ENTRYPOINT [ "reflex", "-c", "/reflex.conf" ]