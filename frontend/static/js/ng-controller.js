angular.module('gapi', [])
    .controller('SampleController', function($scope, $http) {


        $scope.angularTest = "yes";

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
                        $scope.allGames = response.data;
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
                        $scope.allPlayers = response.data;
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
                        $scope.allOutcomes = response.data;
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

        $scope.createOutcome = function(gameId, playerId, win) {
            console.log("In createGame function");

            if (win === "true") {
                var winBool = true;
            } else {
                var winBool = false;
            }

            $http.post("/outcomes", {"gameid": gameId, "playerid": playerId, "win": winBool})
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


    });