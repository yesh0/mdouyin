#!/bin/bash

dir="$1"

if [[ -d "$dir" ]]; then
    cd "$dir"
    ENV_MDOUYIN_ETCD=127.0.0.1:2379                     \
        ENV_MDOUYIN_RDBMS="file::memory:?cache=shared"  \
        ENV_MDOUYIN_SECRET=123456789abcde               \
        ENV_MDOUYIN_CASSANDRA=127.0.0.1                 \
        ENV_MDOUYIN_BASE=http://127.0.0.1:8000          \
        make run -B
else
    echo "Usage: ./run.sh <dir>"
    exit 1
fi
