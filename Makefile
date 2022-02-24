
.PHONY: openapi
openapi: openapi_http openapi_js

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh job src/job/ports ports

.PHONY: openapi_js
openapi_js:
	@./scripts/openapi-js.sh job

.PHONY: proto
proto:
	@./scripts/gen-proto.sh job