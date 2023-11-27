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

### Inbound Rules
Check and Allow a Port through Windows Firewall:

Open Windows Firewall Settings:

Press the Windows key, type "firewall," and select "Windows Defender Firewall" or just "Firewall & network protection" in Windows Security.
Advanced Settings:

In the Windows Defender Firewall window, click on "Advanced settings" on the left-hand side. This opens the Windows Defender Firewall with Advanced Security window. (You might need administrator privileges to access this.)
Inbound Rules:

In the Advanced Security window, click on "Inbound Rules" in the left pane.
Check Existing Rules:

Look through the list of Inbound Rules for any rules related to port 8996.
The rules will be listed by name, and you can see the details, including the specific port number, in the "Ports" column.
Creating a New Rule (if necessary):

If there’s no existing rule for port 8996, you'll need to create one. To do this, click "New Rule..." on the right-hand side.
Select "Port" as the Rule Type and click "Next."
Choose "TCP" or "UDP" based on your application's requirement (TCP is more common for web services). Then, select "Specific local ports," enter "8996," and click "Next."
Choose "Allow the connection" and click "Next."
Ensure all three profiles (Domain, Private, Public) are checked as per your requirement and click "Next."
Give the rule a name (like "Radarbase Port 8996") and an optional description, then click "Finish."
Review and Apply:

After creating the new rule, it should appear in the list of Inbound Rules. Make sure it is enabled (checked).
Restart the Radarbase Service:

After adjusting the firewall settings, it may be a good idea to restart the Radarbase service to ensure the new settings take effect.
Test Connectivity Again:

Try accessing the Radarbase server from another device on your network to see if the firewall changes resolved the issue.

If there are some response from your server, then probably you need to:
1. Check your router's firewall settings to ensure that incoming connections on port `2222` are allowed.
2. Ensure that any firewall on your Windows PC allows connections on port `8996`.

## Testphase 3: DDNS and Port Forwarding
If the first two phases has passed, the last step would be set up your **DDNS** and **Port Forwarding** on your router, in order to let other device outside your network can access the radarbase server.

### DDNS
In the setting page of your router, try to assign a global domain name that forwarding to your lan address of the server.

### Port Forwarding 
After DDNS is set up, make sure to add a Port Forwarding run to forward a specific port (e.g. `2222`) to port `8996` (radarbase) on your machine.
Final test, open a browser try enter: `<DDNS-ADDR>:<PORT>/api/v1/minmax/StkSH` to check whether it can get some response.
