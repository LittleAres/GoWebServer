app = angular.module("myApp")

app.controller "baseCtrl", ($scope, $rootScope)->
  $rootScope.title = "Index"
  $rootScope.small = ""
  $rootScope.items =[]
  $rootScope.active_item = "Index"