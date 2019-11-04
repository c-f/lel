# Automation and Deployment

Currently LEL was programmed for documentation - but documentation needs to be shared. Unfortuantly LEL has some lack of features (which might be implemented), but meanwhile you might take a look at seafile and we're providing you a ansible playbook for easy deployment.

In the end it doesn't matter what you use and how you sync/share/edit your project.

## Why Seafile

LEL is currently missing some manditory features, which \*_[Seafile](https://github.com/haiwen/seafile)_ helps to solve:

- share between members
- separation of projects
- end-to-end encryption
- open source
- good working synchronization client (stable !)
- versioning (thank god !) / history

## Deployment

Just spin up a small linux instance by your favorite Cloud-provider.
Afterwards you can use Ansible to install and configure seafile.

Ansible will install all dependencies, then build the seafile image and start seafile as a persistent docker container. Afterwards the users and the e2e encrypted docmentation is created.

**NOTE**:
Please be aware that the password is sent on the server unencrypted `once!` to the container on the same machine. However if you do not trust your cloud provider, then setup the container locally and just upload the initialized container.

1. Install and configure the host to be present in the `inv.yml` file
   IF you are not a ansible user please follow this guide: [TODO](TODO)

2. Modify the `misc_docu_srv._usernames.yml` for your teammember

3. Start the playbook. All results and passwords, will be saved inside the `seafile.release.txt` file :)

```bash
lel@srv:~/lel/_ansible-playbook$ vim misc_docu_srv._usernames.yml
lel@srv:~/lel/_ansible-playbook$ ansible-playbook sync-server.yml
```

4. Connect to seafile via a SSH-forward or change the docker settings

```bash
ssh -L 127.0.0.1:13337:127.0.0.1:13337 <serverinfos>
```

5. Install the seafile client or even use docker :D - files can be found in the `_ansible-playbook/roles/seafile_server/files` folder.

### Install and Use this ansible playbook (Linux)

```bash
TODO
```
