#!/bin/bash

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    CGO_ENABLED=1 go build -a -x -v --tags "linux sqlite_omit_load_extension sqlite_vtable sqlite_fts5 sqlite_icu sqlite_json" && ldd simple
elif [[ "$OSTYPE" == "darwin"* ]]; then
    CGO_ENABLED=1 go build -a -x -v --tags "darwin sqlite_omit_load_extension sqlite_vtable sqlite_fts5 sqlite_icu sqlite_json" -ldflags "-s -w"
fi
