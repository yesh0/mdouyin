#!/bin/bash

dir="$1"

# CREATE USER 'mdouyin'@'localhost' IDENTIFIED BY 'test';
# CREATE DATABASE mdouyin;
# GRANT ALL PRIVILEGES ON mdouyin.* TO 'mdouyin'@'localhost';
# FLUSH PRIVILEGES;

if [[ -d "$dir" ]]; then
    cd "$dir"
    ENV_MDOUYIN_ETCD=127.0.0.1:2379                     \
        ENV_MDOUYIN_SECRET=123456789abcde               \
        ENV_MDOUYIN_CASSANDRA=127.0.0.1                 \
        ENV_MDOUYIN_BASE=http://127.0.0.1:8000          \
        ENV_MDOUYIN_REDIS=127.0.0.1:6379                \
        ENV_MDOUYIN_RDBMS="mdouyin:test@tcp(127.0.0.1:3306)/mdouyin?charset=utf8mb4&parseTime=True&loc=Local"  \
        make run -B
else
    echo "Usage: ./run.sh <dir>"
    exit 1
fi
