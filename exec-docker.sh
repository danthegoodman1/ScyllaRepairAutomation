HOST=$1
CMD=$2

echo running "$2" on host "$1"
docker exec $1 $2