#!/usr/bin/env bash

SoName="libopenimsdk.so"
OutputPath="../shared"
GoSrc="."
go build -buildmode=c-shared -trimpath -ldflags="-s -w" -o ${OutputPath}/${SoName} ${GoSrc}/export.go ${GoSrc}/listener.go ${GoSrc}/tools.go