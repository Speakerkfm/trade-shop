const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json');

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
        await state.loadFixtures();
    });

    it('login', async () => {
        await postLogin(
            "speaker123@mail.ru",
            "123456",
            function (res, body, next) {
               next();
            },
            302
        );
    })
});