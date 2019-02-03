# CoreDNS database query

Supports sqlite3 and mysql

Insert test data:
```
insert into records(name,type,content,ttl,disabled,blocked)values("wener.pass","A","192.168.1.1",3600,0,1);
insert into records(name,type,content,ttl,disabled,blocked)values("wener.block","A","192.168.1.1",3600,0,0);
```

Test it:
```
nslookup wener.pass 127.0.0.1
nslookup wener.block 127.0.0.1
```