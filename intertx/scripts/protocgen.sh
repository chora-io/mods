#!/usr/bin/env bash

set -e

if ! grep "github.com/choraio/mods/intertx" go.mod &>/dev/null ; then
  echo -e "ERROR: This command must be run from inside the intertx module."
  return 1
fi

echo "Generating gogo files"

cd proto
buf mod update
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &> /dev/null ; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done
cd ..
cp -r github.com/choraio/mods/intertx/* ./
rm -rf github.com

echo "Generating pulsar files"

cd api
find ./ -type f \( -iname \*.pulsar.go -o -iname \*.pb.go -o -iname \*.cosmos_orm.go -o -iname \*.pb.gw.go \) -delete
find . -empty -type d -delete
cd ..
cd proto
buf generate --template buf.gen.pulsar.yaml
