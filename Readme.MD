## simple-syslog

A small goapp that uses gopkg.in/mcuadros/go-syslog.v2 to listen for syslog data and print it out to the screen.  A good app for small scale tests of syslog forwarding agents. 


### installation

```
go install github.com/randomtask1155/simple-syslog
```

### usage

```
PORT=514 ./simple-syslog 
```
