# packet, a CLI tool to manage Packet services

packet is a CLI tool to manage [packet.net](https://www.packet.net) services. You can browse the documetation [here](doc/packet.md).

# Installation

## Download the binary for your platform

Download relevant binary from the following links:

|Platform| Binary |
|:-------|:-------|
|Windows(64bit)|[Download](bin/win_64)|
|Mac(64bit)|[Download](bin/osx_64)|
|Linux(64bit)|[Download](bin/lin_64)|

# Getting started

## Configure your API key
Command syntax: `packet configure`
```sh
$ packet configure
Enter your API key [ *****y7wi ]: <APIKEY>
```
**NOTE:** Without your API key configured, you'll need to specify it in every command in the form: `packet --key <APIKEY> <command> <subcommand> <flags>`

## Manage your projects

### Create a new project
```sh
$ packet project create --name "My Brand New Project"
{
    "id": "52a57c4b-5e28-4f79-9133-f7c953fa0e35",
    "name": "My Brand New Project",
    "created_at": "2016-05-29T15:58:27Z",
    "updated_at": "2016-05-29T15:58:27Z",
...
```

### View all your projects
```sh
$ packet project listall
[
    {
        "id": "13935598-d08c-4bd8-8281-3196b6379452",
        "name": "Demo",
        "created_at": "2016-04-27T02:42:12Z",
        "updated_at": "2016-05-29T15:04:20Z",
...
```
### View project by ID
```sh
$ packet project list --project-id 52a57c4b-5e28-4f79-9133-f7c953fa0e35
{
    "id": "52a57c4b-5e28-4f79-9133-f7c953fa0e35",
    "name": "My Brand New Project",
    "created_at": "2016-05-29T15:58:27Z",
    "updated_at": "2016-05-29T15:58:27Z",
...
```

## Manage your devices

### Create a new device
```sh
$ packet device create --project-id <Project ID>\
	--billing hourly \
	--facility ewr1\
	--hostname CreatedFromCli\
	--os-type coreos_stable\
	--plan baremetal_0

Provisioning of device successfully started...
 [ 2016-05-29T14:58:19Z ] "%device%" (Type 0) was deployed to project "%project%" by %user%
 [ 2016-05-29T14:59:05Z ] Network configured with addresses 147.75.199.11, 2604:1380:0:7b00::17, and 10.100.131.151
 [ 2016-05-29T14:59:16Z ] Configuration written, restarting device
 [ 2016-05-29T14:59:36Z ] Device connected to DHCP system
 [ 2016-05-29T15:01:54Z ] Installation finished, rebooting server
 [ 2016-05-29T15:02:50Z ] Provision complete! Your device is ready to go.
{
    "id": "2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88",
    "href": "/devices/2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88",
    "hostname": "CreatedFromCli",
    "state": "active",
    "created_at": "2016-05-29T14:58:19Z",
    "updated_at": "2016-05-29T15:02:50Z",
    "billing_cycle": "hourly",
...
```

### View device by ID
```sh
$ packet device list --device-id 2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88
{
    "id": "2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88",
    "href": "/devices/2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88",
    "hostname": "CreatedFromCli",
    "state": "active",
    "created_at": "2016-05-29T14:58:19Z",
    "updated_at": "2016-05-29T15:02:50Z",
    "billing_cycle": "hourly",
...
```

## For more help
Type `packet -h` in your console or browse the documetation [here](doc/packet.md).

## License
[MIT](LICENSE)