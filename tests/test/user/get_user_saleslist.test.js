const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
expect = require('chai').expect;

async function getSales(buffer, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() + '/user/sales')
        .expectStatus(status)
        .expect(expect)
        .end();
}

describe('GET /user/sales', () => {
    before(async () => {
        await state.new();
    });

    it('sales list ok', async () => {
        let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

        await getSales(
            buffer,
            function (res, body, next) {
                expect(body).to.be.deep.eql(
                    [
                        {
                            "id": "030e212f-99d7-4329-a69f-f7613d30a499",
                            "items": [
                                {
                                    "count": 4,
                                    "id": "ad663316-7a59-48f2-8b15-e4f99573ebc0",
                                    "name": "item2",
                                    "price": 15.9
                                }
                            ],
                            "total_count": 63.6
                        }
                    ]
                );
                next();
            },
            200
        );
    });

    step("sales list without auth", async () => {
        await getSales(
            {},
            function (res, body, next) {
                expect(body).to.equal(null);
                next();
            },
            401
        )
    })
});