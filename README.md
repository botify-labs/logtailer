Logtailer : multi-hosts "tail -f" for your logs
===============================================

**NB: this repo is in very early alpha, use at your own risks!**

For instance it doesn't clean subprocesses porperly for now when hitting
Ctrl+C.


Installation
------------

Until there are releases, you'd better use `go run` or `go build` yourself:
```
go run logtailer.log <server(s)> <file(s)>
```

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

References
----------

This is my first program in Go, so I'll leave here some articles or
StackOverflow questions that helped me making this. Thanks to their respective
authors!

Useful references:
- running a command: http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
- streaming logs + ANSI colors: http://kvz.io/blog/2013/07/12/prefix-streaming-stdout-and-stderr-in-golang/
- time formatting: http://stackoverflow.com/questions/5885486/how-do-i-get-the-current-timestamp-in-go
- signal handling: http://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
- waiting goroutines with a wait group: http://stackoverflow.com/questions/18405023/how-would-you-define-a-pool-of-goroutines-to-be-executed-at-once-in-golang
