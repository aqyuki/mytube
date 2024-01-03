echo "setting up"
mkdir -p bin

echo "building migrations..."
go build -o ./bin/migration cmd/migrate/migrate.go

echo "running migrations..."
./bin/migration up --path migrations

if [ "$?" -ne 0 ]; then
  echo "migrations failed"
  exit 1
fi

echo "building server..."
go build -o bin/server cmd/server/server.go
./bin/server

case "$?" in
  0):
    echo "server shut donwned successfully"
    eixt 0
    ;;
  *):
    echo "server shut down with error"
    exit 1
    ;;
esac
