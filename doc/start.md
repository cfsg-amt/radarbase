# radarbase troubleshoots
## Run the Radarbase Server

Open terminal, run the following command to start the radarbase server.

```shell
cd ~/Desktop/Radar/Radar/radarbase/ | go run cmd/radarbase/main.go
```

## Testphase 1: Local Server Status
After starting the radarbase server and see the loading successfull output. Minimize the terminal.
Open the browser, enter: `localhost:8996/api/v1/minmax/StkSH`
If it get response like this:
```
{"max":{"5年平均市盈率":643.3,"name":0,"valid":0,"中位目標價變化":24.35,"中文新聞（正負面）情緒平均指標":0.97,"交易價值":5035212338,"交易股數":506974098,"保力加通道  (下線) (20日)":1615.717,"保力加通道 (上線)...
```
It means the local server runs successfully, and can be accessed locally.
If not, please back to that terminal running radarbase, check its log.

## Testphase 2: LAN Server Status
First, in the server, try to get its LAN address. In Windows 10, please following [this](https://support.microsoft.com/en-us/windows/find-your-ip-address-in-windows-f21a9bbc-c582-55cd-35e0-73431160a1b9#Category=Windows_10) instruction.
Then using another device that **in the same network** (wireless or wired) with that server.

Open the browser on that device, and try to access `<LAN-ADDR>:8996/api/v1/minmax/StkSH`, where the `<LAN-ADDR>` is the LAN address of the server.
See if you can get the same response as before in Testphase 1.
If not, they issue may lie in the firewall of your server, try to turn it off and run that again.

Or open another terminal window, try to ping that device:

```shell
ping <LAN-ADDR>
```

### Check and Allow a Port through Windows Firewall:

1. **Open Windows Firewall Settings**:
Press the Windows key, type "firewall," and select "Windows Defender Firewall" or just "Firewall & network protection" in Windows Security.

2. **Advanced Settings**:
In the Windows Defender Firewall window, click on "Advanced settings" on the left-hand side. This opens the Windows Defender Firewall with Advanced Security window. (You might need administrator privileges to access this.)

3. **Inbound Rules**:
In the Advanced Security window, click on "Inbound Rules" in the left pane.

4. **Check Existing Rules**:
Look through the list of Inbound Rules for any rules related to port 8996.
The rules will be listed by name, and you can see the details, including the specific port number, in the "Ports" column.

5. **Creating a New Rule (if necessary):**
If there’s no existing rule for port 8996, you'll need to create one. To do this, click "New Rule..." on the right-hand side.
Select "Port" as the Rule Type and click "Next."
Choose "TCP" or "UDP" based on your application's requirement (TCP is more common for web services). Then, select "Specific local ports," enter "8996," and click "Next."
Choose "Allow the connection" and click "Next."
Ensure all three profiles (Domain, Private, Public) are checked as per your requirement and click "Next."
Give the rule a name (like "Radarbase Port 8996") and an optional description, then click "Finish."

6. **Review and Apply**:
After creating the new rule, it should appear in the list of Inbound Rules. Make sure it is enabled (checked).
After adjusting the firewall settings, it may be a good idea to restart the device (reboot) to ensure the new settings take effect.

## Testphase 3: DDNS and Port Forwarding
If the first two phases has passed, the last step would be set up your **DDNS** and **Port Forwarding** on your router, in order to let other device outside your network can access the radarbase server.

### DDNS
In the setting page of your router, try to assign a global domain name that forwarding to your lan address of the server.

### Port Forwarding 
After DDNS is set up, make sure to add a Port Forwarding run to forward a specific port (e.g. `2222`) to port `8996` (radarbase) on your machine.
Final test, open a browser try enter: `<DDNS-ADDR>:<PORT>/api/v1/minmax/StkSH` to check whether it can get some response.


## HTTPS Setting (win-acme)
This guide explains how to generate an SSL certificate for a Fully Qualified Domain Name (FDQN) on a Windows Server machine using the win-acme tool. 
**[Credit and Source](https://docs.fintechos.com/Platform/22.1/AdminGuide/Content/Installation/generateSSL.htm)**

### Prepare
1. Make sure that ports 80 and 443 on your environment allow Internet connectivity.
2. Download the [win-acme](https://github.com/win-acme/win-acme/releases/download/v2.2.6.1571/win-acme.v2.2.6.1571.x64.pluggable.zip).
3. Run wacs.exe as administrator.

### Create a certificate with full options
```
N: Create certificate (default settings)
M: Create certificate (full options)
R: Run renewals (0 currently due)
A: Manage renewals (1 total)
O: More options...
Q: Quit 
Please choose from the menu: M 
```
Type `M` to create a certificate with full options

### Input your domain name and serve HTTP-01 test from memory
```
Please specify how the list of domain names that will be included in the
certificate should be determined. If you choose for one of the "all bindings"
options, the list will automatically be updated for future renewals to
reflect the bindings at that time.
1: Read bindings from IIS
2: Manual input
3: CSR created by another program
C: Abort 
How shall we determine the domain(s) to include in the certificate?: 2 
```
Type 2 to manually input the domain names included in the certificate.

```
Description:        A host name to get a certificate for. This may be a
                    comma-separated list. 
Host: vm-customer360-dev.westeurope.cloudapp.azure.com
Source generated using plugin Manual: vm-customer360-dev.westeurope.cloudapp.azure.com 
Friendly name '[Manual] vm-customer360-dev.westeurope.cloudapp.azure.com'. <Enter> to accept or type desired name: vm-customer360-dev.westeurope.cloudapp.azure.com 
```
Enter the domain name you want to use in this server.
**`l45411e1993.tplinkdns.com`**

```
The ACME server will need to verify that you are the owner of the domain
names that you are requesting the certificate for. This happens both during
initial setup *and* for every future renewal. There are two main methods of
doing so: answering specific http requests (http-01) or create specific dns
records (dns-01). For wildcard domains the latter is the only option. Various
additional plugins are available from https://github.com/win-acme/win-acme/. 
1: [http-01] Save verification files on (network) path
2: [http-01] Serve verification files from memory
3: [http-01] Upload verification files via FTP(S)
4: [http-01] Upload verification files via SSH-FTP
5: [http-01] Upload verification files via WebDav
6: [dns-01] Create verification records manually (auto-renew not possible)
7: [dns-01] Create verification records with acme-dns (https://github.com/joohoi/acme-dns)
8: [dns-01] Create verification records with your own script
9: [tls-alpn-01] Answer TLS verification request from win-acme
C: Abort 
How would you like prove ownership for the domain(s)?: 2
```
Type 2 to serve the verification files from memory


### Download the RSA key
```
After ownership of the domain(s) has been proven, we will create a
Certificate Signing Request (CSR) to obtain the actual certificate. The CSR
determines properties of the certificate like which (type of) key to use. If
you are not sure what to pick here, RSA is the safe default. 
1: Elliptic Curve key
2: RSA key
C: Abort 
What kind of private key should be used for the certificate?: 2
```
Type 2 to select the RSA key type.

```
When we have the certificate, you can store in one or more ways to make it
accessible to your applications. The Windows Certificate Store is the default
location for IIS (unless you are managing a cluster of them).
1: IIS Central Certificate Store (.pfx per host)
2: PEM encoded files (Apache, nginx, etc.)
3: PFX archive
4: Windows Certificate Store
5: No (additional) store steps
How would you like to store the certificate?: 2 
```
Type 2 to store the certificate as PEM encoded files.


```
Description:        .pem files are exported to this folder.
File path: .
Description:        Password to set for the private key .pem file.
1: None
2: Type/paste in console
3: Search in vault 
Choose from the menu: 2
```
Type 2 to insert the path where you wish to store the certificates from the console. (E.g.: `C:\Users\john.doe\Documents`).

```
Description:        Password to set for the private key .pem file.
1: None
2: Type/paste in console
3: Search in vault 
Choose from the menu: 2
```
Type 1 to disable password protection for the private key file

```
1: IIS Central Certificate Store (.pfx per host)
2: PEM encoded files (Apache, nginx, etc.)
3: PFX archive
4: Windows Certificate Store
5: No (additional) store steps
Would you like to store it in another way too?: 5 
Installation plugin IIS not available: This step cannot be used in combination with the specified store(s) 
```
Type 5 to decline any additional store steps and finish this.







