'use_strict';

module.exports = (server, handlers, validations, defaultConfig) => {
  server.route({
    method: 'POST',
    path: '/upload',
    config: {
      ...defaultConfig,
      handler: handlers.jobs.create,
      validate: validations.jobs.create,
      payload: {
        output: 'stream'
      }
    }
  });
};
