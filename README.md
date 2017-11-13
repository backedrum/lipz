###LIPZ - a tool for network monitoring and troubleshooting.

- backend - golang,
- frontend - vue.js

###How it works?
1. Compile go server and start server binary under sudo.
```
go build server.go
sudo ./server
```
2. Start UI from /client folder. For detailed instructions follow [this link](client/README.md)
```
npm run dev
```
3. Navigate to "Devices" link -> a table with available devices will be displayed.
4. Click on device row -> capture dialog will appear. Press "Start Capture!" button.
5. Live capture on selected device will be taken for 15 seconds. Result will be presented afterwards.

###Main libraries used:
- Lib pcap from https://github.com/google/gopacket <br>
- Table component from https://github.com/xaksis/vue-good-table <br>
- Start go-vue.js template from https://github.com/markcheno/go-vue-starter <br>
