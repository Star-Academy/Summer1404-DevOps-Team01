# Ansible Task: Install Python on a Virtual Machine

## Task Overview

This document provides step-by-step instructions on how to use Ansible to install Python on a virtual machine. This is a fundamental DevOps task that demonstrates Ansible's ability to automate software installation across remote systems.

## Prerequisites

Before beginning this task, you'll need:

1. A control node with Ansible installed
2. A target virtual machine with SSH access
3. Basic understanding of YAML syntax
4. SSH key-based authentication configured between the control node and target VM

## Step 1: Install Ansible on the Control Node

First, let's install Ansible on your control node (the machine from which you'll run the automation):

```bash
# For Debian/Ubuntu systems
sudo apt update
sudo apt install ansible -y

# For Red Hat/CentOS systems
sudo yum install epel-release -y
sudo yum install ansible -y
```

**What is happening here?**

- Ansible is a configuration management tool written in Python
- The control node is the computer where you install Ansible and run commands from
- The package manager (apt or yum) downloads and installs Ansible and its dependencies

## Step 2: Create an Inventory File

Create a file named `inventory.ini` to list your target virtual machine:

```ini
[servers]
target-vm ansible_host=185.60.136.214 ansible_user=root
```

**What is an inventory file?**

- An inventory file defines the hosts and groups of hosts upon which commands, modules, and tasks in a playbook operate
- It lists the IP addresses or hostnames of your target machines
- The `ansible_host` parameter specifies the IP address
- The `ansible_user` parameter specifies the user account to use when connecting

## Step 3: Test the Connection

Verify that Ansible can connect to your target virtual machine:

```bash
# Copy SSH key to target VM
ssh-copy-id root@185.60.136.214
# Test connectivity using the ping module
ansible all -i inventory.ini -m ping
# or if you need to enter a password
ansible all -i inventory.ini -m ping --ask-pass
```

**What is happening here?**

- The `-i` flag specifies the inventory file to use
- The `-m ping` instructs Ansible to use the "ping" module, which tests connectivity
- Successful output will show "pong" responses from your targets
- This verifies SSH connectivity and authentication without making any changes

## Step 4: Create a Playbook to Install Python

Create a YAML file named `install_python.yml` with the following content:

```yaml
---
- name: Install Python on target VM
  hosts: servers
  become: yes  # This enables privilege escalation (sudo)
  tasks:
    - name: Update package cache
      package:
        update_cache: yes
      
    - name: Install Python
      package:
        name: 
          - python3
          - python3-pip
        state: present
      
    - name: Check Python version
      command: python3 --version
      register: python_version
      changed_when: false
      
    - name: Display Python version
      debug:
        msg: "Python version: {{ python_version.stdout }}"
```

**What is this playbook doing?**

- A playbook is a YAML file that contains a list of tasks to be executed on remote hosts
- The `hosts` parameter specifies which hosts or groups from the inventory to target
- The `become: yes` directive enables privilege escalation (sudo) to install software
- The `package` module is a platform-independent way to manage packages
- The `register` keyword stores the output of a command for later use
- The `debug` module displays information during execution

## Step 5: Run the Playbook

Execute the playbook to install Python on your target VM:

```bash
ansible-playbook -i inventory.ini install_python.yml
```

**What is happening here?**

- The `ansible-playbook` command runs a playbook against your inventory
- Ansible connects to each host in the "servers" group
- It executes each task in sequence
- It reports the status of each task (changed, ok, failed)
- The final output will show a summary of changes made

## Step 6: Verify the Installation

To confirm that Python was successfully installed:

```bash
ansible servers -i inventory.ini -m command -a "python3 --version"
```

**What is this command doing?**

- It runs a direct command on all hosts in the "servers" group
- The `-m command` specifies the use of the command module
- The `-a` parameter passes arguments to the module
- The output will show the Python version installed on each target

## Best Practices

1. **Idempotence**: The playbook is designed to be idempotent, meaning it can be run multiple times without causing problems
2. **Platform Independence**: Using the `package` module allows the playbook to work on different Linux distributions
3. **Validation**: The playbook includes steps to verify that the installation succeeded
4. **Documentation**: Comments explain what each section does

## Troubleshooting

If you encounter issues:

1. **Playbook syntax errors**: Use the `--syntax-check` flag to validate your playbook

   ```bash
   ansible-playbook --syntax-check install_python.yml -i inventory.ini
   ```

4. **Verbose output**: Add `-v`, `-vv`, or `-vvv` flags for increasing levels of verbosity

   ```bash
   ansible-playbook -i inventory.ini install_python.yml -vv
   ```

## Conclusion

You have successfully used Ansible to automate the installation of Python on a virtual machine. This basic task demonstrates the core workflow of Ansible automation:

1. Define target hosts in an inventory
2. Create a playbook with desired state
3. Execute the playbook
4. Verify the results

This same pattern can be extended to more complex automation scenarios, making it a fundamental skill for DevOps engineers.
