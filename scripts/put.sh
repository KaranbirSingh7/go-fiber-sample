#!/bin/bash

################## VARIABLES ####################
PORT="3000"
API="api/v1/book"
BOOK_ID="3"
URL="http://localhost:${PORT}/${API}/${BOOK_ID}"


################## REQUEST ######################
curl -d '{
  "title": "WOW",
  "author": "GRRM",
  "rating": 5
}' \
 -H "Content-Type: application/json" \
 -X PUT ${URL}
