const hippie = require('hippie-swagger'),
    state = require('./state'),
    swagger = require('../../tmp/swagger.dereference.json');

let cookie;

module.exports = {
    loginUser: async function postLogin(email, password) {
        const buffer = {};

        await hippie(swagger, state.swaggerOptions)
            .json()
            .post(state.host() + '/login')
            .send({
                email: email,
                password: password
            })
            .expect(function(res, body, next) {
                if ('set-cookie' in res.headers) {
                    cookie = res.headers['set-cookie'][0];
                }
                buffer.cookie = cookie;
                next();
            })
            .end();

        return buffer
    }
};