# ScyllaRepairAutomation

Scylla Manager is cool, but it requires 2 annoying things:

1. You have to setup the agent on every node
2. You have to have an enterprise license for more than 5 nodes

1 is not a huge deal, 2 sucks.

FWIW if you are an enterprise customer definitely use it. This is for the rest of us open source users.

## Go Version

Supports optional logging, failure and success scripts, etc.

```
go run . -exec exec.sh -hosts host1,host2 -s succ.sh -f fail.shgo run
```

## Bash Version

This version is very simple and hard-coded. You can change the command within the contents of the file.

```
./repair.sh "10.0.0.1:9042,10.0.0.2:9042,10.0.0.3:9042"
```
