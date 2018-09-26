'use strict';

const Joi = require('joi');

module.exports = () => {
  return {
    jobs: require('./jobs')(Joi)
  };
};
