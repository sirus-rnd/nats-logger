# Nats Logger

![](https://github.com/sirus-rnd/nats-logger/workflows/prod/badge.svg)

service to log nats event to database

## Development

to develop this project you should have these tools

- `go` with go module enabled
- `make` utility
- `gcc` to compile some deps. that required `CGO` functionality

compile binary using this command

```bash
make all
```

to just run this project, make sure you install deps before run

```bash
make init # (optional) for first build only
make run
```

## Indexing

to make index on payload data

```sql
CREATE INDEX eventpayload ON event_models USING GIN (payload jsonb_path_ops);
```
