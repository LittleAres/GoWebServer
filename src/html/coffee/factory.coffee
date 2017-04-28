angular.module("myApp")
  .factory "index", ($resource)->
    $resource('/j/angular/index/:id', null,{
        'query': { method:'GET', isArray: false },
        'update': {method: 'PUT'}
    })
