FILE=.git/hooks/pre-commit

if test -f "$FILE"; then
  echo "Pre Commit hook is already set"
  exit 0
fi

cp ./scripts/hooks/pre-commit.sh $FILE
chmod +x $FILE

echo "Hook is set"

echo "Running checker script"

./scripts/checker.sh

echo "If exit code is 0, you're good to code"
