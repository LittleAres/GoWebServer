// Generated by CoffeeScript 1.4.0
(function() {

  angular.module("myApp").factory("index", function($resource) {
    return $resource('/j/angular/index/:id', null, {
      'query': {
        method: 'GET',
        isArray: false
      },
      'update': {
        method: 'PUT'
      }
    });
  });

}).call(this);
