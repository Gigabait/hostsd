# hostsd configuration

There is two type of configuration:

* Static configuration that is read at startup from environment variables.
* Dynamic configuration that can be modified with hostsdctl.

## Static configuration

Here goes a table of environment variables that is read at startup:

| Variable | Expected value | Description |
| -------- | -------------- | ----------- |
| DATADIR | Path to directory | Directory with dynamic configuration, caches, and so on. Defaulting to ``/var/lib/hostsd``. |
| DOWNLOADER_DONOTSTART | boolean | Start (``false``) or not (``true``) remote lists downloader. Defaulting to ``false``. |
| HOSTS_PATH | path to hosts file | Custom hosts file to use. Useful for development. |
| HTTP_DONOTSTART | boolean | Start (``false``) or not (``true``) internal HTTP server which is needed for ``hostsdctl``. Defaulting to ``false``. |
| HTTP_LISTEN | ip_address:port | IP address and port to start HTTP server listening on. Defaulting to ``127.0.0.1:61525``. |
| HTTP_WAITFORSECONDS | integer | Maximum number of seconds to wait for HTTP server to reply on startup. If HTTP server won't start - hostsd will exit. Defaulting to 10 seconds. |
| PARSER_FILES | List of URLs to fetch delimited with comma | List of URLs to fetch. These URLs should contain JSON file in format like [here](https://raw.githubusercontent.com/medium-isp/medium-dns/master/hosts/hosts.json). URLs should be encoded. |

## Dynamic configuration

As dynamic configuration might be changed between releases very often - please issue ``hostsdctl config help``. This will show a list of available variables, their current values and description.
