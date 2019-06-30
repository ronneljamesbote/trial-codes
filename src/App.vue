<template>
  <div class="app">
    <Header />
    <Content>
      <h1 class="app-title">Famous Persons</h1>
      <div class="cards-container">
        <AddCard :selected="selectedCard === null" @create-person="handleCreatePerson" />
        <PersonCard
          v-for="person in persons"
          v-bind="{person}"
          :key="person.ID"
          :selected="selectedCard === person.ID"
          @delete-person="handleDeletePerson"
          @selected-changed="handleSelectedChanged"
          @update-person="handleUpdatePerson"
        />
      </div>
    </Content>
    <Footer />
  </div>
</template>

<script>
import persons from "./data/persons";
import moment from "moment";
import AddCard from "./components/AddCard.vue";
import Content from "./components/layout/Content.vue";
import Footer from "./components/layout/Footer.vue";
import Header from "./components/layout/Header.vue";
import PersonCard from "./components/PersonCard.vue";
import { setTimeout } from "timers";

export default {
  name: "app",
  components: {
    AddCard,
    Content,
    Footer,
    Header,
    PersonCard
  },
  data() {
    return {
      selectedCard: null,
      persons: []
    };
  },
  created() {
    this.getPersons();
  },
  methods: {
    handleSelectedChanged(id) {
      this.selectedCard = id;
    },
    handleCreatePerson(person, success) {
      const lastId = this.persons[0].ID;
      const newPerson = { ID: lastId + 1, ...person };

      this.persons.unshift(newPerson);

      typeof success === "function" && success(newPerson);
    },
    handleDeletePerson(id, success) {
      // Save person to be deleted
      const temp = this.persons.filter(p => p.ID === id);

      this.persons = this.persons.filter(p => p.ID != id);

      typeof success === "function" && success(temp);
    },
    handleUpdatePerson(person, success) {
      const index = this.persons.findIndex(p => p.ID === person.ID);
      this.$set(this.persons, index, person);

      typeof success === "function" && success(person);
    },
    getPersons(success) {
      this.persons = persons.reverse();

      typeof success === "function" && success(this.persons);
    }
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
@media (min-width: 992px) {
  /* Correct text direction */
  .app .app-title {
    text-align: left;
  }

  /* Display correct spacing between cards */
  .app .cards-container {
    justify-content: flex-start;
  }
}
</style>
