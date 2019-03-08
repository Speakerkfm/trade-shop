const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
    mysql = require('../../services/mysql'),
    expect = require('chai').expect;

async function getBuy(buffer, saleID, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() +'/sales/{sale_id}/buy')
        .pathParams({
            sale_id: saleID
        })
        .expectStatus(status)
        .expect(expect)
        .end()
}

describe("GET /sales/{sale_id}/buy", () => {
    describe("buy ok", () => {
        before(async () => {
            await state.new();
        });

        step("get buy", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await getBuy(
                buffer,
                "332d3e06-c05a-429b-892b-2f6ce064ee34",
                function (res, body, next) {
                    expect(body).to.equal(null);
                    next();
                },
                200
            )
        });

        step("check buyer inventory and bill in db", async () => {
            const result1 = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv1 = result1[0];
            expect(inv1.count).to.equal(12);

            const result2 = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'ad663316-7a59-48f2-8b15-e4f99573ebc0'"
            );
            const inv2 = result2[0];
            expect(inv2.count).to.equal(1);

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(338.5)
        });

        step("check seller bill in db", async () => {

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '56a8b65c-7432-463f-aadc-0e5b7eca8d06'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(261.5)
        });
    });

    describe("buy bad not enough money", () => {
        before(async () => {
            await state.new();
        });

        step("get buy without money", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await getBuy(
                buffer,
                "59f1640e-b4f4-41dd-879f-17392a54419b",
                function (res, body, next) {
                    expect(body.error.description).to.equal("Not enough money");
                    next();
                },
                400
            )
        });

        step("check buyer inventory and bill in db", async () => {
            const result1 = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv1 = result1[0];
            expect(inv1.count).to.equal(9);

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(500.0)
        });

        step("check seller bill in db", async () => {

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '56a8b65c-7432-463f-aadc-0e5b7eca8d06'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(100.0)
        });
    });

    describe("buy bad this is your lot", () => {
        before(async () => {
            await state.new();
        });

        step("get buy buy lot owner", async () => {
            let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

            await getBuy(
                buffer,
                "030e212f-99d7-4329-a69f-f7613d30a499",
                function (res, body, next) {
                    expect(body.error.description).to.equal("You can't buy your lot");
                    next();
                },
                400
            )
        });

        step("check buyer inventory and bill in db", async () => {
            const result1 = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv1 = result1[0];
            expect(inv1.count).to.equal(9);

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(500.0)
        });

        step("check seller bill in db", async () => {

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '56a8b65c-7432-463f-aadc-0e5b7eca8d06'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(100.0)
        });
    });

    describe("buy bad without auth", () => {
        before(async () => {
            await state.new();
        });

        step("get buy without auth", async () => {
            await getBuy(
                {},
                "59f1640e-b4f4-41dd-879f-17392a54419b",
                function (res, body, next) {
                    next();
                },
                401
            )
        });

        step("check buyer inventory and bill in db", async () => {
            const result1 = await mysql.query(
                "SELECT * FROM `trade-shop`.inventory WHERE user_id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe' AND  item_id = 'a5630737-636c-454c-826f-3002aaf46376'"
            );
            const inv1 = result1[0];
            expect(inv1.count).to.equal(9);

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '017d4ff8-e2c8-42f2-89f3-7822eeca3ebe'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(500.0)
        });

        step("check seller bill in db", async () => {

            const result3 = await mysql.query(
                "SELECT * FROM `trade-shop`.users WHERE id = '56a8b65c-7432-463f-aadc-0e5b7eca8d06'"
            );
            const user = result3[0];
            expect(user.bill).to.equal(100.0)
        });
    });
});