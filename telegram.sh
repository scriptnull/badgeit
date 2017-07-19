#!/bin/bash
set -e
echo "Posting message to Telegram"
echo $1
echo "---------------------------"
curl -H "Content-Type: application/json" -X POST -d "{ \"text\": \"$1\" }" $TELEGRAM_HORN
