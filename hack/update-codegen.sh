#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

source "${CODEGEN_PKG}/kube_codegen.sh"

# generate the code with:
# --output-base    The base directory for output content, required.
#                  it must be set the parent directory of the project root directory

kube::codegen::gen_helpers \
    --input-pkg-root custom-controller/pkg/apis \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root custom-controller/pkg/apis \
    --output-pkg-root custom-controller/pkg \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
