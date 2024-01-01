if [ "$#" -lt 1 ]; then
  echo "argument required"
  exit 1
fi

case "$1" in
  "dev-up"):
    docker compose -f dev-compose.yml up
    ;;
  "dev-down"):
    docker compose -f dev-compose.yml down
    ;;
  *)
    echo "invalid command"
    exit 1
    ;;
esac

exit 0