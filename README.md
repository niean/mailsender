# mailsender
simple mail-sending service, http-api, written in golang


## install
```
	git clone git@github.com:niean/mailsender.git
```

## usage
```
    # install mailserver
	git clone git@github.com:niean/mailsender.git && cd mailsender
	mv cfg.example.json cfg.json //then change cfg.json as needed
	./control start
	./control tail

	# test, assuming http server listens on "tycloudstart.com:1986"
	curl -X POST -d "tos=nieanan@xiaomi.com;subject=ur subject;content=ur content;user=from_user_name" "tycloudstart.com:1986/mail/sender"

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