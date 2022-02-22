#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

~/go/bin/oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/openapi/$service.yml"
~/go/bin/oapi-codegen -generate chi-server -o "$output_dir/openapi_api.gen.go" -package "$package" "api/openapi/$service.yml"
~/go/bin/oapi-codegen -generate types -o "src/common/client/$service/openapi_types.gen.go" -package "$service" "api/openapi/$service.yml"
~/go/bin/oapi-codegen -generate client -o "src/common/client/$service/openapi_client_gen.go" -package "$service" "api/openapi/$service.yml"