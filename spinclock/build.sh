#!/bin/sh
GOOS=darwin GOOARCH=arm64 go build -o au.id.hawkins.sd.spinclock.sdPlugin/spinclock .
GOOS=windows GOOARCH=amd64 go build -o au.id.hawkins.sd.spinclock.sdPlugin/spinclock.exe  .
