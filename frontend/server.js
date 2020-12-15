const express = require('express');
const app = express();
app.use(express.static('dist/frontend'));
app.get('/', function (req, res,next) {
  res.redirect('/');
});
app.listen(8080)
