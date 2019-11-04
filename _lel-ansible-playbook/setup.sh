#!/bin/bash 

# sudo apt install virtualenv curl python-pip jq libnotify-bin

# --[Ansible Virtual Environment]--
echo -e "\n\e[34m[*]\e[39m Check for Virtual ansible envrionment\e[90m"
if [ ! -d "${PWNGADGET_DIR}/ansible2.7" ]; then

	echo -e "\n\e[32m[+]\e[39m Install Ansible to: ${PWNGADGET_DIR}/ansible2.7\e[90m"
	virtualenv ansible2.7
	source "${PWNGADGET_DIR}/ansible2.7/bin/activate"
	pip install ansible==2.7
	pip install dopy

fi
