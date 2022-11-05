<template>
    <div id="app">
      <input v-model="location1" type="text" /><br />
      <button @click="getForecast()">Get Forecast</button>
    </div>
    <ul>
      <li v-for="f of forecast" :key="f.id">
        {{ f }}
      </li>
    </ul>
  </template>
  <script>
  import axios from "axios";
  export default {
    name: "App",
    data() {
      return {
        forecast: ""
      };
    },

    methods: {
      async getForecast() {
        try {
          const getUrl = `http://0.0.0.0:8090/forecast?freetext=` + this.location1
          const res = await axios.get(getUrl);
          this.forecast = res.data;
        } catch (error) {
          console.log(error);
    }

      },
    },
  };
  </script>