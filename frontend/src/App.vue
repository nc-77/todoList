<template>
  <div id="app">
    <home />
  </div>
</template>

<script>
import Home from "./views/Home.vue";
export default {
  name: "App",
  components: {
    Home,
  },
};
</script>

// 引入初始样式
<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap");
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
body {
  font-family: "Poppins", sans-serif;
}
</style>
