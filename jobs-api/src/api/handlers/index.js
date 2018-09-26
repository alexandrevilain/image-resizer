'use_strict';

module.exports = services => {
  return {
    jobs: require('./jobs')(services)
  };
};
