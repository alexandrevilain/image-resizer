'use_strict';

module.exports = (server, services) => {
  const handlers = require('./handlers')(services);
  const validations = require('./validations')();
  return {
    handlers,
    validations,
    routes: require('./routes')(server, handlers, validations)
  };
};
