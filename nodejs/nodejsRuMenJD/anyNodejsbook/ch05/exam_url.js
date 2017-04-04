var http = require('http'),
    url = require('url');

http.createServer(function(req,res) {
    var pathname = url.parse(req.url).pathname;

    if (pathname === '/') {
	res.writeHead(200,
	    
	    
	
