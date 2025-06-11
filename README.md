# portsplz
!["a terminal window runs a quick nmap scan to check the open/closed status of every possible TCP port, and saves the output to a file. it then runs portsplz against the file, which outputs a list of ports are marked as open in that file. then, it runs a script and version scan against the same host, and passes the portsplz output to nmap so that only open ports are enumerated.](img/example.png)

## install
```
make install
```

## usage
```
portsplz <nmap-output-file>
```
only supports nmap -oN file
