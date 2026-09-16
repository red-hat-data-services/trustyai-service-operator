#!/usr/bin/env bash

# Generate an RPM lockfile with rpm-lockfile-prototype.
#
# The generator runs in a container so that the host does not need the DNF
# Python bindings or rpm-lockfile-prototype installed. Run this script from
# the repository root.
#
# Usage:
#   hack/generate-rpm-lockfile.sh
#   hack/generate-rpm-lockfile.sh --rpm-input path/to/rpms.in.yaml
#
# Environment:
#   RPM_LOCKFILE_BASE_IMAGE  Container image used for generation.
#                           Defaults to registry.access.redhat.com/ubi9:9.6.
#   RPM_LOCKFILE_VERSION     rpm-lockfile-prototype version. Defaults to 0.20.0.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

RPM_INPUT="trustyai-operator-module/.konflux/rpms.in.yaml"
BASE_IMAGE="${RPM_LOCKFILE_BASE_IMAGE:-registry.access.redhat.com/ubi9:9.6}"
RPM_LOCKFILE_VERSION="${RPM_LOCKFILE_VERSION:-0.20.0}"
IMAGE="localhost/trustyai-rpm-lockfile:${RPM_LOCKFILE_VERSION}"

usage() {
    cat <<EOF
Usage: $0 [--rpm-input PATH]

Generate rpms.lock.yaml next to the selected rpms.in.yaml.

Options:
  --rpm-input PATH  Path to rpms.in.yaml, relative to the repository root.
  -h, --help        Show this help.

Environment:
  RPM_LOCKFILE_BASE_IMAGE  Base image for the generator (default: $BASE_IMAGE)
  RPM_LOCKFILE_VERSION     rpm-lockfile-prototype version (default: $RPM_LOCKFILE_VERSION)
EOF
}

while [[ $# -gt 0 ]]; do
    case "$1" in
        --rpm-input)
            [[ $# -ge 2 ]] || { echo "error: --rpm-input requires a path" >&2; exit 2; }
            RPM_INPUT="$2"
            shift 2
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            echo "error: unknown argument: $1" >&2
            usage >&2
            exit 2
            ;;
    esac
done

cd "$PROJECT_ROOT"
[[ -f "$RPM_INPUT" ]] || { echo "error: file not found: $RPM_INPUT" >&2; exit 1; }

PREFETCH_DIR="$(dirname "$RPM_INPUT")"
CONTAINER_PREFETCH_DIR="/workspace/$PREFETCH_DIR"

if ! command -v podman >/dev/null 2>&1; then
    echo "error: podman is required" >&2
    exit 1
fi

if ! podman image exists "$IMAGE" 2>/dev/null; then
    echo "Building RPM lockfile generator image: $IMAGE"
    podman build --pull=missing -t "$IMAGE" -f - . <<EOF
FROM $BASE_IMAGE
USER root
RUN dnf install -y python3 python3-pip python3-dnf rpm git && dnf clean all
RUN python3 -m pip install --no-cache-dir \\
    --index-url https://pypi.org/simple/ \\
    https://github.com/konflux-ci/rpm-lockfile-prototype/archive/refs/tags/v${RPM_LOCKFILE_VERSION}.zip
EOF
fi

echo "Generating lockfile from $RPM_INPUT"
podman run --rm \
    -v "$PROJECT_ROOT:/workspace" \
    -w /workspace \
    "$IMAGE" \
    bash -ceu '
        cd "$1"
        rpm-lockfile-prototype rpms.in.yaml
    ' bash "$PREFETCH_DIR"

echo "Generated $PREFETCH_DIR/rpms.lock.yaml"
