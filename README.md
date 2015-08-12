Logtailer : multi-hosts "tail -f" for your logs
===============================================

**NB: this repo is in very early alpha, use at your own risks!**

For instance it doesn't clean subprocesses porperly for now when hitting
Ctrl+C.


Installation
------------

TODO

Usage
-----

Generic usage:
```
logtrailer <host1, [host2, host3, ...]> <file1, [file2, file3, ...]>
```

Given you can use shell expansion on your side, and shell globing for some
patterns on remote side!

Examples:
```
logtrailer elasticsearch{1,2,3}.example.net "/var/log/elasticsearch/*.log"
#=> will follow all logs inside this folder for those 3 machines

logtrailer server.example.net /var/log/syslog "/var/log/**/*.log"
#=> will follow basically all logs on server.example.net
```
