This repo is a weather app composed of a dockererized front end in Vue.js and a dockerized backend written in GoLang.  
The backend takes a plaintext input from the user via the frontend page, makes an api to convert that location to a latitude/longitude pair, then makes another api call to convert that pair to a grid point, then makes a final api call to `Weather.gov` to get the forecast.  Which is then displayed on the page.

Commands to build and run the front end locally:
First, `cd` to the front end directory.
`docker build -t vue .`
`docker run -it -p 8081:8080 --rm --name vue-1 vue`

Command to build and run the back end locally:
First, `cd` to the back end directory.
`docker build -t forecast .`
`docker run -it --rm -p 8090:8090 -v $PWD/src:/go/src/forecast forecast`
