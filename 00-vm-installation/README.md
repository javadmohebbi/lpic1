# Intro
To do the exercises in our course, we assume you have either Ubuntu 18.04 LTS or CentOS 7 available. You can install either or both of these Linux distributions in a virtual environment using the Oracle VirtualBox software, available at [https://virtualbox.org](https://virtualbox.org).

# Download the required images
- **Ubuntu Server 18.04 LTS**: [Download ISO](https://releases.ubuntu.com/18.04/ubuntu-18.04.6-live-server-amd64.iso)
- **CentOS 7**: [Download ISO](http://mirrors.standaloneinstaller.com/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-NetInstall-2009.iso)


# Requirement
1. You need to install VirtualBox (version 6.x is recommended)
    - Also may want to use VMware workstation or other similar software. Since VirtualBox is free software, we will use it here.
2. Internet connection on host OS is required.
3. 64bit Processors

# Installation Guide
I recommend you install both OS to get familiar with different distibutions. Be sure not to config the same IP address for the OS **host-only** adapters. Because you might need to connect to both of them at the same time.
- Ubuntu Server 18.04 LTS: [Installation Guide](UBUNTU-README.md)
- CentOS 7: [Installation Guide](CENTOS-README.md)

# Connect to VMs using SSH
- Windows: [Guide](SSH-README.md#windows-os)
- Linux / macOS: [Guide](SSH-README.md#linux--macos)