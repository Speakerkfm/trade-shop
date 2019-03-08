const amqp = require('amqplib');

const EXCHANGE_NAME = 'mailer';
const EXCHANGE_KEY = '';
const EXCHANGE_TYPE = 'direct';
const EXCHANGE_OPTION = {
    durable: true
};

module.exports = {
    getMessage: async () => {
        const conn = await amqp.connect(
            `amqp://`+ process.env.AMQP_USER+':'+process.env.AMQP_PASSWORD+'@'+process.env.AMQP_HOST+':'+process.env.AMQP_PORT+'/'
        );
        const channel = await conn.createChannel();

        await channel.assertExchange(EXCHANGE_NAME, EXCHANGE_TYPE, EXCHANGE_OPTION);

        const q = await channel.assertQueue(EXCHANGE_NAME, {
            name: EXCHANGE_NAME,
            durable: true,
            autoDelete: false,
            exclusive: false,
            nowait: false,
            arguments: {"x-dead-letter-exchange": "mailer_fail"}
        });

        console.log('Waiting for logs');

        await channel.bindQueue(q.queue, EXCHANGE_NAME, EXCHANGE_KEY);

        let message = await channel.get(q.queue, {noAck: false});

        channel.close();
        conn.close();

        return JSON.parse(message.content.toString());
    },

    clearMessages: async () => {
        const conn = await amqp.connect(
            `amqp://`+process.env.AMQP_USER+':'+process.env.AMQP_PASSWORD+'@'+process.env.AMQP_HOST+':'+process.env.AMQP_PORT
        );
        const channel = await conn.createChannel();
        await channel.purgeQueue(EXCHANGE_NAME);
    }
};