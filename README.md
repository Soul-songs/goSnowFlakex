## goSnowFlake

[![Build Status](https://travis-ci.org/zheng-ji/goSnowFlake.svg)](https://travis-ci.org/zheng-ji/goSnowFlake)
[![GoDoc](https://godoc.org/github.com/zheng-ji/goSnowFlake?status.svg)](https://godoc.org/github.com/zheng-ji/goSnowFlake)

According to the Twitter SnowFlake Theory, A ThreadSafe Unique ID  Generator written by Golang

根据 Twitter SnowFlake 算法, 实现的分布式线程安全 UID 生成器

![goSnowFlake](https://github.com/zheng-ji/goSnowFlake/blob/master/logo/snowflake.png)

Feature
--------

* ThreadSafe unique id generator
* Green pluggable, without external storage like Redis or MySQL
* Suitable for distributed systems
* Implement Twitter's SnowFlake theory


Description
-----------


```
0               41	           51			64
+---------------+----------------+-----------+
|timestamp(ms)  | worker node id | sequence	 |
+---------------+----------------+-----------+

id  = timestamp | workerid | sequence (eg. 1451063443347648410)

```

An id is composed by three part: timestamp in millon second, worker id, and sequence. Sequence is zero default. when timestamp is the same, we use sequence to avoid conflict



Installation
-------------

```
go get github.com/zheng-ji/goSnowFlake
```

Example
-------

```go
import (
	"fmt"
	"github.com/zheng-ji/goSnowFlake"
)

func main() {
    // Params: Given the workerId, 0 < workerId < 1024
	iw, err := goSnowFlake.NewIdWorker(1) 
	if err!= nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		if id, err := iw.NextId(); err != nil {
            fmt.Println(err)
        } else{
            fmt.Println(id)
        }
	}
}
```

Documentation
-------------

- [Twitter Blog Reference](https://blog.twitter.com/2010/announcing-snowflake)
- [Reddit Discuss](https://www.reddit.com/comments/cajap/twitter_announces_snowflake_a_distributed_unique/)

License
-------

Copyright (c) 2016 by [zheng-ji](http://zheng-ji.info) released under MIT License.

