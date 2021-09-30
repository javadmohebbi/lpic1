### Installing *CentOS 7* on VirtualBox
1. Download ISO image from [this link](http://mirrors.standaloneinstaller.com/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-NetInstall-2009.iso)

2. Open **VirtualBox** and go to **Machine** menu and then click on **New**. You can press **CTRL+N**. Create virtual machine window appears.
   - **Name**: lpic-centos
   - **Machine Folder**: Change this if this location has not enough space otherwise leave it as it's default value
   - **Type**: Linux
   - **Version**: Red Hat (64-bit)

![centos-01-vm-wizard-1](./centos/centos-01-vm-wizard-1.png)


3. Click **next**. In this window you should config the virtual machine memory. 1024MB (1GB) memory is enough for our lab. But you can change it according to your need.

![centos-01-vm-wizard-2](centos/centos-01-vm-wizard-2.png)

4. Click **next** and you will see Hard disk confoguration window. Click on **create** button and leave everything default.

![centos-01-vm-wizard-3](centos/centos-01-vm-wizard-3.png)

5. Leave everything default in "Hard disk file type" window and click **next**.

![centos-01-vm-wizard-4](centos/centos-01-vm-wizard-4.png)

6. In this window, select **Dynamically allocated** and click **next**.

![centos-01-vm-wizard-5](centos/centos-01-vm-wizard-5.png)

7. You can change the hard disk size. The default value is enough. But you can increase it as you want but don't decrease it. Then click **create** button.

![centos-01-vm-wizard-6](centos/centos-01-vm-wizard-6.png)


8. Now you will see the newly created virtual machine (**lpic-centos**).

![centos-02-vm-created](centos/centos-02-vm-created.png)


9. **Right click** on **lpic-centos** machine and click **settings**. Then go to **Network** Section and select **Adapter 2** tab. Do the following changes:
    - Check the **Enable Network Adapter** check box.
    - Select **Host-only adapter** for **Attached To** drop-down.
    - Select **vboxnet0** for **Name** drop-down.

![centos-03-vm-setting-1](centos/centos-03-vm-setting-1.png)

10. Now goto **Storage** section and below **Controller: IDE** select **Empty**. Then click on ![disk-icon](centos/disk-icon.png) (**Disk icon**) and from the pop-up menu, select **Choose a disk file**.

![centos-03-vm-setting-2](centos/centos-03-vm-setting-2.png)

11. Now do the following
    - Select your **centos ISO image** and click **open**.
    - Then click **OK** to close the virtual mahchine settings window.

![centos-03-vm-setting-3](centos/centos-03-vm-setting-3.png)

12. Select **lpic-centos** machine and click on **Start** button.

![centos-04-vm-install-1](centos/centos-04-vm-install-1.png)

13.  The virtual machine starts and boots from the provided ISO image. Using **UP arrow key**, select **Install CentOS 7** and press **Enter**.

![centos-04-vm-install-2](centos/centos-04-vm-install-2.png)

14. Choose your language and click on **Continue** button.

![centos-04-vm-install-3](centos/centos-04-vm-install-3.png)

15. Scroll-down a bit and click on **NETWORK & HOST NAME**.

![centos-04-vm-install-4](centos/centos-04-vm-install-4.png)

16. Click on **Ethernet (enp0s3)** and then on the right side change **OFF** to **ON** to enable it

![centos-04-vm-install-5](centos/centos-04-vm-install-5.png)

16. Click on **Ethernet (enp0s8)** and then on the right side change **OFF** to **ON** to enable it. Also click on **Configure...** button to configure a static IPv4 for this ethernet.
    - **\*\*\* Note**: Before go further, write down the current IP address that our network adapter got from DHCP server (In my case this IP address is `192.168.56.104/24` ).

![centos-04-vm-install-6](centos/centos-04-vm-install-6.png)

17. Select **IPv4 Settings** tab. Change the **Method** to **Manual** and click on **Add** button. Now you can fill the **Address** and **Netmask** field. Then click on **Save** button.

![centos-04-vm-install-7](centos/centos-04-vm-install-7.png)

18. Change the `localhost.localdomain` to **lpic-centos.localdomain** and press the **Apply** button. Then click on **Done** to save the configuration.

![centos-04-vm-install-8](centos/centos-04-vm-install-8.png)


19. Click on **INSTALLATION SOURCE** to define it.

![centos-04-vm-install-9](centos/centos-04-vm-install-9.png)

20. Type **mirror.centos.org/centos/7/os/x86_64** in the box and then click on **Done** button.

![centos-04-vm-install-10](centos/centos-04-vm-install-10.png)

21. Click on **SOFTWARE SELECTION** to choose type of installation.

![centos-04-vm-install-11](centos/centos-04-vm-install-11.png)

22. On the left side, select **Minimal Install** and the right side select **Compatibility Libraries** and **System Administration Tools**. Then click on **Done** to save the configuration.

![centos-04-vm-install-12](centos/centos-04-vm-install-12.png)

23. Click on **INSTALLATION DESTINATION** to configure hard disk partitioning.

![centos-04-vm-install-13](centos/centos-04-vm-install-13.png)

24. We use automatic method, so check the **Automatically configure partitioning** and click on **Done** button.

![centos-04-vm-install-14](centos/centos-04-vm-install-14.png)

25. Now we can start the installation by clicking on **Begin Installation**

![centos-04-vm-install-15](centos/centos-04-vm-install-15.png)

26. To set the **root** account's password, click on **ROOT PASSWORD**.

![centos-04-vm-install-16](centos/centos-04-vm-install-16.png)

27. Define the password and press **Done** button. **Don't forget** this password. You might need it in the future.

![centos-04-vm-install-17](centos/centos-04-vm-install-17.png)

28. We don't want to use root account for our works. We must create a user with administrative access. To do so, click on **USER CREATION** button.

![centos-04-vm-install-18](centos/centos-04-vm-install-18.png)

29.  In this step, you should provide your server and user information and also the credentials. for **Name**, **Username** you could use the below information and fill the password as you wish. **\*\*\* Please don't forget username and the password. We need them to login to our server!**. Also make sure the **Make this user administrator** is checked. Then click on **Done** button to save the configuration.

![centos-04-vm-install-19](centos/centos-04-vm-install-19.png)

30. When installation ends, the **Reboot** button will appear. Click on **Reboot** buton and wait for machine to reboot.

![centos-04-vm-install-20](centos/centos-04-vm-install-20.png)

31. To login to your server, you should use the credentials you have set in **Step 25**. Next to **login:**, type your username (my username is **student**), then press **Enter**. **\*\*\*Note**: Password will not show while you are typing. Don't worry! Type your password and press **Enter**.

![centos-04-vm-install-21](centos/centos-04-vm-install-21.png)

32. We want to update our repository and upgrade the applications which need it. So, run this command `sudo yum update` to do it. You might be asked to enter your password when you use **sudo**. So enter it and press **Enter**.

![centos-04-vm-install-22](centos/centos-04-vm-install-22.png)

33. If downlaod needed, It will ask you to confirm it by typing **y** and pressing **Enter** key.

![centos-04-vm-install-23](centos/centos-04-vm-install-23.png)

34.  To confirm newly downloaded firmware key, type **y** and press **Enter** key.

![centos-04-vm-install-24](centos/centos-04-vm-install-24.png)

35. When you see **Complete!**, it means our upgrade is completed and no more action is required.

![centos-04-vm-install-25](centos/centos-04-vm-install-25.png)

36. To be sure that our network adapters are configured properly, run `nmcli` command. Sometime after CentOS installtion, 2nd adapter is disabled as you can see in the below image

![centos-04-vm-install-26](centos/centos-04-vm-install-26.png)

37. To fix that run the command `sudo nmtui`.

![centos-04-vm-install-27](centos/centos-04-vm-install-27.png)

38. Highlight **Activate a connection** and press **Enter**.

![centos-04-vm-install-28](centos/centos-04-vm-install-28.png)

39. Highlight **enp0s8** and then using **Right arrow key**, highlight the **\<Activate\>** and hit **Enter**.

![centos-04-vm-install-29](centos/centos-04-vm-install-29.png)

40. Choose **\<Back\>**

![centos-04-vm-install-30](centos/centos-04-vm-install-30.png)

41. To quit this app, select **Quit** and press **Enter**.

![centos-04-vm-install-31](centos/centos-04-vm-install-31.png)

42. Run **`nmcli`** again to confirm that all network adapters are active and has the IP address we configured during the installation. You could use this **IP address** to login to the server using **SSH**.

![centos-04-vm-install-32](centos/centos-04-vm-install-32.png)


43.  Run **`sudo yum install nano git -y`** to insall both **nano** editor and **git** application.



44.  Check [Login to server using SSH](SSH-README.md#login-to-server-using-ssh) guide. We will login to our server using **SSH**. In real-wolrd scenarios, no one logins to server directly from the console and we do it too. So check [this guide](SSH-README.md#login-to-server-using-ssh) and do the same for your future logins.











