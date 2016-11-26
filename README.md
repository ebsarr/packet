# packet, a CLI tool to manage Packet services
[![Go Report Card](https://goreportcard.com/badge/github.com/ebsarr/packet)](https://goreportcard.com/report/github.com/ebsarr/packet)

packet is a CLI tool to manage [packet.net](https://www.packet.net) services. You can browse the help [here](doc/packet.md).

        .--~~~~~~~~~~~~~------.
       /--===============------\\
       | |^^^^^^^^^^^^^^^|     |
       | |               |     |
       | | > packet v2.0 |     |
       | |               |     |
       | |_______________|     |
       |                   ::::|
       '======================='
       //-"-"-"-"-"-"-"-"-"-"-\\\\
      //_"_"_"_"_"_"_"_"_"_"_"_\\\\
      [-------------------------]
      \\_________________________/

# Installation

Install with `go get`
```
$ go get -u github.com/ebsarr/packet
```
`packet` executable will be installed in `$GOPATH/bin` or `$GOBIN`
<br>
<br>
You can also run it in a docker container:
```
$ docker run -it ebsarr/packet bash
```

Try the help
```
$ packet -h
CLI tool to manage packet.net services

Usage:
  packet [flags]
  packet [command]

Available Commands:
  admin           Manage projects, ssh keys, etc...
  baremetal       Manage server devices.
  network         Manage packet network services
  storage         Manage your storages

Flags:
  -k, --key string       API key
  -p, --profile string   Profile name (default "default")
  -v, --version          Show version and exit

Use "packet [command] --help" for more information about a command.
```

# Getting started

## Configure your API key by adding a profile
Command syntax: `packet admin add-profile`
```
$ packet admin add-profile
Enter your API key [ *****y7wi ]: <APIKEY>
Enter your default project ID [  ]: <Project ID>
```
This command will add a profile named "default".
<br>
**NOTE:** Without your API key configured, you'll need to specify it in every command in the form: `packet --key <APIKEY> <command> <subcommand> <flags>`. You can also optionnaly configure a default project ID.

If you have multiple accounts, or if you are working on multiple projects, you can set profiles to make it easy to switch between accounts or projects. After setting multiple profiles, you can use `-p` or `--profile` option to switch between accounts and projects.

### Setting a profile
Here I'm creating a new profile named `ebsarr`
```
$ packet admin add-profile -p ebsarr
Enter your API key [ *****y7wi ]: <APIKEY>
Enter your default project ID [  ]: <Project ID>
```
Without the `-p` switch, the `add-profile` command sets up a profile named "default". You can view your profiles with the `list-profiles` command:
```
$ packet admin list-profiles
NAME      	APIKEY                          	DEFAULT PROJECT
----      	------                          	---------------
default   	XMiR----------------------------	13935598-d08c-4bd8------------------
ebsarr    	XMiR----------------------------	e30d25da-728d-47fe------------------
```
Now I can switch easily between projects when running the `packet` command:
```
$ packet -p ebsarr baremetal list-devices
[]
```
Without the `-p` option, the default profile will be used:
```
$ packet baremetal list-devices
[
    {
        "id": "69148e7c-44e1-4b4a-ac1a-f9e08b552fe8",
        "href": "/devices/69148e7c-44e1-4b4a-ac1a-f9e08b552fe8",
        "hostname": "ebsarrtest01",
        "state": "active",
        "created_at": "2016-07-15T07:16:28Z",
        "updated_at": "2016-07-15T07:31:27Z",
        "billing_cycle": "hourly",
        "ip_addresses": [
...
``` 

## Manage your projects

### Create a new project
```
$ packet admin create-project --name "My Brand New Project"
{
    "id": "52a57c4b-5e28-4f79-9133-f7c953fa0e35",
    "name": "My Brand New Project",
    "created_at": "2016-05-29T15:58:27Z",
    "updated_at": "2016-05-29T15:58:27Z",
...
```

### View all your projects
```
$ packet admin list-projects
[
    {
        "id": "13935598-d08c-4bd8-8281-3196b6379452",
        "name": "Demo",
        "created_at": "2016-04-27T02:42:12Z",
        "updated_at": "2016-05-29T15:04:20Z",
...
```
### View project by ID
```
$ packet admin list-project --project-id 52a57c4b-5e28-4f79-9133-f7c953fa0e35
{
    "id": "52a57c4b-5e28-4f79-9133-f7c953fa0e35",
    "name": "My Brand New Project",
    "created_at": "2016-05-29T15:58:27Z",
    "updated_at": "2016-05-29T15:58:27Z",
...
```

## Manage your devices

### Create a new device
```
$ packet baremetal create-device --project-id <Project ID>\
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
```
$ packet baremetal list-device --device-id 2f027ea7-e5e9-4768-b2ba-fc03f3fa2b88
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

## Managing SSH keys
### Create SSH key
First generate the key
```
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/sarre27/.ssh/id_rsa): ./id_rsa
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
...
```
Register to packet
```
$ packet admin create-sshkey --label bass2@packet --file id_rsa.pub
{
    "id": "02b76cb4-ebeb-4eee-8d5d-a6d744aa793b",
    "label": "bass2@packet",
    "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDGENCZL3i+rrBZb2sQDt6H2xyziJuNB2ylFOnEywlBNdn0bGvXBFpitpOFmumYU7d0c2i2AYqDP8JgCu1sBiv1hENdMoMe8nmJRR8kjz7q+xWP18R+hYCAgvLEsfhW6fp7xpwK4cYTq07axg0hXSg+Lu8zjcm79EfucXPcNeYLq+27yPI3r8WnglMxfxhz7woBOlAjru6UBoCo+kpXwxA987rHoZEOeBpKGe8OzNt4Vqd8XOEYfwslCbsKBNfMJT0Eh/HeQ6WOOpZMsf3p6ufAocCsP5BeyZsChkuyNeNnPjMQG6chL8GzbMjb463IyiRJHkrM0zQPB+ysIhU5o8CP",
    "fingerprint": "f2:ca:e8:ea:dc:bf:a3:67:ad:4c:21:7a:92:bb:ed:ff",
    "created_at": "2016-05-30T05:36:29Z",
    "updated_at": "2016-05-30T05:36:29Z",
    "user": {
        "id": "",
        "href": "/users/0193dabc-a51d-4190-a01e-b270d664db3a"
    },
    "href": "/ssh-keys/02b76cb4-ebeb-4eee-8d5d-a6d744aa793b"
}
```
### View registerd SSH keys
```
$ packet admin list-sshkeys
[
    {
        "id": "76dee787-07a7-4510-9760-d27d0c51531e",
        "label": "bass2@pckt",
        "key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFj5UBo8zP4Uet9xsHn83HG9g7fFeGuSLHcjIWpE/WTLqV76H2DA9PCmmUivYi8f/VMSUpKSET2zC8wrzamepPrWXGqoFsB/I1za8ghcjhzN66und5dOPJzH2C+QihU1aH5cBoaPekb03HSK3qKUa1lCa0nmtdkvWxsspC42zXnf4PqOKkolang7sUe3tj4QvajEFxnWozcOc2Nfukv4q9Ml5vqSePCnYWby1oXWafezzQqThNe0+1DUYYKRe46D11E9wPamQiU6v7edlCSrzD2lpKxWGdtXPNZJYivpUdvCoj38sKuEfATsFZKd/HAiuKVe/o7Mpp3ZccbnLNgV1R",
        "fingerprint": "1f:e9:bf:0d:66:10:5f:6c:47:70:0b:70:c0:d5:db:7c",
        "created_at": "2016-05-27T10:17:47Z",
        "updated_at": "2016-05-27T10:17:47Z",
...
```
### Delete key by ID
```sh
$ packet admin delete-sshkey --key-id 02b76cb4-ebeb-4eee-8d5d-a6d744aa793b
$ echo $?
0
$ packet admin list-sshkey --key-id 02b76cb4-ebeb-4eee-8d5d-a6d744aa793b
Error: GET https://api.packet.net/ssh-keys/02b76cb4-ebeb-4eee-8d5d-a6d744aa793b: 404 Not found
```

## For more help
Type `packet -h` in your console or browse the help [here](doc/packet.md).

## License
[MIT](LICENSE)

## Changelog

| Version | Description |
|---------|-------------|
| **2.0**     | - Changed command structure, many bugs fixed.
| **1.3**     | - Can now delete profiles |
| **1.2**     | - Added profile support: use `--profile` option to switch between profiles | 
|         | - `ssh` command now reads keys from files, use `--file` instead of `ssh-key` to read from files.         |
| **1.1**     | Added support for userdata |
| **1.0**     | First version |