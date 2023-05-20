IFS=',' read -ra HOSTS <<< "$1" # you could replace $1 with a hard-coded list if you prefer
for i in "${HOSTS[@]}"; do
  output=$(ssh root@$i -C "nodetool repair -pr") # replace with your command (ssh, kubectl, fly ssh, etc.)
  exit_code=$?
  if [ $exit_code -ne 0 ]; then
      echo "😥 SSH command failed with exit code $exit_code for host '$i'"
      exit 1
  fi
  echo "✅ All repairs complete!"
done
