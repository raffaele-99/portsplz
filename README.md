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

## example of intended use case
1. quick scan all ports and output in normal format
```
$ sudo nmap -Pn -T4 -v -p- localhost -oN tcp-first.nmap
```
```
$ cat tcp-first.nmap
# Nmap 7.95 scan initiated Wed Jun 11 09:56:48 2025 as: nmap -Pn -T4 -p- -oN tcp-first.nmap localhost
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00011s latency).
Other addresses for localhost (not scanned): ::1
Not shown: 65524 closed tcp ports (reset)
PORT      STATE    SERVICE
5000/tcp  open     upnp
7000/tcp  open     afs3-fileserver
59318/tcp open     unknown
```
2. use tool to run deeper scan on open ports 
```
$ sudo nmap -Pn -v -sCV -p $(portsplz tcp-first.nmap) localhost -oN tcp-second.nmap
```

```
$ cat tcp-second.nmap
# Nmap 7.95 scan initiated Wed Jun 11 10:00:22 2025 as: nmap -Pn -v -sCV -p T:5000,7000,57621,58949,59318 -oN tcp-second.nmap localhost
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00011s latency).
Other addresses for localhost (not scanned): ::1

PORT      STATE SERVICE        VERSION
5000/tcp  open  rtsp
|_rtsp-methods: ERROR: Script execution failed (use -d to debug)
| fingerprint-strings:
|   FourOhFourRequest:
            ...SNIPPET...
|   GetRequest:
            ...SNIPPET...
|   HTTPOptions:
            ...SNIPPET...
|   RTSPRequest:
            ...SNIPPET...
|   SIPOptions:
            ...SNIPPET...
7000/tcp  open  rtsp
|_irc-info: Unable to open connection
|_rtsp-methods: ERROR: Script execution failed (use -d to debug)
| fingerprint-strings:
|   FourOhFourRequest:
|     HTTP/1.1 403 Forbidden
            ...SNIPPET...
|   GetRequest:
            ...SNIPPET...
|   HTTPOptions:
            ...SNIPPET...
|   RTSPRequest:
            ...SNIPPET...
|   SIPOptions:
            ...SNIPPET...

1 services unrecognized despite returning data. If you know the service/version, please submit the following fingerprints at https://nmap.org/cgi-bin/submit.cgi?new-service :
            ...SNIPPET...
Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
# Nmap done at Wed Jun 11 10:01:04 2025 -- 1 IP address (1 host up) scanned in 41.77 seconds
```