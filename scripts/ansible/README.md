# Ansible for Axia

[Ansible](https://ansible.com) playbooks, roles, & inventories to install
[Axia](https://github.com/Axia-Tech/axiav2) as a systemd service.
Target(s) can be

- localhost
- Cloud VMs (e.g. Amazon, Azure, Digital Ocean)
- Raspberry Pi
- any host running a supported operating system
- any combination of the above


## Using

To create an Axia service on localhost

1. Check you have Ansible 2.9+ (see [Installing](#installing))
2. Clone the Axia git repository
    ```
    $ git clone https://github.com/Axia-Tech/axiav2
    ```

3. Change to this directory
    ```
    $ cd Axia/scripts/ansible
    ```

4. Run the service playbook
    ```
    $ ansible-playbook \
        -i inventories/localhost.yml \
        --ask-become-pass \
        service_playbook.yml
    ```

   You don't need `--ask-become-pass` if your account doesn't require a sudo
   password. To install on remote hosts you will need to create an inventory,
   see [customising].

   Output should look similar to the [example run](#example-run).

5. Check the service is running
    ```
    $ systemctl status Axia
    ```

    The output should look similar to
    ```
    ● Axia.service - Axia node for Axia consensus network
    Loaded: loaded (/etc/systemd/system/Axia.service; enabled; vendor preset: enabled)
    Active: active (running) since Wed 2020-10-21 10:00:00 UTC; 32s ago
    ...
    ```


## Installing

Ansible 2.9 (or higher) is required. To check, run

```
ansible --version
```

the first line of output should be like `ansible 2.9.x`, or `ansible 2.10.x`
(`x` can be any number). If the output includes `ansible: command not found`,
or an earlier version (e.g. `ansible 2.8.16`), then you need to install a
supported version.

To install a supported version

1. Create a Python Virtualenv
    ```
    $ python3 -m venv venv/
    ```

2. Activate the Virtualenv
    ```
    $ source venv/bin/activate
    ```

4. Install Ansible inside the virtualenv
    ```
    $ pip install "ansible>=2.9"
    ```


## Customising

To run against remote targets you'll need an [Ansible inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html#inventory-basics-formats-hosts-and-groups).
Here are some examples to use as a starting point.

### Amazon

```yaml
axia_nodes:
  hosts:
    ec2-203-0-113-42.us-east-1.compute.amazonaws.com:
    ec2-203-0-113-9.ap-southeast-1.compute.amazonaws.com:
  vars:
    ansible_ssh_private_key_file: "~/.ssh/aws_identity.pem"
    ansible_user: "ubuntu"
```

### Raspberry Pi

```yaml
axia_nodes:
  hosts:
    raspberrypi.local:
  vars:
    ansible_user: "pi"
```

## Requirements

Target operating systems supported by these roles & playbooks are

- CentOS 7
- CentOS 8
- Debian 10
- Raspberry Pi OS
- Ubuntu 18.04 LTS
- Ubuntu 20.04 LTS


## Example run

```
PLAY [Configure Axia service] *********************************************

TASK [Gathering Facts] *********************************************************
ok: [localhost]

TASK [axia_download : Query releases] *************************************
ok: [localhost]

TASK [axia_download : Fetch release] **************************************
changed: [localhost] => (item=Axia-linux-arm64-v1.5.0.tar.gz)
changed: [localhost] => (item=Axia-linux-arm64-v1.5.0.tar.gz.sig)

TASK [axia_download : Create temp gnupghome] ******************************
changed: [localhost]

TASK [axia_download : Import keys] ****************************************
changed: [localhost]

TASK [axia_download : Verify signature] ***********************************
ok: [localhost]

TASK [axia_download : Cleanup temp gnupghome] *****************************
changed: [localhost]

TASK [axia_download : Unpack release] *************************************
changed: [localhost] => (item=Axia-linux-arm64-v1.5.0.tar.gz)

TASK [axia_user : Create Axia daemon group] **************************
changed: [localhost]

TASK [axia_user : Create Axia daemon user] ***************************
changed: [localhost]

TASK [axia_install : Create shared directories] ***************************
changed: [localhost] => (item={'path': '/usr/local/bin'})
changed: [localhost] => (item={'path': '/var/local/lib'})
changed: [localhost] => (item={'path': '/var/local/log'})

TASK [axia_install : Create Axia directories] ************************
changed: [localhost] => (item=/var/local/lib/Axia)
changed: [localhost] => (item=/var/local/lib/Axia/db)
changed: [localhost] => (item=/var/local/lib/Axia/staking)
changed: [localhost] => (item=/var/local/log/Axia)
changed: [localhost] => (item=/usr/local/lib/Axia)

TASK [axia_install : Install Axia binary] ****************************
changed: [localhost]

TASK [axia_install : Remove outdated support files] **********************
ok: [localhost] => (item={'path': '/usr/local/lib/Axia/evm'})
ok: [localhost] => (item={'path': '/usr/local/lib/Axia/Axia-preupgrade'})
ok: [localhost] => (item={'path': '/usr/local/lib/Axia/Axia-latest'})

TASK [axia_install : Install support files] *******************************
changed: [localhost] => (item=/usr/local/lib/Axia/plugins)

TASK [axia_staker : Create staking key] ***********************************
changed: [localhost]

TASK [axia_staker : Create staking certificate signing request] ***********
changed: [localhost]

TASK [axia_staker : Create staking certificate] ***************************
changed: [localhost]

TASK [axia_service : Configure Axia service] *************************
changed: [localhost]

TASK [axia_service : Enable Axia service] ****************************
changed: [localhost]

RUNNING HANDLER [axia_service : Restart Axia service] ****************
changed: [localhost]

PLAY RECAP *********************************************************************
localhost : ok=6 changed=24 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
```
