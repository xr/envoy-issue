var express = require('express')
var app = express()

app.post('/', function (req, res) {
    let payloadSize = 0;
    req.on('data', (chunk) => {
      payloadSize += chunk.length;
      console.log(`received chunk length: ${chunk.length}, total: ${payloadSize}`);
    });

    req.on('end', () => {
      console.log(`received total payload size: ${payloadSize}`);
      res.send({
        headers: req.headers,
        query: req.query,
        params: req.params,
        url: req.url,
        body: req.body,
        hostname: req.hostname,
        ip: req.ip,
        ips: req.ips,
        method: req.method,
        baseUrl: req.baseUrl,
        originalUrl: req.originalUrl,
        protocol: req.protocol,
        cookies: req.cookies,
        payloadSize: payloadSize,
      });
    });

})


app.get('/', function (req, res) {
  res.send({
    headers: req.headers,
    query: req.query,
    params: req.params,
    url: req.url,
    hostname: req.hostname,
    ip: req.ip,
    ips: req.ips,
    method: req.method,
    baseUrl: req.baseUrl,
    originalUrl: req.originalUrl,
    protocol: req.protocol,
    cookies: req.cookies,
  });
})
app.listen(4000);

