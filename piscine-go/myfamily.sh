#! /bin/bash

curl -s https://01.gritlab.ax/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ( $hero_id | tonumber)) | .connections.relatives' | tr -d '"'

# curl "https://01.gritlab.ax/assets/superhero/all.json" | jq ' .[] | 
# select( .id== '$HERO_ID') | .connection.relatives' | tr -d '"'