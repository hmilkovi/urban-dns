# Urban-DNS

Fast DNS on for large scale that filters ad domains, phishing domains, malware domains, supports dynamic DNS records and supports integration with route53.

Database mysql banchmarks (100 queries 10000 users):
```
$ dnstrace -n 100 -c 1000 --server 127.0.0.1 --recurse wener.test

Total requests:	 100000 of 100000 (100.0%)
DNS success codes:	100000

DNS response codes
	NOERROR:	100000

Time taken for tests:	 15.089083207s
Questions per second:	 6627.3

DNS timings, 100000 datapoints
	 min:		 0s
	 mean:		 144.613691ms
	 [+/-sd]:	 127.832583ms
	 max:		 1.476395007s

```


Health check endpoint:
```
http://127.0.0.1:8080/health
```

Metrics:
```
http://127.0.0.1:9153/metrics
```