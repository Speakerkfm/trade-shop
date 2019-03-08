let SwaggerParser = require('swagger-parser'),
    parser = new SwaggerParser(),
    path = require('path');

parser.dereference(path.join(__dirname, '../../tmp/swagger.yaml'), function (err, api) {
    if (err) throw err;
    console.log(JSON.stringify(api));
});