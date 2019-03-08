const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
    expect = require('chai').expect;

async function getLogout(buffer, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() +'/user/logout')
        .expectStatus(status)
        .expect(expect)
        .end()
}

describe("GET /user/logout", () => {
    before(async () => {
        await state.new();
    });

    step("user inventory", async () => {
        let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

        await getLogout(
            buffer,
            function (res, body, next) {
                expect(body).to.equal(null);
                next();
            },
            200
        )
    });

    step("logout without auth", async () => {
        await getLogout(
            {},
            function (res, body, next) {
                expect(body).to.equal(null);
                next();
            },
            401
        )
    })
});