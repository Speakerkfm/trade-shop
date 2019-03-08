const fixtures = require('./fixtures'),
    mockproxy = require('./mockproxy'),
    params = require('./params'),
    redis = require('./redis');

module.exports = {
  new: async function(proxyUrl) {
      let promises = [this.loadFixtures(), this.redisFlushAll()];
      if (proxyUrl !== undefined) {
          promises.push(mockproxy.load(proxyUrl));
      }
      await Promise.all(promises);
  },
  loadFixtures: async () => {
      let promises = [];

      for (let table of params.mysqlFixtures) {
          promises.push(fixtures.loadTable(table));
      }

      await Promise.all(promises)
  },
  redisFlushAll: async () => {
      await redis.flushall();
  },
  redisSet: async (key, value) => {
      await redis.set(key, value, redis.print);
  },
  host: () => {
      return params.apiHost;
  },
  swaggerOptions: {
      validateResponseSchema: true,
      validateParameterSchema: true,
      errorOnExtraHeaderParameters: true,
      errorOnExtraParameters: false
  }
};