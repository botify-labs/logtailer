logtailer : multi-hosts "tail -f"
=================================

**NB: this repo is in very early alpha, use at your own risks!**

For instance it doesn't clean subprocesses porperly for now when hitting
Ctrl+C.


Installation
------------

Go download the correct version of the tool on the [releases page](https://github.com/botify-labs/logtailer/releases).

If you better like command-line, here we go:

... for Linux 64bits users:
```
sudo curl -L https://github.com/botify-labs/logtailer/releases/download/0.0.1/logtailer_linux-amd64 -o /usr/local/bin/logtailer
sudo chmod +x /usr/local/bin/logtailer
```

... for Mac OSX users:
```
sudo curl -L https://github.com/botify-labs/logtailer/releases/download/0.0.1/logtailer_darwin-amd64 -o /usr/local/bin/logtailer
sudo chmod +x /usr/local/bin/logtailer
```

If you want to build a *development* version, use `go run` yourself:
```
go run logtailer.go <server(s)> <file(s)>
```


Usage
-----

Generic usage:
```
logtrailer [-n150] <host1 [host2 host3 ...]> <file1 [file2 file3 ...]>
```

Given you can use shell expansion on your side, and shell globing for some
patterns on remote side!

Examples:
```
logtrailer elasticsearch{1,2,3}.example.net "/var/log/elasticsearch/*.log"
#=> will follow all logs inside this folder for those 3 machines

logtrailer -n25 server.example.net /var/log/syslog "/var/log/**/*.log"
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
- signal handling: http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in
- terminating processes: http://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
- waiting goroutines with a wait group: http://stackoverflow.com/questions/18405023/how-would-you-define-a-pool-of-goroutines-to-be-executed-at-once-in-golang
