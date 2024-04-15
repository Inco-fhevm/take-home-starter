# Inco Take-Home Starter Template

## Run a local node

1. Build the node

```bash
make build
```

The binary will be in `./build/incod`.

2. Make sure the `incod` binary is in your $PATH.

3. Run the node

```bash
./scripts/test_node.sh
```

You should see the node's logs running in your terminal. If you stop the node, and wish to re-run from the same height (not create a new chain), you can pass in `CLEAN=false ./scripts/test_node.sh`.

## Build Protobuf files

If you modify the `*.proto` files, you need to regenerate the `*.pb.go` files. Make sure you have Docker installed, and you can use the following command:

```bash
make proto-gen
```
