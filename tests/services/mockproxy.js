const request = promisifyRequest(require('request'));

module.exports = {
    load: async path => {
        return request(
            `http://localhost:8080/_mockproxy?dir=tests/${path}`,
            (error, response, body) => {
                if (error != null) {
                    console.log('mockproxy error ', error);
                }
            }
        )
    },
    setDelay: async sec => {
        return request(
            `http://localhost:8080/_mockproxy/runtime?delayOnce=${sec}`,
            (error, response, body) => {
                if (error != null) {
                    console.log('mockproxy error ', error);
                }
            }
        )
    }
};

function promisifyRequest(req) {
    return function (options) {
        let args = [options];
        return new Promise((resolve, reject) => {
            args.push((e, response, body) =>
            e ? reject({error: e, response, body}) : resolve({response, body})
            );
            req.apply(null, args);
        });
    };
}