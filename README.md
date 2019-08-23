# hostsd [![](https://img.shields.io/badge/Powered_by_Medium-555.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAIGNIUk0AAHolAACAgwAA+f8AAIDpAAB1MAAA6mAAADqYAAAXb5JfxUYAAAHKSURBVHjavJK9btNwFMWP/06stFLUOorT2FbiKk4NNHSqmvICSLQSvAILlAVGJNSJmZWOvAELQsBC2yQM/UBICGVqShNXxaaxLZymIpbstJeBVFgMDB040l3uPT/dD12OiHAZMVxSDMBdACaArwDu/MN7C0ALwCGAeyCiH/RHPhGViQh/hUZEbszXg9PtfiGis1iyQUSJGMSI6H2sfu553mdeKxRX8nn5dGJyIjsaSQMAIqr3ej1EUfSE5/n7jP0+h21Ze+/evB0wRVVmapsbLBgEVmyf1SiKbuy3WvOu6z7t9/sYDocIgsBu1OvniqJUWKmsR4IglD80GjYRhSMwwXHcC9/3n307OhIcx4Hv+1FtY9NMJJJXS7oesqncFM0YBuc4Xb190F6/aJlMJivXZmdznYO2c2ia2N3ZWfc8t2xcMZiUk4hNiiKTZRna9HTG7HTWANQu4GKxaFTmrtuDn4OXx9+Pn5d0PZeXZaTTaY6lUql9MZOBpmkAh2Xbsh4AOB2xwkK1qlcXq49VVb1ZKBQgiiJ4nm8zAA/Hxsc/ZrOSp6rqfLPZ9MIwXALwCsDrMAxvb29tnSiqsiBJkicIwicAj7j//qu/BgCMz/3oOPduNwAAAABJRU5ErkJggg==)](https://github.com/medium-isp)

Cross-platform daemon for managing hosts file.

The main idea was simplifying access to Medium DNS without tinkering with DNS servers but with updating hosts file. After seconds idea appeared: make it configurable and able to use not only Medium DNS repository for obtaining domain names.

What it do:

* Fetching Medium DNS JSON
* Comparing with entries in hosts file
* Adds, updates or deleting domains in hosts file

Planned:

* More than one repository for domain names
* API for getting statistics or for configuring hostsd
* hostsdctl for unprivileged API access
* Launch scripts for all popular init systems (systemd, launchd, openrc, etc.)

## Installation

For now you can do:

```
go install -u -v github.com/medium-isp/hostsd/cmd/hostsd
```

and move file to ``/usr/local/bin``, for example.

## Configuration

Nothing should be configured by default. See [documentation](/doc/configuration.md) about list of available environment variables and expected values.
