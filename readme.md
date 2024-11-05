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
You could use the `./test.sh` in the repo to make the request and usage:

```
./test.sh <PAYLOAD_SIZE_IN_BYTES>
```

and follow the logs

5. Some examples

```
./test.sh 1000 -> ok
./test.sh 2000 -> ok
./test.sh 10000 -> ok
./test.sh 50000 -> ok
./test.sh 100000 -> ok
./test.sh 1000000 -> NOT OK, got duplicated payload and received 2000000 from the upstream
```

