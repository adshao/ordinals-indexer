# Ordinals Indexer

Ordinals Indexer is an API server and index synchronizer for ordinals inscriptions. It provides a robust and efficient way to interact with and manage ordinals inscriptions data.

## Features

- Implements the [BRC-721](https://github.com/adshao/brc-721) protocol.
- Easily parse inscription content for other protocols, including ordinals domains and BRC NFTs.
- Provides a robust and efficient way to interact with and manage ordinals inscriptions data.

## Requirements

- It's recommended to run your own ordinals server before using this indexer.
- Create a database for the indexer to store data.

## Getting Started

Here are the steps to setup and run Ordinals Indexer:

### Configuration

Copy the `configs/config_example.yaml` to `configs/config.yaml` and modify the configuration file as needed.

### Database

Use `atlas` to apply the database schema to your database.

```bash
cd internal/data
atlas migrate apply --dir 'file://ent/migrate/migrations' --url 'postgres://test:test@127.0.0.1/test'
```

### Build binary

Build the binaries for the API server and syncer:

```bash
make build
```

### API Server

Run API server to start the indexer:

```bash
./bin/server -conf configs/config.yaml
```

### Syncer

Run syncer to start syncing data with the ordinals server:

```bash
./bin/sync -conf configs/config.yaml
```

## Documentation

You can find the complete API documentation [here](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/adshao/ordinals-indexer/main/openapi.yaml#/).

## Contributing

We welcome contributions from the community! Please check out our Contributing Guide for more details.

## Support

If you encounter any issues or have questions, please open an issue on this GitHub.

## License

Ordinals Indexer is licensed under the MIT License.

---
Let me know if you'd like me to fill in more specific details, or if there's any additional information you'd like me to include.
