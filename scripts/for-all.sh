#!/bin/sh

for package in $(grep "/" "go.work" | sed -e "s#.\+/##"); do
    cd "$package"
    echo ">> $package"
    $@
    cd ..
done
