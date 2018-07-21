package cmd

const version = "2.3.1"

const releaseNotes = `------------------------------------------------------------------------------|
| Version   | Description                                                     |
|-----------|-----------------------------------------------------------------|
| 2.3.1     | Allow creating volumes without a snapshot policy by passing     |
|           | --count 0                                                       |
|-----------|-----------------------------------------------------------------|
| 2.3       | Added support for the "organization" API                        |
|           | Support pagination for packet baremetal list-devices            |
|-----------|-----------------------------------------------------------------|
| 2.2.2     | Fixed a bug that blows away all tags on device updates          |
|           | Added --tags flag(not mandatory) to create-device and           |
|           | update-device commands                                          |
|-----------|-----------------------------------------------------------------| 
| 2.2.1     | Fixed compilation issue that emerged with updated Packet API    |
|-----------|-----------------------------------------------------------------|
| 2.2       | "update-device" command added; more options for "create-device" |
|           | command                                                         |
|-----------|-----------------------------------------------------------------|
| 2.1.3     | Fixed an issue that emerged with the updated Packet API         |
|-----------|-----------------------------------------------------------------|
| 2.1.2     | Fixed an issue that emerged with the updated Packet API         |
|-----------|-----------------------------------------------------------------|
| 2.1.1     | Bug fix around profile configuration. Now you can use --name    |
|           | or -n to configure and name a profile                           |
|-----------|-----------------------------------------------------------------|
| 2.1       | Add support for spot market                                     |
|-----------|-----------------------------------------------------------------|
| 2.0       | Changed command structure, many bugs fixed                      |
|-----------|-----------------------------------------------------------------|
| 1.3       | Can now delete profiles                                         |
|-----------|-----------------------------------------------------------------|
| 1.2       | Added profile support: use --profile option to switch between   |
|           | profiles.                                                       | 
|           | ssh command now reads keys from files, use --file instead of    |
|           | ssh-key to read from files.                                     |
|-----------|-----------------------------------------------------------------|
| 1.1       | Added support for userdata                                      |
|-----------|-----------------------------------------------------------------|
| 1.0       | First version                                                   |
-------------------------------------------------------------------------------`
