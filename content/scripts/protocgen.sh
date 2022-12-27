#!/bin/sh

set -e

if ! grep "github.com/choraio/mods/content" go.mod &>/dev/null ; then
  echo -e "ERROR: This command must be run from inside the content module."
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
cp -r github.com/choraio/mods/content/* ./
rm -rf github.com

echo "Generating pulsar files"

go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/cosmos/cosmos-sdk/orm/cmd/protoc-gen-go-cosmos-orm@latest

cd api
find ./ -type f \( -iname \*.pulsar.go -o -iname \*.pb.go -o -iname \*.cosmos_orm.go -o -iname \*.pb.gw.go \) -delete
find . -empty -type d -delete

cd ../proto
buf generate --template buf.gen.pulsar.yaml
