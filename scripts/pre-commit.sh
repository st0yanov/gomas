#!/bin/sh

GOFMT_FILES=$(gofmt -l .)
printf >&2 'Running "gofmt -l" ...\n'
if [ -n "${GOFMT_FILES}" ]; then
  printf >&2 'gofmt failed for the following files:\n%s\n\nplease run "gofmt -w ." on your changes before committing.\n' "${GOFMT_FILES}\n"
  exit 1
else
  printf >&2 'No formatting errors were found.\n'
fi
echo "------------------------------------------\n"

GOLINT_ERRORS=$(golint ./... | grep -v "Id should be")
printf >&2 'Running "golint ./..." ...\n'
if [ -n "${GOLINT_ERRORS}" ]; then
  printf >&2 'golint failed for the following reasons:\n%s\n\nplease run 'golint ./...' on your changes before committing.\n' "${GOLINT_ERRORS}\n"
  exit 1
else
  printf >&2 'No style mistakes were found.\n'
fi
echo "------------------------------------------\n"

GOVET_ERRORS=$(go tool vet ./ 2>&1)
printf >&2 'Running "go tool vet ./" ...\n'
if [ -n "${GOVET_ERRORS}" ]; then
  printf >&2 'go vet failed for the following reasons:\n%s\n\nplease run "go tool vet ./" on your changes before committing.\n' "${GOVET_ERRORS}\n"
  exit 1
else
  printf >&2 'No suspicious constructs were found.\n'
fi
echo "------------------------------------------\n"

printf >&2 'Running "go test -v -race ./..." ...\n\n'
ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --compilers=4
if [ $? -eq 0 ]; then
  printf >&2 '\nAll tests passed successfully.\n'
else
  printf >&2 '\ngo test failed, please fix before committing.\n'
  exit 1
fi
echo "------------------------------------------\n"
