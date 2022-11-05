<script setup>
  import Forecast from './Forecast.vue'

</script>

<template>
  <div id="app">
    <h1>Local Weather Forecast</h1>
  </div>
  <ul>
    <li v-for="f of forecast" :key="f.id">
      {{ f }}
    </li>
  </ul>
  <main>
    <Forecast />
  </main>
</template>
<script>
</script>
<style>
#app {
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
li {
  list-style: none;
}
</style>
