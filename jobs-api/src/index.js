'use strict';

module.exports = server => {
  return require('./services')().then(services => {
    return require('./api')(server, services);
  });
};
