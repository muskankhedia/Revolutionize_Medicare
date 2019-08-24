var app = angular.module('pt_management', ['ngRoute']);

var global = {
    url: 'http://0.0.0.0:5000',
    username: 'default',
    patientid: '1'
};

app.config(function($routeProvider,$locationProvider) {
    $routeProvider
    .when("/login/",{
        templateUrl:'./html_components/welcome.html',
        controller:'loginController',
        title:'Login | SignUp',
    })
    .when("/login",{
        templateUrl:'./html_components/login.html',
        controller:'loginController',
        title:'Login | SignUp',
    })
    .when('/dashboard', {
        templateUrl:'./html_components/dashboard.html',
        controller:'mainController',
        title:'Dashboard',
    })
    .when('/', {
        templateUrl:'./html_components/dashboard.html',
        controller:'mainController',
        title:'Dashboard',
    })
})


app.controller('mainController', function($scope,$location,$rootScope,$http) {
    console.warn('assignee controller called')
    $rootScope.showSidebar = true;
    $rootScope.settingsOption = true;
    $scope.refreshStop = global.refresh;
    $scope.getAlEventsPatient = function() {
        let data = 'username='+$scope.assignee_form.username+'&patientid='+global.patientid;
        $http(
            {url: global.url+'/allevents',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        })
        .then(resp=>{
            res=resp.data;
            if(res){
                $scope.wrongpass = 'Success';
                $rootScope.showSidebar = true;
            }
            else{
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });

    };
})