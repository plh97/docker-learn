const http = require('http');
const url = require('url');

http.createServer((req,res)=>{
  res.writeHead(200, { 
    'Content-Type': 'application/json',
  });

  const query = url.parse(req.url,true).query;
  res.end(`${query.callback}(${JSON.stringify(query)})`);
}).listen(80,()=>{
  console.log('success listen 80');
})