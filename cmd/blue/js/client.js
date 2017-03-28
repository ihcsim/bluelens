// This module exports functions that give access to the bluelens API hosted at localhost.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'https';
    host = host || 'localhost';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // createMusic calls the create action of the music resource.
  // path is the request path, the format is "/bluelens/music"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createMusic = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // createUser calls the create action of the user resource.
  // path is the request path, the format is "/bluelens/user"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createUser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Update a user's followees list with a new followee.
  // path is the request path, the format is "/bluelens/user/:id/follows/:followeeID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.followUser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List up to N music resources. N can be adjusted using the 'limit' and 'offset' parameters.
  // path is the request path, the format is "/bluelens/music"
  // limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listMusic = function (path, limit, offset, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List up to N user resources. N can be adjusted using the 'limit' and 'offset' parameters.
  // path is the request path, the format is "/bluelens/user"
  // limit, offset are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listUser = function (path, limit, offset, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        limit: limit,
        offset: offset
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Add a music to a user's history.
  // path is the request path, the format is "/bluelens/user/:id/listen/:musicID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listenUser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Make music recommendations for a user.
  // path is the request path, the format is "/bluelens/recommendations/:userID/:limit"
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

  // Get a music resource with the given ID
  // path is the request path, the format is "/bluelens/music/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showMusic = function (path, config) {
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

  // Get a user resource with the given ID.
  // path is the request path, the format is "/bluelens/user/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showUser = function (path, config) {
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
