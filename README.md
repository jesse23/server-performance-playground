# server-performance-playground
Server Performance Playground


## Performance test
```sh
wrk2 -t2 -c30 -d30s -R60 -L http://localhost:8000
wrk2 -t2 -c10 -d30s -R60 --timeout 500 http://localhost:8000
```

# Reference
https://juejin.cn/post/6844904095409504270