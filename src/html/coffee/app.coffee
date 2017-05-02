app = angular.module "myApp",
    ['ui.router',
     "ngResource",
#     "ui.bootstrap"
    ]

app.config ($stateProvider, $urlRouterProvider)->
    $urlRouterProvider.otherwise("/")
    $stateProvider
    .state('index', {
        url: "/",
        templateUrl: "templates/index.html"
        controller: "indexCtrl"
    })
    .state('edit', {
        url: "/edit",
        templateUrl: "templates/edit.html"
        controller: "indexCtrl"
    })
    .state('register', {
        url: "/register",
        templateUrl: "templates/register.html"
        controller: "registerCtrl"
    })