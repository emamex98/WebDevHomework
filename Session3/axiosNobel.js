import axios from 'axios';

let url = "http://api.nobelprize.org/v1/laureate.json?bornCountryCode=mx";

axios.get(url)
    .then(function (response) {
        var respJson = response.data["laureates"]
        for (var i = 0; i < respJson.length; i++) {
            if (respJson[i]["surname"] == "Paz") {
                console.log(respJson[i]["surname"])
            }
        }
    })
    .catch(function (error) {
        console.log(error);
    });
