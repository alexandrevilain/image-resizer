'use strict';

module.exports = () => {
  return require('./StorageService')().then(StorageService => ({
    NatsService: require('./NatsService')(),
    StorageService
  }));
};
