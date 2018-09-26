'use_strict';

const uuid = require('uuid');
const path = require('path');

module.exports = ({ StorageService, NatsService }) => {
  return {
    create(request, h) {
      const { width, height, file } = request.payload;
      const id = uuid.v4();
      const filename = `${id}${path.extname(file.hapi.filename)}`;
      return StorageService.uploadFile(filename, file)
        .then(url => ({
          id,
          width,
          height,
          originalUrl: url
        }))
        .then(job => {
          NatsService.publish(job);
          return job;
        });
    },
    get(request, h) {}
  };
};
