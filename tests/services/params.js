params = {
    mysqlFixtures: [
        'inventory',
        'item_sale',
        'items',
        'sales',
        'users'
    ],
    mysql: {
        connectionLimit: 10,
        host: process.env.DATABASE_HOST,
        user: process.env.DATABASE_USER,
        password: process.env.DATABASE_PASSWORD,
        database: process.env.DATABASE_NAME,
        port: process.env.DATABASE_PORT
    },
    redis: {
        host: process.env.REDIS_HOST.split(':')[0],
        port: process.env.REDIS_HOST.split(':')[1]
    },
    apiHost: 'http://localhost:8000'
};

module.exports = params;