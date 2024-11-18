## Readme
This repo is used to reproduce the issue reported [here](https://github.com/envoyproxy/envoy/issues/36119)


## Requirements
- nodejs
- golang
- envoy


## Run
1. start the upstream server:
```
cd echo-server && node index.js
```

2. start the ext-proc grpc server:
```
cd ext-proc && go run main.go
```

3. start envoy:
```
envoy -f envoycfg.yaml
```

4. Run the request

You will need to install `autocannon` to perform the load testing
```
autocannon -c 2 -m GET -d 60 -R 10 --renderStatusCodes http://localhost:10000/
```

this will mostly be fine, there is no cancelled context, however with a bit more traffics like:


```
autocannon -c 2 -m GET -d 60 -R 30 --renderStatusCodes http://localhost:10000/
```

We will get quite a lot of:
```
Error reading message from stream: rpc error: code = Canceled desc = context canceled
```
