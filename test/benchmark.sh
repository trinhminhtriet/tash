#!/bin/bash

set -euo pipefail

function parse_options() {
  SAVE=
  while [[ $# -gt 0 ]]; do
    case "${1}" in
      --save|-s)
        SAVE=YES
        shift
        ;;
      *)
        printf "Unknown flag: ${1}\n\n"
        help
        exit 1
        ;;
    esac
  done
}

function __main__() {
  export TASH_USER_CONFIG="$PWD/user-config.yaml"
  parse_options $@
    if [[ "$SAVE" ]]; then
        hyperfine -N --runs 10 '../dist/tash run ping -s server-9' > ./profiles/ping-no-key
        hyperfine -N --runs 10 '../dist/tash run ping --forks=1 -t reachable' > ./profiles/ping
        hyperfine -N --runs 10 '../dist/tash run ping --strategy=free -t reachable' > ./profiles/ping-parallel
        hyperfine -N --runs 10 '../dist/tash run d --forks=1 -t reachable' > ./profiles/nested
        hyperfine -N --runs 10 '../dist/tash run d --strategy=free -t reachable' > ./profiles/nested-parallel
        hyperfine -N --runs 10 '../dist/tash list servers' > ./profiles/list-servers
    else
        hyperfine -N --runs 10 '../dist/tash run ping -s server-9'
        hyperfine -N --runs 10 '../dist/tash run ping --forks=1 -t reachable'
        hyperfine -N --runs 10 '../dist/tash run ping --strategy=free -t reachable'
        hyperfine -N --runs 10 '../dist/tash run d --forks=1 -t reachable'
        hyperfine -N --runs 10 '../dist/tash run d --strategy=free -t reachable'
        hyperfine -N --runs 10 '../dist/tash list servers'
    fi
}

__main__ $@
