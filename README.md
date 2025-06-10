# portsplz

helper tool

## install
```
make install
```

## usage
```
portsplz <nmap-output-file>
```
only supports nmap -oN file

## intended use case
```
$ sudo nmap -Pn -T4 -v -p- <target_addr> -oN tcp-first.nmap
$ sudo nmap -Pn -v -sCV -p $(notesplz tcp-first.nmap)$ <target_addr> -oN tcp-second.nmap
```