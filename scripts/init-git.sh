#!/bin/sh

cp ./scripts/pre-commit.sh ./.git/hooks/pre-commit
cp ./scripts/pre-push.sh ./.git/hooks/pre-push

chmod 755 ./.git/hooks/pre-commit
chmod 755 ./.git/hooks/pre-push