Commands to build and run the front end locally:
First, `cd` to the front end directory.
`docker build -t vue .`
`docker run -it -p 8081:8080 --rm --name vue-1 vue`

Command to build and run the back end locally:
First, `cd` to the back end directory.
`docker build -t forecast .`
`docker run -it --rm -p 8090:8090 -v $PWD/src:/go/src/forecast forecast`
