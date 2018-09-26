'use strict';

// As we use config everywhere in the app, make it global!
global.config = require('common-env/withLogger')(console).getOrElseAll({
  api: {
    port: 8000
  },
  storage: {
    bucketName: 'uploads',
    server: 'localhost',
    port: 9000,
    accessKey: 'supinfo',
    secretKey: 'supinfo1234',
    ssl: false
  },
  nats: {
    connectionStrings: 'nats://localhost:4222',
    queue: 'jobs'
  }
});
