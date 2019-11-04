# Documentation Documentation :D
> This file contains all information in order to setup and configure the client properly.

## Prequisition
docker, docker-compose, socat

```bash
sudo apt install docker.io socat 
# can also be docker-ce (based on the official installation)
sudo pip install docker-compose
```

## Setup 
> (simple VM has internet access, direct SSH connection possible)

0. Create a new Directory for the docu 
```bash
lel@srv:~$ mkdir -p ~/_current/docu
```
1. Download the following files into your documentation VM (except for the docker-compose.yml)
```bash
README.md  client.Dockerfile  docker-compose.client.yml  setup.sh
```
2. Create the environment (./setup.sh)
```bash
lel@srv:~/_current/docu$ ./setup.sh
```
3. Create SSH-Connection 
Forward the remote port 127.0.0.1:13331 to the docker network of the container
```bash
lel@srv:~/_current/docu$ export DOCU_NET=$(docker network inspect client-test_default -f "{{range  .IPAM.Config}}{{.Gateway}}{{end}}")
lel@srv:~/_current/docu$ ssh -L 127.0.0.1:13331:$(DOCU_NET):13331 <docu_srv>

# or in the SSH-config (vim ~/.ssh/config)
LocalForward <docker-network>:13331 127.0.0.1:13331
```
------------------------------------

## Setup Restricted
> (VM has no dns/internet - only socks)
0. Use the same step until step 2.
1. Configure Docker to use the socks proxy
```bash
root@srv:/home/lel/_current/docu# mkdir -p /etc/systemd/system/docker.service.d
root@srv:/home/lel/_current/docu# echo -e '[Service]\nEnvironment="ALL_PROXY=socks5://<>:1080"' > /etc/systemd/system/docker.service.d/socks-proxy.conf
root@srv:/home/lel/_current/docu# systemctl daemon-reload && systemctl restart docker
```
2. Configure docker-compose to use socks proxy.
> Add additional args to the docker-compose.client.yml file/ replace <socks>:1080 with your addrs.
> Please add it to both client  
```
client:
 build:
  context: .
  dockerfile: client.Dockerfile
  args:
    - USER_ID=${UID}
    - http_proxy=socks5h://<socks>:1080
    - https_proxy=socks5h://<socks>:1080
    - ALL_PROXY=socks5h://<socks>:1080

```
3. Create the environment (./setup.sh) (same)
```bash
lel@srv:~/_current/docu$ ./setup.sh
```
4. Use Socat to tunnel from the other SSH-IP to the documentation container
```bash
lel@srv:~/_current/docu$ export DOCU_NET=$(docker network inspect client-test_default -f "{{range  .IPAM.Config}}{{.Gateway}}{{end}}")
lel@srv:~/_current/docu$ socat TCP-LISTEN:13331,bind=${DOCU_NET},fork TCP:<ssh-forward-ip>:13331
```
---------------------------------------

## Usage Seafile
1. Configure seafile:
```bash
/seafile-client is the location

Server is the docu_net, which was previously exported into ${DOCU_NET}
Protocol is http, Port is 13331

e.g. http://172.22.0.1:13331

User and Password are generated and will be provided.
```
2. After login you'll see the library "docu", which is a end-to-end lib
```bash
Right-click and use "Sync this Library", specify /data

E2E-Password was allready generated and will be provided.
```
3. Finished
> The documentation folder is now synced and can also be accessed in the normal folder 
```bash
`pwd`/seafile-data
```

