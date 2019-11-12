#!/bin/bash
cd static;
npm run build;
cd ..
goxc;
docker build . -t $(jq -r '"invist/lel:" + .version' static/package.json)
docker-compose -f docker-compose.yml.build up --build

find . -path "./pkg/*" -name "LeL_*" -type f -exec sha256sum {} \; > pkg/LeL/$(jq -r '.version' static/package.json)/SHA256SUMS
docker push $(jq -r '"invist/lel:" + .version' static/package.json)
