const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    expect = require('chai').expect;

async function postLogin(email, password, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .json()
        .post(state.host() + '/login')
        .send({
            email: email,
            password: password
        })
        .expectStatus(status)
        .expect(expect)
        .end();
}

describe('POST /login', () => {
    before(async () => {
        await state.new();
    });

    it('login ok', async () => {
        await postLogin(
            "speaker123@mail.ru",
            "123456",
            function (res, body, next) {
                expect(res.headers['set-cookie']).not.equal("");
                next();
            },
            302
        );
    });

    it('login wrong password', async () => {
        await postLogin(
            "speaker123@mail.ru",
            "12345678",
            function (res, body, next) {
                expect(body.error.description).equal('Wrong email or password');
                next();
            },
            401
        );
    });

    it('login wrong username', async () => {
        await postLogin(
            "speakerzxcv123@mail.ru",
            "12345678",
            function (res, body, next) {
                expect(body.error.description).to.be.equal('Wrong email or password');
                next();
            },
            401
        );
    })
});