echo "If exit code is 0, you're good to code"

# Path to pre-commit hook.
FILE=.git/hooks/pre-commit

if [ "$GO111MODULE" != "on" ]
then
  echo "Please, enable GO111MODULE first"
  exit 1
fi

echo "Downloading dependencies"
go mod download

echo "Vendoring"
go mod vendor

echo "Running checker script"
./scripts/checker.sh

if test -f "$FILE"; then
  echo "Pre Commit hook is already set"
  exit 0
fi

echo "Setting pre-commit git hook"
cp ./scripts/hooks/pre-commit.sh $FILE
chmod +x $FILE
