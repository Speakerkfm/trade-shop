const hippie = require('hippie-swagger'),
    state = require('../../services/state'),
    swagger = require('../../../tmp/swagger.dereference.json'),
    auth = require('../../services/auth'),
    expect = require('chai').expect;

async function getExchangeRates(buffer, expect, status) {
    return hippie(swagger, state.swaggerOptions)
        .header('Cookie', buffer.cookie)
        .get(state.host() +'/exchange_rates')
        .expectStatus(status)
        .expect(expect)
        .end()
}

describe("GET /exchange_rates", () => {
    before(async () => {
        await state.new('proxy_data/get_exchange_rates');
    });

    step("exchange rates", async () => {
        let buffer = await auth.loginUser("speaker123@mail.ru", "123456");

        await getExchangeRates(
            buffer,
            function (res, body, next) {
                expect(body).to.eql(
                    {
                        "Date": "2019-03-09T11:30:00+03:00",
                        "PreviousDate": "2019-03-07T11:30:00+03:00",
                        "PreviousURL": "\/\/www.cbr-xml-daily.ru\/archive\/2019\/03\/07\/daily_json.js",
                        "Timestamp": "2019-03-10T23:00:00+03:00",
                        "Valute": {
                            "AUD": {
                                "ID": "R01010",
                                "NumCode": "036",
                                "CharCode": "AUD",
                                "Nominal": 1,
                                "Name": "Австралийский доллар",
                                "Value": 46.4457,
                                "Previous": 46.3074
                            },
                            "AZN": {
                                "ID": "R01020A",
                                "NumCode": "944",
                                "CharCode": "AZN",
                                "Nominal": 1,
                                "Name": "Азербайджанский манат",
                                "Value": 38.8828,
                                "Previous": 38.8111
                            },
                            "GBP": {
                                "ID": "R01035",
                                "NumCode": "826",
                                "CharCode": "GBP",
                                "Nominal": 1,
                                "Name": "Фунт стерлингов Соединенного королевства",
                                "Value": 86.8292,
                                "Previous": 86.5638
                            },
                            "AMD": {
                                "ID": "R01060",
                                "NumCode": "051",
                                "CharCode": "AMD",
                                "Nominal": 100,
                                "Name": "Армянских драмов",
                                "Value": 13.4704,
                                "Previous": 13.4483
                            },
                            "BYN": {
                                "ID": "R01090B",
                                "NumCode": "933",
                                "CharCode": "BYN",
                                "Nominal": 1,
                                "Name": "Белорусский рубль",
                                "Value": 30.8606,
                                "Previous": 30.762
                            },
                            "BGN": {
                                "ID": "R01100",
                                "NumCode": "975",
                                "CharCode": "BGN",
                                "Nominal": 1,
                                "Name": "Болгарский лев",
                                "Value": 38.1541,
                                "Previous": 38.0639
                            },
                            "BRL": {
                                "ID": "R01115",
                                "NumCode": "986",
                                "CharCode": "BRL",
                                "Nominal": 1,
                                "Name": "Бразильский реал",
                                "Value": 17.1769,
                                "Previous": 17.4395
                            },
                            "HUF": {
                                "ID": "R01135",
                                "NumCode": "348",
                                "CharCode": "HUF",
                                "Nominal": 100,
                                "Name": "Венгерских форинтов",
                                "Value": 23.6644,
                                "Previous": 23.5967
                            },
                            "HKD": {
                                "ID": "R01200",
                                "NumCode": "344",
                                "CharCode": "HKD",
                                "Nominal": 10,
                                "Name": "Гонконгских долларов",
                                "Value": 84.0335,
                                "Previous": 83.8775
                            },
                            "DKK": {
                                "ID": "R01215",
                                "NumCode": "208",
                                "CharCode": "DKK",
                                "Nominal": 1,
                                "Name": "Датская крона",
                                "Value": 10.0019,
                                "Previous": 99.7833
                            },
                            "USD": {
                                "ID": "R01235",
                                "NumCode": "840",
                                "CharCode": "USD",
                                "Nominal": 1,
                                "Name": "Доллар США",
                                "Value": 65.9646,
                                "Previous": 65.843
                            },
                            "EUR": {
                                "ID": "R01239",
                                "NumCode": "978",
                                "CharCode": "EUR",
                                "Nominal": 1,
                                "Name": "Евро",
                                "Value": 74.573,
                                "Previous": 74.4158
                            },
                            "INR": {
                                "ID": "R01270",
                                "NumCode": "356",
                                "CharCode": "INR",
                                "Nominal": 100,
                                "Name": "Индийских рупий",
                                "Value": 94.2217,
                                "Previous": 93.348
                            },
                            "KZT": {
                                "ID": "R01335",
                                "NumCode": "398",
                                "CharCode": "KZT",
                                "Nominal": 100,
                                "Name": "Казахстанских тенге",
                                "Value": 17.3829,
                                "Previous": 17.4174
                            },
                            "CAD": {
                                "ID": "R01350",
                                "NumCode": "124",
                                "CharCode": "CAD",
                                "Nominal": 1,
                                "Name": "Канадский доллар",
                                "Value": 49.0699,
                                "Previous": 49.2652
                            },
                            "KGS": {
                                "ID": "R01370",
                                "NumCode": "417",
                                "CharCode": "KGS",
                                "Nominal": 100,
                                "Name": "Киргизских сомов",
                                "Value": 94.6,
                                "Previous": 94.4121
                            },
                            "CNY": {
                                "ID": "R01375",
                                "NumCode": "156",
                                "CharCode": "CNY",
                                "Nominal": 10,
                                "Name": "Китайских юаней",
                                "Value": 98.3519,
                                "Previous": 98.1106
                            },
                            "MDL": {
                                "ID": "R01500",
                                "NumCode": "498",
                                "CharCode": "MDL",
                                "Nominal": 10,
                                "Name": "Молдавских леев",
                                "Value": 38.7003,
                                "Previous": 38.5667
                            },
                            "NOK": {
                                "ID": "R01535",
                                "NumCode": "578",
                                "CharCode": "NOK",
                                "Nominal": 10,
                                "Name": "Норвежских крон",
                                "Value": 76.096,
                                "Previous": 75.884
                            },
                            "PLN": {
                                "ID": "R01565",
                                "NumCode": "985",
                                "CharCode": "PLN",
                                "Nominal": 1,
                                "Name": "Польский злотый",
                                "Value": 17.3555,
                                "Previous": 17.3134
                            },
                            "RON": {
                                "ID": "R01585F",
                                "NumCode": "946",
                                "CharCode": "RON",
                                "Nominal": 1,
                                "Name": "Румынский лей",
                                "Value": 15.7235,
                                "Previous": 15.6814
                            },
                            "XDR": {
                                "ID": "R01589",
                                "NumCode": "960",
                                "CharCode": "XDR",
                                "Nominal": 1,
                                "Name": "СДР (специальные права заимствования)",
                                "Value": 91.7126,
                                "Previous": 91.662
                            },
                            "SGD": {
                                "ID": "R01625",
                                "NumCode": "702",
                                "CharCode": "SGD",
                                "Nominal": 1,
                                "Name": "Сингапурский доллар",
                                "Value": 48.6034,
                                "Previous": 48.5067
                            },
                            "TJS": {
                                "ID": "R01670",
                                "NumCode": "972",
                                "CharCode": "TJS",
                                "Nominal": 10,
                                "Name": "Таджикских сомони",
                                "Value": 69.897,
                                "Previous": 69.746
                            },
                            "TRY": {
                                "ID": "R01700J",
                                "NumCode": "949",
                                "CharCode": "TRY",
                                "Nominal": 1,
                                "Name": "Турецкая лира",
                                "Value": 12.1169,
                                "Previous": 12.223
                            },
                            "TMT": {
                                "ID": "R01710A",
                                "NumCode": "934",
                                "CharCode": "TMT",
                                "Nominal": 1,
                                "Name": "Новый туркменский манат",
                                "Value": 18.874,
                                "Previous": 18.8392
                            },
                            "UZS": {
                                "ID": "R01717",
                                "NumCode": "860",
                                "CharCode": "UZS",
                                "Nominal": 10000,
                                "Name": "Узбекских сумов",
                                "Value": 78.6186,
                                "Previous": 78.4504
                            },
                            "UAH": {
                                "ID": "R01720",
                                "NumCode": "980",
                                "CharCode": "UAH",
                                "Nominal": 10,
                                "Name": "Украинских гривен",
                                "Value": 24.9937,
                                "Previous": 24.8765
                            },
                            "CZK": {
                                "ID": "R01760",
                                "NumCode": "203",
                                "CharCode": "CZK",
                                "Nominal": 10,
                                "Name": "Чешских крон",
                                "Value": 29.1492,
                                "Previous": 29.0492
                            },
                            "SEK": {
                                "ID": "R01770",
                                "NumCode": "752",
                                "CharCode": "SEK",
                                "Nominal": 10,
                                "Name": "Шведских крон",
                                "Value": 70.9038,
                                "Previous": 70.6303
                            },
                            "CHF": {
                                "ID": "R01775",
                                "NumCode": "756",
                                "CharCode": "CHF",
                                "Nominal": 1,
                                "Name": "Швейцарский франк",
                                "Value": 65.656,
                                "Previous": 65.5546
                            },
                            "ZAR": {
                                "ID": "R01810",
                                "NumCode": "710",
                                "CharCode": "ZAR",
                                "Nominal": 10,
                                "Name": "Южноафриканских рэндов",
                                "Value": 46.1837,
                                "Previous": 46.3944
                            },
                            "KRW": {
                                "ID": "R01815",
                                "NumCode": "410",
                                "CharCode": "KRW",
                                "Nominal": 1000,
                                "Name": "Вон Республики Корея",
                                "Value": 58.4443,
                                "Previous": 58.3177
                            },
                            "JPY": {
                                "ID": "R01820",
                                "NumCode": "392",
                                "CharCode": "JPY",
                                "Nominal": 100,
                                "Name": "Японских иен",
                                "Value": 59.0314,
                                "Previous": 58.8778
                            }
                        }
                    }
                );
                next();
            },
            200
        )
    });

    step("exchange rates without auth", async () => {
        await getExchangeRates(
            {},
            function (res, body, next) {
                expect(body).to.equal(null);
                next();
            },
            401
        )
    })
});