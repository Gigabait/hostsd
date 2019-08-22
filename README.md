# hostsd

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

Nothing should be configured by default. See [config struct](/internal/configuration/config.go) for available options. They should be specified as environment variables, like ``HOSTS_PATH``.