const http = require('http');

const server = http.createServer((req,res)=>{
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('nodejs run in docker');
});

const PORT = 80;

server.listen(PORT,()=>{
  console.log(`success run port ${PORT}`);
})