var app = angular.module('pt_management', ['ngRoute']);

var global = {
    url: 'http://0.0.0.0:9000',
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
    .when('/home', {
        templateUrl:'./html_components/dashboard.html',
        controller:'mainController',
        title:'Dashboard',
    })
    .when('/', {
        templateUrl:'./html_components/dashboard.html',
        controller:'mainController',
        title:'Dashboard',
    })
    .when('/events', {
        templateUrl:'./html_components/events.html',
        controller:'eventsController',
        title:'Events',
    })
    .when('/addEvent', {
        templateUrl:'./html_components/addEvent.html',
        controller:'eventsController',
        title:'Events',
    })
    .when('/login', {
        templateUrl:'./html_components/login.html',
        controller:'primaryController',
        title:'Login',
    })
    .when('/signUp', {
        templateUrl:'./html_components/signup.html',
        controller:'primaryController',
        title:'Login',
    })
    .when('/profile', {
        templateUrl:'./html_components/profile.html',
        controller:'profileController',
        title:'Login',
    })
})

app.factory('eventsStore', function () {
	let usageArray = [];
	let updateUsageStore = function (x) {
        console.log('received as');
        console.log(x);
		for (let i in x) {
            usageArray.push(x[i]);
        }
	};
	let getUsageStore = function () {
		return usageArray;
	};
	return {
		updateEventsStore: updateUsageStore,
		getEventsStore   : getUsageStore,
	};
});

