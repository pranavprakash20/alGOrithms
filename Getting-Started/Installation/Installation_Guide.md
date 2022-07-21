===============================================================================================================================================================================================================================================================================================================
# LINUX

Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz
(You may need to run the command as root or through sudo).

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

export PATH=$PATH:/usr/local/go/bin
Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.


Verify that you've installed Go by opening a command prompt and typing the following command:
go version
Confirm that the command prints the installed version of Go.

===============================================================================================================================================================================================================================================================================================================

# Mac

Open the package file you downloaded and follow the prompts to install Go.
The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

Verify that you've installed Go by opening a command prompt and typing the following command:
go version
Confirm that the command prints the installed version of Go.


===============================================================================================================================================================================================================================================================================================================

#WINDOWS

Open the MSI file you downloaded and follow the prompts to install Go.
By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

Verify that you've installed Go.
In Windows, click the Start menu.
In the menu's search box, type cmd, then press the Enter key.
In the Command Prompt window that appears, type the following command:
$ go version
Confirm that the command prints the installed version of Go.


===============================================================================================================================================================================================================================================================================================================


Please refer https://go.dev/doc/install for further details

Few good reads/views on installation 

1. https://www.youtube.com/watch?v=nWZRZUg2lOM
2. https://www.geeksforgeeks.org/go-programming-language-introduction/?ref=lbp
3. https://golangdocs.com/install-go-windows
