const express = require('express')
const path = require('path')
const PORT = process.env.PORT || 5000

var app = express()

app.use('/dist', express.static(path.join(__dirname, '../dist')))

app.get('*', function (request, response) {
  response.sendFile(path.resolve(__dirname, '../index.html'));
});

app.listen(PORT, () => console.log(`Listening on ${ PORT }`))
