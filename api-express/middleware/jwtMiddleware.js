const { expressjwt: expressJwt } = require('express-jwt');

const jwtSecret = 'mySecretKey';

exports.authenticateJWT = expressJwt({
  secret: jwtSecret,
  algorithms: ['HS256']
});
