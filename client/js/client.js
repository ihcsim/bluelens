// This module exports functions that give access to the bluelens API hosted at localhost:8080.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'localhost';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // Update a user's followees list with a new followee.
  // path is the request path, the format is "/bluelens/user/:id/follows/:followeeID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.followUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Get a music resource with the given ID
  // path is the request path, the format is "/bluelens/music/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.getMusic = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Get a user resource with the given ID
  // path is the request path, the format is "/bluelens/user/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.getUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Add a music to a user's history.
  // path is the request path, the format is "/bluelens/user/:id/listen/:musicID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listenUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Make music recommendations for a user.
  // path is the request path, the format is "/bluelens/recommendations/:userID/:maxCount"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.recommendRecommendations = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }
  return client;
  };
});
