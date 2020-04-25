#!/bin/bash

################## VARIABLES ####################
PORT="3000"
API="api/v1/book"
URL="http://localhost:${PORT}/${API}"


################## REQUEST ######################
curl -d '{
  "title": "ASOIAF",
  "author": "GRRM",
  "rating": 5
}' \
 -H "Content-Type: application/json" \
 -X POST ${URL}
