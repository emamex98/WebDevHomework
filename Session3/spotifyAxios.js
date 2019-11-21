import axios from 'axios';

let url = "https://api.spotify.com/v1/artists/00FQb4jTyendYWaN8pK0wa/top-tracks?country=mx";
let config = {
        headers: {
        'Authorization': 'Bearer BQAmQO6C6q0XSlZqMCVi7Vl8zTGwKFfpYArSa2x_Cr0C0g8K-tyYsoRXzeWrudoRy3u1_SphNChI5eNPMKrtHLpGeyqUWeqmapqsjhVG7chX0OKILHkNNg1O0xASBZidvgiVdmPCUL5eCJc'
        }
    };

axios.get(url, config)
    .then(function (response) {
         var respJson = response["data"]["tracks"];
         for (var i = 0; i < respJson.length; i++) {
            console.log(respJson[i]["name"])
         }
    })
    .catch(function (error) {
        console.log(error);
    });
