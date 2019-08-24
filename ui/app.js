var app = angular.module('pt_management', ['ngRoute']);

var global = {
    url: 'http://0.0.0.0:5000',
    username: 'default'
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
    $scope.addAssignee = function() {
        let data = 'username='+$scope.assignee_form.username+'&password='+$scope.assignee_form.password
            +'&name='+$scope.assignee_form.name+'&master='+global.username+'&task='+$scope.assignee_form.task
        $http(
            {url:global.url+'/assigneeAdd',
            method:'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data}
        )
        .then(resp=>{
            res=resp.data;
            if(res['Success']=='Y'){
                $scope.wrongpass = 'Success';
                $rootScope.showSidebar = true;
                setTimeout($location.path('/assignees'),2000)
            }
            else{
                $scope.wrongpass = 'Error occurred while Adding assignee'
            }
        })

    }
    $scope.retriveAssignees = function(){
        $scope.showLoading = true;
        $http(
            {url:global.url+'/assignee',
            method:'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:'master='+global.username
        }
        )
        .then(resp=>{
            res=resp.data;
            if(res['Success']=='Y'){
                $scope.showLoading = false;
                $scope.ass = res['result'];
            }
            else{
                console.error('error occurred while requesting for asignee list')
                $scope.showLoading = false;
            }
        })
    }
    $scope.removeAss = function(username) {
        $http(
            {url:global.url+'/delAssignee',
            method:'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:'master='+global.username+'&username='+username
        }
        )
        .then(resp=>{
            res=resp.data;
            if(res['Success']=='Y'){
                alert('Removed Assignee '+username+'. Refresh to see the effect.');
            }
            else{
                alert('Error Removing Assignee '+username);
            }
        })
    } 
})