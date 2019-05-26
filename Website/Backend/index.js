const express = require('express');
/*
const session = require('express-session');
const cookies = require('cookie-parser');
const body = require('body-parser');
const config = require('./config.js');
const routes = require('./routes');
*/

let app = express();
app.set('views', __dirname + '/views');
app.set('view engine', 'jsx');
app.engine('jsx', require('express-react-views').createEngine());