app.controller('primaryController', function($scope,$location,$rootScope,$http) {
    console.warn('primaryController called')
    $rootScope.showSidebar = false;
    $rootScope.settingsOption = false;
    $scope.refreshStop = global.refresh;
    $scope.patientid = '';
    $scope.handleSignUp = function() {
        let data = 'name=' + $scope.signup.name + '&email=' + $scope.signup.email + '&age=' + $scope.signup.age + '&dob=' + $scope.signup.dob + '&bg=' + $scope.signup.bg + '&b=' + $scope.signup.b;
        console.log('data is', data);
        $http(
            {url: global.url+'/signup',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        }).then(resp => {
            let res = resp.data;
            console.log('resp is ', res)
            console.log(resp)
            if (res) {
                $scope.patientid = res;
                $rootScope.showSidebar = true;
                eventsStore.updateEventsStore(res);
            } else {
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });
    }

    $scope.handleLogin = function() {
        let data = 'patientid=' + $scope.patientid;
        console.log('data is', data);
        $http(
            {url: global.url+'/login',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        }).then(resp => {
            let res = resp.data;
            console.log('res is ', res)
            if (true) {
                $location.path('/home');
                global.patientid = patientid;
                $rootScope.showSidebar = true;
                $rootScope.settingsOption = true;
            } else {
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });
    }
});

app.controller('mainController', function($scope,$location,$rootScope,$http) {
    console.warn('assignee controller called')
    $rootScope.showSidebar = true;
    $rootScope.settingsOption = true;
    $scope.refreshStop = global.refresh;
    $scope.eventsArr = [];
    $scope.getAllEventsPatient = function() {
        let data = 'patientid='+global.patientid;
        console.warn('fetching')
        $http(
            {url: global.url+'/allevents',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        }).then(resp => {
            res = resp.data;
            console.warn('thisss')
            console.warn(res)
            if (res) {
                $rootScope.showSidebar = true;
                $scope.eventsArr = res;
                // eventsStore.updateEventsStore(res);
            } else {
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });

    };
});

app.controller('profileController', function($scope,$location,$rootScope,$http) {
    console.warn('profile controller called')
    $rootScope.showSidebar = true;
    $rootScope.settingsOption = true;
    $scope.refreshStop = global.refresh;
    $scope.eventsArr = [];
    $scope.getProfile = function() {
        let data = 'patientid='+global.patientid;
        console.warn('fetching')
        $http(
            {url: global.url+'/get_profile',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        }).then(resp => {
            res = resp.data;
            console.warn('get profiles ...')
            console.warn(res)
            if (res) {
                $rootScope.showSidebar = true;
                // eventsStore.updateEventsStore(res);
                $scope.wrongpass = 'Updated successfully'
            } else {
                $scope.wrongpass = 'Error occurred';
            }
        });

    };

    $scope.updateProfile = function() {
        let data = 'asugar='+$scope.profile.asugar + '&bsugar=' + $scope.profile.bsugar + '&sbp=' + $scope.profile.sbp + '&dbp=' + $scope.profile.dbp + '&bmi=' + $scope.profile.bmi + '&temp='
         + $scope.profile.temp + '&pulse=' + $scope.profile.pulse + '&resp=' + $scope.profile.resp + '&gender=' + $scope.profile.gender + '&patientid=' + global.patientid;
        console.warn('fetching')
        $http(
            {url: global.url+'/update_profile',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        }).then(resp => {
            res = resp.data;
            console.warn('update profiles ...')
            console.warn(res)
            if (res) {
                $rootScope.showSidebar = true;
                // eventsStore.updateEventsStore(res);
            } else {
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });

    };
});

app.controller('eventsController', function($scope,$location,$rootScope,$http) {
    console.warn('events controller called')
    $rootScope.showSidebar = true;
    $rootScope.settingsOption = true;
    $scope.eventsArr = [];
    $scope.refreshStop = global.refresh;
    $scope.getAllEventsPatient = function() {
        let data = 'patientid='+global.patientid;
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
            console.warn(res)
            if(res){
                $scope.wrongpass = 'Success';
                $rootScope.showSidebar = true;
                $scope.eventsArr = res;
            } else{
                $scope.wrongpass = 'Error occurred while Adding assignee';
            }
        });
    };

    $scope.addEvent = function() {
        $location.path('/addEvent');
    };

    $scope.addEventParticular = function() {
        // format date
        console.warn('date is ', $scope.event.date);
        let date = $scope.event.date.toString();
        let rawArr = date.split(' ');
        let monthNumber = s => {
            switch (s) {
                case 'Jan':
                    return '01';
                case 'Feb':
                    return '02';
                case 'Mar':
                    return '03';
                case 'Apr':
                    return '04';
                case 'May':
                    return '05';
                case 'Jun':
                    return '06';
                case 'Jul':
                    return '07';
                case 'Aug':
                    return '08';
                case 'Sep':
                    return '09';
                case 'Oct':
                    return '10';
                case 'Nov':
                    return '11';
                case 'Dec':
                    return '12';
            }
        };
        date = rawArr[2] + '/' + monthNumber(rawArr[1]) + '/' + rawArr[3];
        let data = 'patientid=' + global.patientid + '&disease=' + $scope.event.disease + '&medicine=' + $scope.event.medicine + '&time=' + $scope.event.time + '&date=' + date;
        console.warn('data is');
        console.warn(data);
        $http({
            url: global.url+'/add_data',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data:data
        })
        .then(resp => {
            res = resp.data;
            if (res) {
                $scope.wrongpass = 'Success';
                $rootScope.showSidebar = true;
                $location.path('/events');
            } else {
                $scope.wrongpass = 'Error occurred while Adding Events';
            }
        });
    }

    $scope.getSuggestions = function() {
        let data = 'event=' + $scope.event.disease + '&patientid=' + global.patientid;
        $http({
            url: global.url+'/suggestmedicines',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data: data
        })
        .then(resp => {
            let res = resp.data;
            console.warn('get suggestions ...')
            console.warn(res)
            // if (res.length) {
            //     let min = {
            //         medicine: res[0].Medicine,
            //         prob: res[0].Probability
            //     };
            //     for(let x in res) {
            //         if (min[prob])
            //     }
            //     $rootScope.showSidebar = true;
            //     $location.path('/events');
            // } else {
            //     $scope.wrongpass = 'Error occurred while Adding Events';
            // }
        });
    }
});