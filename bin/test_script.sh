#!/bin/bash

PROGNAME="$(basename $0)"

error_exit()
{
  echo "${PROGNAME}: ${1:-"Unknown Error"}" 1>&2
  exit 1
}

echo "Hello World"
echo "Example of error with line number and message"
error_exit "$LINENO: An error has occurred."
