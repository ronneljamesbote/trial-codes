<template>
  <div class="app">
    <Header/>
    <Content>
      <h1 class="app-title">Famous Persons</h1>
      <div class="cards-container">
        <Card v-for="person in persons" :key="person.id" v-bind="{ person }"/>
      </div>
    </Content>
    <Footer/>
  </div>
</template>

<script>
import personList from "./data/persons";
import Card from "./Card.vue";
import Content from "./layout/Content.vue";
import Footer from "./layout/Footer.vue";
import Header from "./layout/Header.vue";

export default {
  name: "app",
  components: {
    Card,
    Content,
    Footer,
    Header
  },
  data() {
    return {
      persons: []
    };
  },
  created() {
    this.persons = personList;
  }
};
</script>

<style>
html {
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  color: #2c3e50;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.app {
  background: rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.app .app-title {
  margin-bottom: 1.5em;
  text-align: center;
}

.app .cards-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
}

/* For big screen sizes */
@media screen and (min-width: 769px) {
  .app .app-title {
    text-align: left;
  }

  .app .cards-container {
    justify-content: space-between;
  }
}
</style>
