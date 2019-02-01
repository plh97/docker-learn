const axios = require('axios');


axios({
  method: 'options',
  url: 'http://localhost/'
}).then(res=>{
  console.log(res.data);
})




