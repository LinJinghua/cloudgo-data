# cloudgo-data

## 概述
详见[golang 构建数据服务](http://blog.csdn.net/pmlpml/article/details/78602290)

## 内容

使用 ``xorm`` 或 ``gorm`` 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 ``database/sql`` 与 ``orm`` 实现的异同！
- 编程效率：orm架构编程效率更高。
- 程序结构：orm架构的程序更优。
- 服务性能：sql更好。

### ``orm`` 是否就是实现了 ``dao`` 的自动化？
``orm`` 在一定程度上确实是实现了 ``dao`` 的自动化。实际上并不彻底。

### 使用 ``ab`` 测试性能
##### POST
- ``database/sql``

```plain
$ ab -n 1000 -c 100 -p data.txt -T "application/x-www-form-urlencoded" "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        120 bytes

Concurrency Level:      100
Time taken for tests:   5.215 seconds
Complete requests:      1000
Failed requests:        982
   (Connect: 0, Receive: 0, Length: 982, Exceptions: 0)
Total transferred:      245803 bytes
Total body sent:        205000
HTML transferred:       121803 bytes
Requests per second:    191.74 [#/sec] (mean)
Time per request:       521.549 [ms] (mean)
Time per request:       5.215 [ms] (mean, across all concurrent requests)
Transfer rate:          46.02 [Kbytes/sec] received
                        38.38 kb/s sent
                        84.41 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:   117  513 147.9    544     863
Waiting:      116  512 147.8    543     863
Total:        117  513 147.9    544     863

Percentage of the requests served within a certain time (ms)
  50%    544
  66%    593
  75%    631
  80%    646
  90%    694
  95%    728
  98%    753
  99%    771
 100%    863 (longest request)
```

- ``orm``

```plain
ab -n 1000 -c 100 -p data.txt -T "application/x-www-form-urlencoded" "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        89 bytes

Concurrency Level:      100
Time taken for tests:   7.494 seconds
Complete requests:      1000
Failed requests:        994
   (Connect: 0, Receive: 0, Length: 994, Exceptions: 0)
Total transferred:      213902 bytes
Total body sent:        205000
HTML transferred:       90902 bytes
Requests per second:    133.44 [#/sec] (mean)
Time per request:       749.377 [ms] (mean)
Time per request:       7.494 [ms] (mean, across all concurrent requests)
Transfer rate:          27.87 [Kbytes/sec] received
                        26.71 kb/s sent
                        54.59 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:   105  730 259.3    754    1333
Waiting:      105  729 259.3    753    1332
Total:        105  730 259.3    754    1333

Percentage of the requests served within a certain time (ms)
  50%    754
  66%    856
  75%    908
  80%    951
  90%   1065
  95%   1159
  98%   1233
  99%   1261
 100%   1333 (longest request)
```

##### GET

- ``database/sql``

```plain
$ ab -n 1000 -c 100 "http://127.0.0.1:8080/service/userinfo"
ab -n 1000 -c 100 "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        111 bytes

Concurrency Level:      100
Time taken for tests:   3.531 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      235000 bytes
HTML transferred:       111000 bytes
Requests per second:    283.20 [#/sec] (mean)
Time per request:       353.110 [ms] (mean)
Time per request:       3.531 [ms] (mean, across all concurrent requests)
Transfer rate:          64.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       4
Processing:     2  340 236.6    370    1039
Waiting:        2  339 236.5    369    1039
Total:          2  340 236.6    370    1040

Percentage of the requests served within a certain time (ms)
  50%    370
  66%    430
  75%    493
  80%    535
  90%    642
  95%    736
  98%    874
  99%    963
 100%   1040 (longest request)
```

- ``orm``

```plain
$ ab -n 1000 -c 100 "http://127.0.0.1:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        127247 bytes

Concurrency Level:      100
Time taken for tests:   8.720 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      127350000 bytes
HTML transferred:       127247000 bytes
Requests per second:    114.68 [#/sec] (mean)
Time per request:       871.986 [ms] (mean)
Time per request:       8.720 [ms] (mean, across all concurrent requests)
Transfer rate:          14262.30 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       1
Processing:   121  857 404.9    815    2143
Waiting:       31  641 382.8    687    1637
Total:        121  857 404.9    816    2143

Percentage of the requests served within a certain time (ms)
  50%    816
  66%   1066
  75%   1178
  80%   1225
  90%   1409
  95%   1599
  98%   1705
  99%   1741
 100%   2143 (longest request)
```

### 性能测试评价

在此处测试看来，``database/sql`` 和 ``orm`` 在处理请求的 web 服务上并没有对性能造成太大的影响，可能是这次测试数据量不大或者测试查询样例太简单以致orm优化很好。
