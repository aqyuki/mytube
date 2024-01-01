if [ "$#" -lt 1]; then
  echo "argument required"
  exit 1
fi

case "$1" in
  "create"):
    if [ "$#" -lt 2 ]; then
      echo "migration name required"
      exit 1
    fi
    migrate create -ext sql -dir migrations $2
    ;;
  "up"):
    go run cmd/migrate/migrate.go --path migrations
    ;;
  "help")
    echo "create <name> - create a new migration"
    echo "up - run all migrations"
    ;;
  *)
    echo "invalid command"
    exit 1
    ;;
esac