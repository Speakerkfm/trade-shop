const redis = require('redis'),
    params = require('./params');

let client = redis.createClient(params.redis);

client.on('connect', function () {
    console.log("Redis client connected");
});

module.exports = client;