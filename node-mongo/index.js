const express = require('express');
const mongoose = require('mongoose');
const bodyparser = require('body-parser');

const Item = require('./models/Item.model')

const app = express();

app.use(bodyparser())


mongoose
.connect('mongodb://mongo:27017/docker-node-mongo', {useNewUrlParser: true})
.then((e)=>{
  console.log('mongoose connect success');
})
.catch(err=>{
  console.log(err);
});

app.get('/', function (req, res) {
  Item
    .find()
    .then(item=>res.render('index', {item}))
    .catch(err=>res.status(400).json({
      msg:'no item find.的的22人'
    }))
});



app.post('/item/add', function (req, res) {
  const newItem = new Item({
    name: res.body.name
  })
  newItem.save().then(item=>res.redirect('/'))
})

const PORT = 3000;

app.listen(PORT, ()=>{
  console.log(`listen port:${PORT}`);
})