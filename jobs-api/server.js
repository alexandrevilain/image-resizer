'use strict';

require('./config');

const Hapi = require('hapi');
const Inert = require('inert');
const Vision = require('vision');

const init = async () => {
  try {
    const server = new Hapi.Server({
      host: '0.0.0.0',
      port: config.api.port,
      routes: {
        cors: {
          origin: ['*']
        },
        validate: {
          failAction: async (request, h, err) => {
            console.error(`Validation error: ${err.message}`);
            throw process.env.NODE_ENV === 'production'
              ? Boom.badRequest(`Invalid request payload input`)
              : err;
          }
        }
      }
    });
    await require('./src')(server);
    await server.register([Inert, Vision, require('hapi-log-requests')(console)]);
    await server.start();
    return server;
  } catch (err) {
    throw err;
  }
};

init()
  .then(server => {
    console.log('Server running at', server.info.uri);
  })
  .catch(err => {
    console.error(err);
    process.exit(1);
  });
