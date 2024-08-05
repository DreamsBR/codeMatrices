const jwt = require('jsonwebtoken');

const jwtSecret = 'mySecretKey';

exports.generateJWT = function() {
  return jwt.sign({}, jwtSecret, { expiresIn: '24h' });
};
