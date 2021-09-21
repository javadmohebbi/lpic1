# Intro
To do the exercises in our course, we assume you have either Ubuntu 18.04 LTS or CentOS 7 available. You can install either or both of these Linux distributions in a virtual environment using the Oracle VirtualBox software, available at [https://virtualbox.org](https://virtualbox.org).

# Download the required images
- **Ubuntu Server 18.04 LTS**: [Download ISO](https://releases.ubuntu.com/18.04/ubuntu-18.04.6-live-server-amd64.iso)
- **CentOS 7**: [Download ISO](http://mirrors.standaloneinstaller.com/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-NetInstall-2009.iso)


## Requirement
1. You need to install VirtualBox (version 6.x is recommended)
    - Also may want to use VMware workstation or other similar software. Since VirtualBox is free software, we will use it here.
2. Internet connection on host OS is required.

### Installing *Ubuntu 18.04* on VirtualBox
1. Download ISO image from [this link](https://releases.ubuntu.com/18.04/ubuntu-18.04.6-live-server-amd64.iso)

2. Open **VirtualBox** and go to **Machine** menu and then click on **New**. You can press **CTRL+N**. Create virtual machine window appears.
   - **Name**: lpic-ubuntu
   - **Machine Folder**: Change this if this location has not enough space otherwise leave it as it's default value
   - **Type**: Linux
   - **Version**: Ubuntu (64-bit)

![ubuntu-01-vm-wizard-1](./ubuntu/ubuntu-01-vm-wizard-1.png)

3. Click **next**. In this window you should config the virtual machine memory. 1024MB (1GB) memory is enough for our lab. But you can change it according to your need.

![ubuntu-01-vm-wizard-2](ubuntu/ubuntu-01-vm-wizard-2.png)

4. Click **next** and you will see Hard disk confoguration window. Click on **create** button and leave everything default.

![ubuntu-01-vm-wizard-3](ubuntu/ubuntu-01-vm-wizard-3.png)

5. Leave everything default in "Hard disk file type" window and click **next**.

![ubuntu-01-vm-wizard-4](ubuntu/ubuntu-01-vm-wizard-4.png)

6. In this window, select **Dynamically allocated** and click **next**.

![ubuntu-01-vm-wizard-5](ubuntu/ubuntu-01-vm-wizard-5.png)

7. You can change the hard disk size. The default value is enough. But you can increase it as you want but don't decrease it. Then click **create** button.

![ubuntu-01-vm-wizard-6](ubuntu/ubuntu-01-vm-wizard-6.png)

8. Now you will see the newly created virtual machine (**lpic-ubuntu**).

![ubuntu-02-vm-created](ubuntu/ubuntu-02-vm-created.png)

9. **Right click** on **lpic-ubuntu** machine and click **settings**. Then go to **Network** Section and select **Adapter 2** tab. Do the following changes:
    - Check the **Enable Network Adapter** check box.
    - Select **Host-only adapter** for **Attached To** drop-down.
    - Select **vboxnet0** for **Name** drop-down.

![ubuntu-03-vm-setting-1](ubuntu/ubuntu-03-vm-setting-1.png)

10. Now goto **Storage** section and below **Controller: IDE** select **Empty**. Then click on ![disk-icon](ubuntu/disk-icon.png) (**Disk icon**) and from the pop-up menu, select **Choose a disk file**.

![ubuntu-03-vm-setting-2](ubuntu/ubuntu-03-vm-setting-2.png)

11. Now do the following
    - Select your **Ubuntu ISO image** and click **open**.
    - Then click **OK** to close the virtual mahchine settings window.

![ubuntu-03-vm-setting-3](ubuntu/ubuntu-03-vm-setting-3.png)

12. Select **lpic-ubuntu** machine and click on **Start** button.

![ubuntu-04-vm-install-1](ubuntu/ubuntu-04-vm-install-1.png)

13. The virtual machine starts and boots from the provided ISO image. In the first step we should select the language. I prefer **English**. Select your language and press **Enter**.

![ubuntu-04-vm-install-2](ubuntu/ubuntu-04-vm-install-2.png)

14. Choose your Keyboard layout and variant and select **Done** to confirm your layout. (I leave all the setting default)

![ubuntu-04-vm-install-3](ubuntu/ubuntu-04-vm-install-3.png)

15. Now, it's time to configure our virtual machine network configuration. If you did all the step above, you **must** see **two** network card. In my case **enp0s3** (NAT) and **enp0s8** (Host-only).

![ubuntu-04-vm-install-4](ubuntu/ubuntu-04-vm-install-4.png)

16. Don't change the **enp0s3**. We want to configure our **Host-only** adapter which is **enp0s8**. To do so, using **Up arrow key** select **enp0s8** and press **Enter** key. The using **Down arrow key** and highlight the **Edit IPv4** and press **Enter** again.
    - **\*\*\* Note**: Before go further, write down the current IP address that our network adapter got from DHCP server (In my case this IP address is `192.168.56.103/24` ).

![ubuntu-04-vm-install-5](ubuntu/ubuntu-04-vm-install-5.png)

17. In **IPv4 Method**, press **Enter**, select **Manual** and again press **Enter**.

![ubuntu-04-vm-install-6](ubuntu/ubuntu-04-vm-install-6.png)

18. Provide the IP address and subnet mask from the **step 16** in the relevant fields. (We don't need to setup DNS and gateway for this adapter for our Lab). Then select **Save** and press **Enter** to save this configuration

![ubuntu-04-vm-install-7](ubuntu/ubuntu-04-vm-install-7.png)

19. Select **Done** and press **Enter**.

![ubuntu-04-vm-install-8](ubuntu/ubuntu-04-vm-install-8.png)

20. If you use proxy to access to the Internet, provide your proxy information here. Otherwise leave **Proxy address** empty and select **Done** and press **Enter**.

![ubuntu-04-vm-install-9](ubuntu/ubuntu-04-vm-install-9.png)

21.  Installation will fill the best **Mirror address**. You can change it but I recommend you leave it empty for now and select **Done** and press **Enter**.

![ubuntu-04-vm-install-10](ubuntu/ubuntu-04-vm-install-10.png)

22. In our Lab, we don't want to change the default configuration. So, select **Done** and press **Enter**.

![ubuntu-04-vm-install-11](ubuntu/ubuntu-04-vm-install-11.png)

23. Again, select **Done** and press **Enter**.

![ubuntu-04-vm-install-12](ubuntu/ubuntu-04-vm-install-12.png)

24. Installation needs you to confirm the hard disk configuration. select **Continue** and press **Enter**

![ubuntu-04-vm-install-13](ubuntu/ubuntu-04-vm-install-13.png)

25. In this step, you should provide your server and user information and also the credentials. for **Name, Server's Name and Username** you could use the below information and fill the password as you will. **\*\*\* Please don't forget username and the password. We need them to login to our server!**

![ubuntu-04-vm-install-14](ubuntu/ubuntu-04-vm-install-14.png)


26. Using **Up arrow key** highlight the **Install OpenSSH server** and press **Space**. Then select **Done** and press **Enter**.

![ubuntu-04-vm-install-15](ubuntu/ubuntu-04-vm-install-15.png)

27. We don't want more changes. So please select **Done** and press **Enter**.

![ubuntu-04-vm-install-16](ubuntu/ubuntu-04-vm-install-16.png)

28. Once setup completed, it will download and install security updates. Please wait until you see **[ reboot now ]** in the bottom of the page.

![ubuntu-04-vm-install-17](ubuntu/ubuntu-04-vm-install-17.png)

29. Congradulation. You have successfully installed **Ubuntu Server 18.04 LTS**. Select **Reboot now** and press **Enter**.

![ubuntu-04-vm-install-18](ubuntu/ubuntu-04-vm-install-18.png)

30. When you see this message _'Please remove the installation medium, then press ENTER'_, press **Enter** and wait for reboot.

![ubuntu-04-vm-install-19](ubuntu/ubuntu-04-vm-install-19.png)

1.  To login to your server, you should use the credentials you have set in **Step 25**. Next to **login:**, type your username (my username is **student**), then press **Enter**. **\*\*\*Note**: Password will not show while you are typing. Don't worry! Type your password and press **Enter**.

![ubuntu-05-vm-login-1](ubuntu/ubuntu-05-vm-login-1.png)

32. If you provide the valid credentials, you will be able to login to your server.

![ubuntu-05-vm-login-2](ubuntu/ubuntu-05-vm-login-2.png)

33. Check [Login to server using SSH](#login-to-server-using-ssh) guide. We will login to our server using **SSH**. In real-wolrd scenarios, no one logins to server directly from the console and we do it too. So check [this guide](#login-to-server-using-ssh) and do the same for your future logins.

### Installing *CentOS 7* on VirtualBox



### Login to server using SSH
##### Windows OS:
  1. Using PowerShell on Windows 10 or above.
     - Open **PowerShell** and use this schema **ssh username@server-address** to login to your server using **SSH**. In my case `ssh student@192.168.56.103`.
     - **\*\*\*Note**: First time, you might be asked to trust the host. Type **yes** and press **Enter** to login
     ![ssh-powershell-1](ssh/ssh-powershell-1.png)
     - Then provide your password and press **Enter**.
     - If the provided infromation are correct, you will see the below page:
     ![ssh-powershell-2](ssh/ssh-powershell-2.png)

  2. Using Putty software.
     - Download Putty from [this link](https://the.earth.li/~sgtatham/putty/latest/w64/putty-64bit-0.76-installer.msi).
     - Install it using the installation wizard.
     - Open **Putty** software.
     - Provide the IP address. In my case `192.168.56.103`
     ![ssh-putty-1](ssh/ssh-putty-1.png)
     - **\*\*\*Note**: First time, you might be asked to trust the host. Press the **Yes** button to trust it.
     ![ssh-putty-2](ssh/ssh-putty-2.png)
     - Provide your username and password. My username is `student`
     ![ssh-putty-3](ssh/ssh-putty-3.png)
     - If the provided infromation are correct, you will see the below page:
     ![ssh-putty-4](ssh/ssh-putty-4.png)

##### Linux & macOS:
- Open **Terminal** and use this schema **ssh username@server-address** to login to your server using **SSH**. In my case `ssh student@192.168.56.103`.
- **\*\*\*Note**: First time, you might be asked to trust the host. Type **yes** and press **Enter** to login
![ssh-linux-mac-1](ssh/ssh-linux-mac-1.png)
- Then provide your password and press **Enter**.
- If the provided infromation are correct, you will see the below page:
![ssh-linux-mac-2](ssh/ssh-linux-mac-2.png)
