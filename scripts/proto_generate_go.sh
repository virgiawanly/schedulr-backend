#!/bin/sh

SCRIPT_DIR="$(dirname "$(realpath "$0")")"
WORKDIR="$SCRIPT_DIR/.."
PROTODIR="$WORKDIR/proto"
OUTDIR="$WORKDIR/protogen/go"
GATEWAYOUTDIR="$WORKDIR/protogen/gateway/go"

# Ensure output directory exists
mkdir -p "$OUTDIR"
mkdir -p "$GATEWAYOUTDIR"

# Find and compile all .proto files recursively
find "$PROTODIR" -name "*.proto" | while read -r entry; do
  echo "Processing: $entry"
  protoc --proto_path="$PROTODIR" --proto_path="$WORKDIR" \
    --go_out="$OUTDIR" --go_opt=paths=source_relative \
    --go-grpc_out="$OUTDIR" --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out "$GATEWAYOUTDIR" \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt standalone=true \
    --grpc-gateway_opt generate_unbound_methods=true \
    "$entry"
done

echo "Proto files generated successfully in $OUTDIR"
