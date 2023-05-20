IFS=',' read -ra HOSTS <<< "$1"
for i in "${HOSTS[@]}"; do
  echo "🛠️  Repairing host '$i'..."
  output=$(docker exec $i nodetool repair -pr)
  exit_code=$?
  if [ $exit_code -ne 0 ]; then
      echo "😥 SSH command failed with exit code $exit_code for host '$i'"
      exit 1
  fi
  echo "☑️ Repaired host '$i'"
done
echo "✅ All repairs complete!"
