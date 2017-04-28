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
        template:"<h1>靖康耻犹未雪，臣子恨何时灭，<br>驾长车踏破贺兰山缺。壮士饥餐胡虏肉，<br>笑谈渴饮匈奴血。待从头，收拾旧山河，朝天阙。</h1>"
#        templateUrl: "src/html/templates/index.html"
#        controller: "indexCtrl"
    })
    .state('edit', {
        url: "/edit",
        templateUrl: "templates/edit.html"
        controller: "indexCtrl"
    })