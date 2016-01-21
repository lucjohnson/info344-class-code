// tsd install express --save will provide parameter hinting for Express (do within Vagrant ssh /vagrant)
// tsd install morgan --save wil provide the same for morgan library
// download Postman chrome extension
// sudo npm install -g nodemon, instead of node server.js use nodemon -L server.js (this will make it easier than restarting the server every time, may want to turn off auto save)
// npm install <package name> --save --no-bin-links when working on Vagrant on a Windows machine
// sudo npm config set bin-links false is a command that could also be used

'use strict';

// require the express module
var express = require('express');
var morgan = require('morgan');
var bodyParser = require('body-parser');

// create a new express application
var app = express();

// log requests
app.use(morgan('dev'));

// parse JSON post bodies
app.use(bodyParser.json());

// serve static files from /static (replaces function below)
app.use(express.static(__dirname + '/static'));

// call this function for GET on root resource (/)
// app.get('/', function(req, res) {
// 	res.setHeader('Content-Type', 'text/plain');
// 	res.send('Hello World!');
// });

// call this function for GET /time
app.get('/time', function(req, res) {
	res.setHeader('Content-Type', 'text/plain');
	res.send(new Date());
});

app.get('/api/v1/users', function(req, res) {
	var users = [
		{
			email: 'blah@blah.com',
			displayName: 'Blah user'
		}
	];
	res.json(users);
});

app.post('/api/v1/users', function(req, res) {
	console.log(req.body);
	
	res.json({message: 'new user created'});
});

// listen for HTTP requests on port 80
app.listen(80, function() {
	console.log('server is listening');
});
