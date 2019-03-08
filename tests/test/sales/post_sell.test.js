const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
    mysql = require('../../services/mysql'),
    expect = require('chai').expect;

async function postSale(buffer, data, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .json()
        .header('Cookie', buffer.cookie)
        .post(state.host() +'/user/sale')
        .send(data)
        .expectStatus(status)
        .expect(expect)
        .end()
}

describe("POST /user/sale", () => {
    describe("sale ok", () => {
        before(async () => {
            await state.new();
        });

        step("post sale", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await postSale(
                buffer,
                [
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46376",
                        "count": 1,
                        "price": 10.2
                    },
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46376",
                        "count": 1,
                        "price": 3.4
                    }
                ],
                function (res, body, next) {
                    expect(body).to.equal(null);
                    next();
                },
                200
            )
        });

        step("check user inventory in db", async () => {
            const result = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv = result[0];
            expect(inv.count).to.equal(7)
        });
    });

    describe("sale bad without auth", () => {
        before(async () => {
            await state.new();
        });

        step("post sale without auth", async () => {
            await postSale(
                {},
                [],
                function (res, body, next) {
                    expect(body).to.equal(null);
                    next();
                },
                401
            )
        });

        step("check user inventory in db", async () => {
            const result = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv = result[0];
            expect(inv.count).to.equal(9)
        });
    });

    describe("sale bad with not enough items", () => {
        before(async () => {
            await state.new();
        });

        step("post sale", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await postSale(
                buffer,
                [
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46376",
                        "count": 100,
                        "price": 10.2
                    },
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46376",
                        "count": 1,
                        "price": 3.4
                    }
                ],
                function (res, body, next) {
                    expect(body.error.description).to.equal("Not enough items");
                    next();
                },
                400
            )
        });

        step("check user inventory in db", async () => {
            const result = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv = result[0];
            expect(inv.count).to.equal(9)
        });
    });

    describe("sale bad with wrong item id", () => {
        before(async () => {
            await state.new();
        });

        step("post sale", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await postSale(
                buffer,
                [
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46300",
                        "count": 1,
                        "price": 10.2
                    },
                    {
                        "id": "a5630737-636c-454c-826f-3002aaf46376",
                        "count": 1,
                        "price": 3.4
                    }
                ],
                function (res, body, next) {
                    expect(body.error.description).to.equal("Not enough items");
                    next();
                },
                400
            )
        });

        step("check user inventory in db", async () => {
            const result = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv = result[0];
            expect(inv.count).to.equal(9)
        });
    });
});