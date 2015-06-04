# mailsender
simple mail-sending service, http-api, written in golang


## install
install mail-sender from src
```bash
	# get src
	mkdir -p $GOPATH/src/github.com/niean && cd $GOPATH/src/github.com/niean
	git clone git@github.com:niean/mailsender.git && cd mailsender
	
	# build
	./control build
	
	# config
	mv cfg.example.json cfg.json
	vim cfg.json
	
	# start
	./control start
	...
	
	# stop
	./control stop
```

install mail-sender from bin
```bash
	# get bin
	mkdir -p /home/to/mailsender && cd /home/to/mailsender
	wget https://github.com/niean/mailsender/tree/master/bin/tycs-mailsender.tar.gz
	tar -zxf tycs-mailsender.tar.gz
	
	# config
	mv cfg.example.json cfg.json
	vim cfg.json
	
	# start
	./control start
	...
	
	# stop
	./control stop
```

## usage
```bash
	# test, send one mail to tycloudstart.com:1986
	curl -X POST -d "tos=anddynie@gmail.com;subject=some_subject;content=some_content;user=anddy" "tycloudstart.com:1986/mail/sender"

    # http-api
    POST /mail/sender HTTP/1.1
    Host host:port
    Args
	   -- tos: receivers, separated by comma(,)
	   -- subject: subject of ur mail
	   -- content: content of ur mail
       -- user: optional, displayed name of sender at receiver's side
```

## configuration
```
    debug: true/false, whether or not open debug-logging

    http
        - enable: true/false, whether or not enable http-server 
        - listen: int, port of the http-server

   	mail 
        - enable: true/false, whether or not enable mail-sending service
        - sendConcurrent: int, max number of concurrent threads sending mails
        - maxQueueSize: int, max number of cached mails which are to be sent
        - fromUser: string, default display name of sender at receiver's side
        - mailServerHost: string, hostname of the smtp (or pop3) server 
        - mailServerPort: int, port of the mailServerHost
        - mailServerAccount: string, authorized user account on mailServerHost
        - mailServerPasswd: int, password for mailServerAccount
```

## reference
TODO
