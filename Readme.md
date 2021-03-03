# Dummy CDROM Ejector
Ok, This is one of the stupid idea that cross my mind. a service to remotely eject your system cdrom :D  
Obviously this is totally useless, it's just for fun, this just work on linux.

### Compile And Run
For building this, you need to install golang `apt install golang` or checkout [official website](https://golang.org)
```
# git clone https://github.com/sina-ghaderi/romdum
# cd romdum
# go build
# ./romdum --net 0.0.0.0:9940 --dev /dev/cdrom
info: application is listening on address 0.0.0.0:9940
...
```

Next connect the server using netcat a.k.a `nc`, may you want to install it? `apt install netcat`

```
# nc 127.0.0.1 9940
help banner for this terminal: system cdrom --> type 1 to eject and 0 to rollin 
1
operation completed successfully, check the cdrom.
0
operation completed successfully, check the cdrom.
... 
```

  
### Help And Licence
Copyright (c) 2021 blg.snix.ir, All rights reserved.
Developed BY sina@snix.ir --> FYI: this is just for fun and totally useless :D
This work is licensed under the terms of the MIT license.

```
# ./dumrom -h
usage of dummy cdrom-dvdrom ejector server:
./dumrom --net [ipv4:port] --dev <rom block device>
options:
	--net string     tcp network to listen on <ipv4:port> (default "127.0.0.1:9940")
	--dev string	 rom block device to mess with. (default is /dev/cdrom)
	--h              print this banner and exit
example: 
	./dumrom --net 0.0.0.0:9940 --dev /dev/cdrom

Copyright (c) 2021 blg.snix.ir, All rights reserved.
Developed BY sina@snix.ir --> FYI: this is just for fun and totally useless :D
This work is licensed under the terms of the MIT license.
```
