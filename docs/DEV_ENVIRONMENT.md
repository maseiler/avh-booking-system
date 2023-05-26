# Getting started with Development
## System Requirenments
We develop with `VS-Code`, but you are free to use your preferred IDE. Keep in mind, that we might not be capable to help you getting started if you are using another IDE.

- [X] Local Admin rights
- [X] VS-Code
- [X] Git
- [x] Go Lang
- [X] SQL Database
- [X] Node.js
- [X] Local DNS Entry (optional)
- [X] Local trusted SSL Certificates
- [X] A `launch.json` File
- [X] Sample entrys for your Database

### Installing VS-Code
[Download VS-Code](https://code.visualstudio.com/download) and Install it on your machine.

If you want to do it via CLI [follow this Tutorial](https://www.makeuseof.com/how-to-install-visual-studio-code-ubuntu/) since the package name is simply `code`. And that is apparently an incredibly stupid name for a package and can get misinterpreted sometimes.

### Installing Git
#### On Linux
```
sudo apt install git-all
```

#### On Windows
[Download Git](https://git-scm.com/download/win) and install it.

### Installing Go Lang
Follow the instructions from
[Go Website](https://go.dev/doc/install) to install Go Lang on your machine.

### SQL Database
You can use almost any SQL Server you like (As long as it does not produce errors). We Reccomend MySQL or MariaDB.

#### On Linux
[Visit MySQL installation Guidelines](https://dev.mysql.com/doc/refman/8.0/en/linux-installation.html) to get your local SQL Server started.

#### On  Windows
You can [Dwonload xampp](https://www.apachefriends.org/de/download.html) and install it on your machine.

#### Creating the Database
You will need a Database and a DB-User with CRUD-Permissions on that DB corresponding with the credentials specified in the environment variables.
The values can be found in the `.vscode/launch.json` File.

##### On Linux
Please visit the MySQL Documentation on how to create a new DB and User.

##### On Windows
1. Start Xampp
1. Start the Apache Web Server
1. Start the MySQL Server
1. Click on `Admin` next to the MySQL Server
1. Add the DB and User via the `phpMyAdmin` Web-GUI.
1. Stop the Apache Server and leave the MySQL Server running

### Installing Node.js
#### On Linux
[Visit this site](https://nodejs.org/en/download/package-manager) to learn how to install Node.js via the CLI.

#### On Windows
[Download Node.js](https://nodejs.org/en/download) and install it on your machine.

### Local DNS Entry
Some Browsers do not allow a SSL Certificate for localhost, since it is a loopback and not owned by anyone. So the ownership cannot be certified.

In order to develop locally without this getting in our way, we need a local DNS entry to reroute our localhost.

#### On Linux
Go to `/etc/` and edit your `hosts`-File.
Add this entry at the end of the file:
```
127.0.0.1 avhbs.local localhost
```
Save it and restart your Browser. Whenever you enter the Domain `avhbs.local` you should get redirected to your localhost.

#### On Windows
Navigate to `C:/Windows/System32/drivers/etc` and copy the `hosts`-File to your Desktop. From there, open it in Notepad and add this entry at the end:
```
127.0.0.1 avhbs.local localhost
```
Save the File and make shure to set the charset to UTF-8 (without DOM). Now **copy** the File back to the original Path (not move! copy it) and overwrite the original.
Now restart your Browser and you should see your localhost sites, when you enter the domain `avhbs.local`.

##### Troubleshooting
- Restart PC after changing the hosts file
- open CMD with Admin Rights and enter `ipconfig /flushdns`

### Trusted SSL Certificates
SSL is a Standard in Web-Technology and should be the standard for development as well. Also it is needed to run PWAs and since Version 2 we are developing a PWA.

#### On Linux
Ask @maseiler

#### On Windows
You can choose if you want it quick and easy or if you want to spend a weekend to set this up.
##### Do everything on your own (not recommended)
- follow [this Tutorial](https://zeropointdevelopment.com/how-to-get-https-working-in-windows-10-localhost-dev-environment/)
- Set up your own Certificate Authority
- Generate and Sign your own Certificates for your development environment
- Copy the newly generated certs in the folder `server/configs` (yes, you have to create the folder first)
##### Use a helper software (recommended)
- follow [this Tutorial](https://realtimelogic.com/articles/How-to-act-as-a-Certificate-Authority-the-Easy-Way)
- Trust the newly generated RootCA
- Generate Certs for your development environment
- Copy the newly generated certs in the folder `server/configs` (yes, you have to create the folder first)

After you have successfully created and placed your certificates, you can modify the `.vscode/launch.json`-File and enter the Paths to your newly generated Certificates.
This will store the Paths as environment variables.

### The `Launch.json` File
Speaking of the launch file...

If you want to work on this project, you will need the launch file. Just ask us and we will send it to you.

The File will be used to start the debugging Go Web Server with the correct Environment-Variables.

If you already have the file and followed through the Tutorial to this point, you just have to change the Frontend Path Variable to match your Local path of the Project followed by the folder `dist`.
Don't worry if you do not see the folder now. It will appear soon.

### Sample Data for your DB
You don't need sample Data. But you could run into some issues not having them.

If you want to get sample Data, please reach out to us and we will send you a sql dump with sample data.

You just need to put that dump in your DB and you should be ready to go.

## Getting Started
After your System is configured you will need to install all neccessary packages to run the application.

Run the following Commands from your project root:
```
cd client
npm install
```

## Run the Dev Server
### 1. Make a Build
Navigate to `/client` and run
```
npm run build
```
If you encounter an Error mentioning `ERR_OSSL_EVP_UNSUPPORTED` run this command 
```
export NODE_OPTIONS=--openssl-legacy-provider
```
and try to build again.

### 2. Start the MySQL DB
Start XAMPP or your MySQL installation. And start your mySQL Server.

### 3. Start the Server
Open the `.vscode/launch.json`-File and hit `F5` Key.
The Server should start in your Debug-Console and show Errors if something went wrong.

### 4. See what you are doing
Just open a Browser and navigate to `https://avhbs.local:8081` if you configured your local DNS Route.

If you did not configure the route, you can access the server by entering this domain `https://127.0.0.1:8081`.

Happy Hacking 😘

Prost 🍻