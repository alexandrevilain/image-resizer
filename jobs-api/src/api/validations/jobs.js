'use_strict';

module.exports = Joi => {
  return {
    create: {
      payload: {
        width: Joi.number().required(),
        height: Joi.number(),
        file: Joi.object()
      }
    }
  };
};
