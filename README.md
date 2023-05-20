# ScyllaRepairAutomation

Scylla Manager is cool, but it requires 2 annoying things:

1. You have to setup the agent on every node
2. You have to have an enterprise license for more than 5 nodes

1 is not a huge deal, 2 sucks.

FWIW if you are an enterprise customer definitely use Scylla Manager. This is for the rest of us open source users.

## How to run this

You can run this manually on some interval, setup a cron job, a kubernetes job per-cluster, etc. You will want to run this every 7 days (by default, change according to your settings).

## Go Version

Supports optional logging, failure and success scripts, etc.

```
go run . -exec exec.sh -hosts host1,host2 -s succ.sh -f fail.shgo run
```

You can test with the local `docker-compose.yaml` by running:

```
docker compose up -d
```

(wait for ~30 seconds)

```
bash test.sh
```

## Bash Version

This version is very simple and hard-coded. You can change the command within the contents of the file.

```
./repair.sh "10.0.0.1:9042,10.0.0.2:9042,10.0.0.3:9042"
```

For example `repair-docker.sh` exists for repairing with a local docker setup (tested using the `docker-compose.yaml`).
