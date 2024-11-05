#!/bin/bash

URL="http://localhost:10000"

# Usage: ./send_post.sh <PAYLOAD_SIZE_IN_BYTES>
# Example: ./send_post.sh 1048576

if [ $# -ne 1 ]; then
    echo "Usage: $0 <PAYLOAD_SIZE_IN_BYTES>"
    exit 1
fi

PAYLOAD_SIZE=$1

TEMP_FILE=$(mktemp)
head -c "$PAYLOAD_SIZE" </dev/urandom > "$TEMP_FILE"
response=$(curl -s -X POST "$URL" -H "Content-Type: application/octet-stream" --data-binary "@$TEMP_FILE")

echo "Response: $response"

rm "$TEMP_FILE"
