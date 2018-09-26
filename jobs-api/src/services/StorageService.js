'use strict';

const minio = require('minio');

const { server, port, accessKey, secretKey, ssl, bucketName } = config.storage;

const minioClient = new minio.Client({
  endPoint: server,
  port: port,
  useSSL: ssl,
  accessKey,
  secretKey
});

const policy = `
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": { "AWS": ["*"] },
      "Action": ["s3:ListBucketMultipartUploads", "s3:GetBucketLocation", "s3:ListBucket"],
      "Resource": ["arn:aws:s3:::${bucketName}"]
    },
    {
      "Effect": "Allow",
      "Principal": { "AWS": ["*"] },
      "Action": [
        "s3:ListMultipartUploadParts",
        "s3:PutObject",
        "s3:AbortMultipartUpload",
        "s3:DeleteObject",
        "s3:GetObject"
      ],
      "Resource": ["arn:aws:s3:::${bucketName}/*"]
    }
  ]
}`;
const getPublicUrl = filename => {
  const protocol = ssl ? 'https' : 'http';
  return `${protocol}://${server}:${port}/${bucketName}/${filename}`;
};

module.exports = () => {
  return minioClient
    .bucketExists(bucketName)
    .then(exists => {
      return exists
        ? Promise.resolve()
        : minioClient
            .makeBucket(bucketName)
            .then(() => minioClient.setBucketPolicy(bucketName, policy));
    })
    .then(() => {
      return {
        uploadFile: (name, fileStream) => {
          return minioClient.putObject(bucketName, name, fileStream).then(() => getPublicUrl(name));
        }
      };
    });
};
