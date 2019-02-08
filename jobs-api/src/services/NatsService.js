'use strict';

const nats = require('nats');
const { connectionStrings, queue } = config.nats;
const connection = nats.connect({ servers: connectionStrings.split(',') });

module.exports = () => {
  return {
    publish: job => {
      return connection.publish(queue, JSON.stringify(job));
    }
  };
};
