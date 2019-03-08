const util = require('util'),
    pool = require('mysql'),
    params = require('./params');

const mysql = pool.createPool(params.mysql);

//замена коллбэков на промисы
mysql.query = util.promisify(mysql.query);

module.exports = mysql;