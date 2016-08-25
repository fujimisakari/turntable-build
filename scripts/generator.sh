#! /bin/bash

for file in `ls $GOPATH/src/github.com/fujimisakari/turntable-build/yaml`; do
    target=`basename ${file} .yaml`
    go run $GOPATH/src/github.com/fujimisakari/turntable-build/scripts/template_generator.go ${target}
    gofmt -w $GOPATH/src/github.com/fujimisakari/turntable-build/domain/${target}/service_master.go
    gofmt -w $GOPATH/src/github.com/fujimisakari/turntable-build/model/${target}_master.go
done
