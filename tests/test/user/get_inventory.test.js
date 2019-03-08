const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
    expect = require('chai').expect;

async function getInventory(buffer, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() +'/user/inventory')
        .expectStatus(status)
        .expect(expect)
        .end()
}

describe("GET /user/inventory", () => {
    describe("inventory ok", () => {
        before(async () => {
            await state.new();
        });

        step("get user inventory", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await getInventory(
                buffer,
                function (res, body, next) {
                    expect(body).to.eql(
                        {
                            "bill": "500.00",
                            "items": [
                                {
                                    "count": 9,
                                    "id": "a5630737-636c-454c-826f-3002aaf46376",
                                    "name": "item1"
                                }
                            ]
                        }
                    );
                    next();
                },
                200
            )
        });

        step("check redis", async () => {
           await state.redisAssert("rate_inventory_017d4ff8-e2c8-42f2-89f3-7822eeca3ebe", "[{\"UserID\":\"017d4ff8-e2c8-42f2-89f3-7822eeca3ebe\",\"ItemID\":\"a5630737-636c-454c-826f-3002aaf46376\",\"Name\":\"item1\",\"Count\":9}]");
        });
    });

    describe("inventory bad", () => {
        before(async () => {
            await state.new();
        });

        step("get user inventory without auth", async () => {
            await getInventory(
                {},
                function (res, body, next) {
                    expect(body).to.equal(null);
                    next();
                },
                401
            )
        });

        step("check redis", async () => {
            await state.redisAssert("rate_inventory_017d4ff8-e2c8-42f2-89f3-7822eeca3ebe", null);
        });
    });
});