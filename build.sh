#!/bin/bash
cd static;
npm run build;
cd ..
goxc;
docker build . -t $(jq '"invist/lel:" + .version' static/package.json)
docker-compose -f docker-compose.yml.build up --build