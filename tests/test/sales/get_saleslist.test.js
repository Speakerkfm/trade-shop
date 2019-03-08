const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth')
    expect = require('chai').expect;

async function getSales(buffer, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() + '/sales')
        .expectStatus(status)
        .expect(expect)
        .end();
}

describe('GET /sales', () => {
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
                            "id": "332d3e06-c05a-429b-892b-2f6ce064ee34",
                            "items": [
                                {
                                    "count": 3,
                                    "id": "a5630737-636c-454c-826f-3002aaf46376",
                                    "name": "item1",
                                    "price": 50.5
                                },
                                {
                                    "count": 1,
                                    "id": "ad663316-7a59-48f2-8b15-e4f99573ebc0",
                                    "name": "item2",
                                    "price": 10
                                }
                            ],
                            "total_count": 161.5
                        },
                        {
                            "id": "59f1640e-b4f4-41dd-879f-17392a54419b",
                            "items": [
                                {
                                    "count": 30,
                                    "id": "a5630737-636c-454c-826f-3002aaf46376",
                                    "name": "item1",
                                    "price": 30.3
                                }
                            ],
                            "total_count": 909
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