#!/usr/bin/env bash

EXIT_CODE=0

echo "" > coverage.txt

for d in $(go list ./...); do
    go test -coverpkg=./... -coverprofile=profile.out $d

    if [ $? -eq 1 ]; then
      EXIT_CODE=1
    fi

    if [ -f profile.out ] && [ $EXIT_CODE != 1 ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

if [ $EXIT_CODE == 1 ]; then
  exit $EXIT_CODE
fi


if [ -n "$CODECOV_TOKEN" ]; then
  curl -s https://codecov.io/bash | bash
  rm coverage.txt
fi
