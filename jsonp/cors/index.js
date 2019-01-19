const http = require('http');

http.createServer((req,res)=>{
  res.writeHead(200, { 
    'Content-Type': 'text/plain',
    'Access-Control-Allow-Origin': 'http://localhost:5500'
    // 'Access-Control-Allow-Origin': '*'
  });
  res.end('nodejs run in docker');
}).listen(80,()=>{
  console.log('success listen 80');
})