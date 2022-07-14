#!/bin/bash

for dir in cmd/*/; do
  appname=$(basename "$dir")
  for filepath in "$dir"*; do
    file=$(basename "$filepath")
    if [ "$file" == "Dockerfile" ]; then
      echo "Found dockerfile at $filepath. Running docker build. Image name will be memepen/$appname."
      docker build . -f "$filepath" -t "memepen/$appname"
    fi
  done
done