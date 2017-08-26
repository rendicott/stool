angular.module('gapi', [])
    .controller('SampleController', function($scope, $http) {

        $scope.loadData = function() {
            console.log("Loading data...")
            $scope.getGames();
            $scope.getPlayers();
            $scope.getOutcomes();
        }

        $scope.getGames = function() {
            console.log("In getGames function");

            $http.get("/games")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding game data to scope");
                        console.log(response.data.data)
                        $scope.allGames = response.data.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get game data...");
                    }
                );
        };

        $scope.getPlayers = function() {
            console.log("In getPlayers function");

            $http.get("/players")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding player data to scope");
                        $scope.allPlayers = response.data.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get player data...");
                    }
                );
        };

        $scope.getOutcomes = function() {
            console.log("In getOutcomes function");

            $http.get("/outcomes")
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding outcome data to scope");
                        $scope.allOutcomes = response.data.data;
                    },
                    function errorCallback(response) {
                        console.log("Unable to get outcome data...");
                    }
                );
        };

        $scope.createGame = function(name) {
            console.log("In createGame function");

            $http.post("/games", {"name": name})
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding game data to scope");
                        $scope.allGames.push(response.data);
                    },
                    function errorCallback(response) {
                        console.log("Unable to create game: ");
                        console.log(response);
                    }
                );

        };

        $scope.createPlayer = function(name) {
            console.log("In createPlayer function");

            $http.post("/players", {"name": name})
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding player data to scope");
                        $scope.allPlayers.push(response.data);
                    },
                    function errorCallback(response) {
                        console.log("Unable to create player: ");
                        console.log(response);
                    }
                );

        };

        $scope.createOutcome = function(gameId, playerId, result, score, date) {
            console.log("In createOutcome function");


            $http.post("/outcomes", {"gameid": gameId, "playerid": playerId, "result": result, "score": score, "date": date})
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("adding outcome data to scope");
                        $scope.allOutcomes.push(response.data);
                    },
                    function errorCallback(response) {
                        console.log("Unable to create outcome: ");
                        console.log(response);
                    }
                );

        };

        $scope.deleteGame = function(id) {
            console.log("In deleteGame function");

            var endpoint = "/games/" + id;
            console.log("Deleting: " + endpoint);
            $http.delete(endpoint)
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("refreshing game data");
                        $scope.getGames();
                    },
                    function errorCallback(response) {
                        console.log("Unable to delete game: ");
                        console.log(response);
                    }
                );

        };

        $scope.deletePlayer = function(id) {
            console.log("In deletePlayer function");

            var endpoint = "/players/" + id;
            console.log("Deleting: " + endpoint);
            $http.delete(endpoint)
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("refreshing player data");
                        $scope.getPlayers();
                    },
                    function errorCallback(response) {
                        console.log("Unable to delete player: ");
                        console.log(response);
                    }
                );

        };

        $scope.deleteOutcome = function(id) {
            console.log("In deleteOutcome function");

            var endpoint = "/outcomes/" + id;
            console.log("Deleting: " + endpoint);
            $http.delete(endpoint)
                .then(
                    function successCallback(response) {
                        console.log(response.data);
                        console.log("refreshing outcome data");
                        $scope.getOutcomes();
                    },
                    function errorCallback(response) {
                        console.log("Unable to delete outcome: ");
                        console.log(response);
                    }
                );

        };


    });