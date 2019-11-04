#!/bin/bash

# This script setup the environment

echo -e "[+] Allow docker GUIs to be displayed"
xhost +local:

echo -e "[+] Create document structure"
mkdir -p ./seafile-data/docu_notes
chown -R ${UID}:${UID} ./seafile-data

echo -e "[+] Create Docker network (docu_net"
docker network create docu_net 2> /dev/null

echo -e "[*] Information about the network"
export DATA_BIND=$(docker network inspect docu_net -f "{{range  .IPAM.Config}}{{.Gateway}}{{end}}")
echo -e "[*] Docker network: ${DATA_BIND}"

echo -e "[*] Building the container"
export UID 
docker-compose -f docker-compose.client.yml up -d client

echo -e "[*] Logoutput:" 
docker-compose -f docker-compose.client.yml logs -f client

