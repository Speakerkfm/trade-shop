const params = require('./params'),
    mysql = require('./mysql'),
    fs = require('fs'),
    yaml = require('yamljs');

module.exports = {
    loadTable: async table => {
        await mysql.query(`TRUNCATE TABLE ${table}`);

        const items = yaml.load(`./fixtures/${params.mysql.database}/${table}.yaml`)

        for (let item of items) {
            await mysql.query(`INSERT INTO ${table} SET ?`, item);
        }
    },
    dumpTable: async table => {
        const results = await mysql.query(`SELECT * FROM trade-shop.${table};`);
        console.log(`Table ${table} dumped`);
    }
};