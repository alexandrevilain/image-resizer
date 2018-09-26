'use strict';

const defaultConfig = {
	description: 'no notes',
	notes: 'no description',
	tags: ['api']
};

module.exports = (server, handlers, validations) => {
	require('./routes')(server, handlers, validations, defaultConfig);
};
