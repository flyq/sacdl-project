var http = require('http');
http.createServer(function (req, res) {
  res.end('Hello World\n');
}).listen(8888, "127.0.0.1");
console.log('Server running at http://127.0.0.1:3000/');
