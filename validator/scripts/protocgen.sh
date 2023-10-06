#!/bin/sh

set -e

if ! grep -q "github.com/choraio/mods/validator" go.mod ; then
  echo -e "ERROR: This command must be run from inside the validator module."
  return 1
fi

echo "Updating dependencies"

cd proto
buf mod update
cd ..

echo "Creating tmp directory"

mkdir -p proto-tmp/chora/validator
cd proto
find . -maxdepth 1 -mindepth 1 -type f -exec cp '{}' ../proto-tmp/ \;
find . -maxdepth 1 -mindepth 1 -type d -exec cp -r '{}' ../proto-tmp/chora/validator/ \;
cd ..

echo "Generating gogo files"

cd proto-tmp

proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package "$file" &> /dev/null ; then
      buf generate --template buf.gen.gogo.yaml "$file"
    fi
  done
done

cd ..

cp -r github.com/choraio/mods/validator/* ./
rm -rf github.com

echo "Generating pulsar files"

go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm@latest

cd api

find ./ -type f \( -iname \*.pulsar.go -o -iname \*.pb.go -o -iname \*.cosmos_orm.go -o -iname \*.pb.gw.go \) -delete
find . -empty -type d -delete

cd ../proto-tmp

buf generate --template buf.gen.pulsar.yaml

cd ../api/chora/validator
find . -maxdepth 1 -mindepth 1 -type d -exec cp -r '{}' ../../ \;
cd ../..
rm -rf chora

cd ..

echo "Generating swagger files"

cd proto-tmp

proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' \))
  if [[ ! -z "$file" ]]; then
    buf generate --template buf.gen.swagger.yaml "$file"
  fi
done

cd ../docs/chora/validator
find . -maxdepth 1 -mindepth 1 -type d -exec cp -r '{}' ../../ \;
cd ../..
rm -rf chora

npm list -g | grep swagger-combine > /dev/null || npm install -g swagger-combine --no-shrinkwrap

swagger-combine config.json -f yaml \
  -o swagger.yaml \
  --continueOnConflictingPaths true \
  --includeDefinitions true

cd ..

echo "Removing tmp directory"

rm -rf proto-tmp
