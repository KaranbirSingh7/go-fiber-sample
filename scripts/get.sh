#!/bin/bash

################## VARIABLES ####################
PORT="3000"
API="api/v1/book"
URL="http://localhost:${PORT}/${API}"


################## REQUEST ######################
curl -X GET ${URL}
