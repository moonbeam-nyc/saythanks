#!/bin/bash
# This script is used to test the address validation endpoint
curl -X POST http://localhost:8080/api/address/validate \
     -H "Content-Type: application/json" \
     -d '{ "address": "3050 21st st", "city": "queens", "state": "NY" }' | jq -r .data | jq