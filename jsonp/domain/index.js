const http = require('http');

http.createServer((req,res)=>{
  res.writeHead(200, { 
    'Content-Type': 'text/javascript',
    // 'Content-Type': 'application/json',
    // 'Access-Control-Allow-Origin': 'http://localhost:5500'
    // 'Access-Control-Allow-Origin': '*'
  });
  res.end(`{
    name: 'peng',
    age: '22',
  }`);
}).listen(80,()=>{
  console.log('success listen 80');
})