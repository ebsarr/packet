# packet, a CLI tool to manage Packet services
[![Latest Version](https://img.shields.io/badge/release-v2.1.2-yellowgreen.svg)](https://github.com/ebsarr/packet/releases) [![Build Status](https://travis-ci.org/ebsarr/packet.svg?branch=master)](https://travis-ci.org/ebsarr/packet) [![Go Report Card](https://goreportcard.com/badge/github.com/ebsarr/packet)](https://goreportcard.com/report/github.com/ebsarr/packet) [![GoDoc](https://godoc.org/github.com/ebsarr/packet?status.svg)](https://godoc.org/github.com/ebsarr/packet)

packet is a CLI tool to manage [packet.net](https://www.packet.net) services. You can browse the help [here](doc/packet.md).

        .--~~~~~~~~~~~~~------.
       /--===============------\\
       | |^^^^^^^^^^^^^^^|     |
       | |               |     |
       | | > packet      |     |
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
`packet` executable will be installed in `$GOPATH/bin`
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
$ packet admin add-profile -n ebsarr
Enter your API key [ *****y7wi ]: <APIKEY>
Enter your default project ID [  ]: <Project ID>
```
Without the `-n` switch, the `add-profile` command sets up a profile named "default". You can view your profiles with the `list-profiles` command:
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

# Available commands

## `packet admin` : project management

Syntax: `packet admin [subcommand]`

|Subcommand | Description |
|-----------|-------------|
| `add-profile`         |Set default configs for the packet cli |
| `create-project`      |Create a new project |
| `create-sshkey`       |Configure a new SSH key |
| `delete-profile`      |Delete a profile |
| `delete-project`      |Delete a project |
| `delete-sshkey`       |Delete SSH key associated with the given ID |
| `list-facilities`     |View a list of facilities(packet DCs) |
| `list-os`             |View available operating systems |
| `list-plans`          |View available plans |
| `list-profiles`       |List configured profiles |
| `list-project`        |Retrieve a project by ID |
| `list-project-events` |View events by project ID |
| `list-projects`       |Retrieve all projects |
| `list-sshkey`         |View configured SSH key associated with the given ID |
| `list-sshkeys`        |View all configured SSH keys |
| `spot-prices`         |View spot market prices. For more details on the Packet spot market, see the [Packet spot market documentation](https://help.packet.net/technical/deployment-options/spot-market). |
| `update-project`      |Update a project |
| `update-sshkey`       |Update a SSH key: change the key or its label |

## `packet baremetal`: Manage server devices

Syntax: `packet baremetal [subcommand]`

|Subcommand | Description |
|-----------|-------------|
| `create-device`   | Create a new device |
| `delete-device`   | Delete a device |
| `list-device`     | Retrieve a device |
| `list-devices`    | Retrieve all devices in a project |
| `list-events`     | View events by device ID |
| `lock-device`     | Lock a device |
| `poweroff-device` | Power off a device |
| `poweron-device`  | Power on a device |
| `reboot-device`   | Reboot a device |
| `unlock-device`   | Unlock a device |

## `packet network`: Manage packet network services

Syntax: `packet network [subcommand]`

|Subcommand | Description |
|-----------|-------------|
| `assign-ip`                     | Assign IP address to a device by ID |
| `list-ip`                       | Retrieve IP address by ID |
| `list-ip-reservation`           | Retrieve a single IP reservation object by ID |
| `list-ip-reservations`          | Retrieve IP resevations for a single project |
| `remove-ip-reservation`         | Remove IP reservation |
| `request-more-ip-reservations`  | Request more IP space |
| `unassign-ip`                   | Unassign IP address from a device |

## `packet storage`: manage your storages

Syntax: `package storage [subcommand]`
<br>
Some of the snapshot related commands may not make sense. I will revisit the API and correct this on the next update.

|Subcommand | Description |
|-----------|-------------|
| `attach-volume`          | Attach a volume to a device |
| `clone-volume`           | Clone a volume or snapshot into a new volume |
| `create-snapshot`        | Create a snapshot of your volume |
| `create-snapshot-policy` | Create a snapshot policy |
| `create-volume`          | Create a volume |
| `delete-snapshot`        | Delete a snapshot of your volume |
| `delete-snapshot-policy` | Delete a snapshot policy |
| `delete-volume`          | Delete storage |
| `detach-volume`          | Detach a volume from a device |
| `list-snapshots`         | View a list of the current volume’s snapshots |
| `list-volume`            | Retrieve a volume by ID |
| `list-volume-events`     | View a list of the current volume's events |
| `list-volumes`           | Retrieve all volumes |
| `restore-volume`         | Restore a volume to the given snapshot |
| `update-snapshot-policy` | Update a snapshot policy |
| `update-volume`          | Update a volume |

## For more help
Type `packet -h` in your console or browse the help [here](doc/packet.md) to view command options.

## License
[![MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Changelog

| Version | Description |
|---------|-------------|
| **2.1.2**     | Fix compilation error. |
| **2.1.1**     | Bug fix around profile configuration. Now you can use `--name` or `-n` to configure and name a profile |
| **2.1**     | Add support for spot market |
| **2.0**     | Changed command structure, many bugs fixed.|
| **1.3**     | Can now delete profiles |
| **1.2**     | Added profile support: use `--profile` option to switch between profiles | 
|         | `ssh` command now reads keys from files, use `--file` instead of `ssh-key` to read from files.         |
| **1.1**     | Added support for userdata |
| **1.0**     | First version |
