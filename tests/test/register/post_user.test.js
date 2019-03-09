const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    expect = require('chai').expect;

async function postUser(email, password, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .json()
        .post(state.host() + '/user')
        .send({
            email: email,
            password: password
        })
        .expectStatus(status)
        .expect(expect)
        .end();
}

describe('POST /user', () => {
    before(async () => {
        await state.new();
    });

    it('register ok', async () => {
        await postUser(
            "speaker12345@mail.ru",
            "123456",
            function (res, body, next) {
                expect(body).equal(null);
                next();
            },
            200
        );
    });

    it('register with email taken', async () => {
        await postUser(
            "speaker123@mail.ru",
            "123456",
            function (res, body, next) {
                expect(body.error.description).equal('Email address is already taken');
                next();
            },
            400
        );
    });
});