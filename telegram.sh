#!/bin/bash
set -e
curl -H "Content-Type: application/json" -X POST -d "{ \"text\": \"$1\" }" $TELEGRAM_HORN
