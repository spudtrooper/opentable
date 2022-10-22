# opentable

Go client API for http://opentable.com

## Usage

### Deployed frontend

https://api-dumpster-fire.herokuapp.com/opentable

### Running local front end

```
./scripts/frontend_local.sh
```

### Running the CLI

Install the CLI:

```bash

go install github.com/spudtrooper/opentable
```

or run from local with:

```bash

./scripts/run.sh
```

### Using the API

For examples using the API, see the CLI in [cli/main.go](https://github.com/spudtrooper/opentable/blob/main/cli/main.go)

### Ingesting

If you want to ingest restaurants by URI, run:

```bash

./scripts/search.sh <uri-1> <uri-2> ... <uri-N>
```

This will require a local mongo DB, controlled by `--mongo_db_name` and `--mongo_db_port`.