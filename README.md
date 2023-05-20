# ScyllaRepairAutomation

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
