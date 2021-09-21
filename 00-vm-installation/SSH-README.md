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
