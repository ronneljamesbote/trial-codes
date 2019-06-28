<template>
  <div class="app">
    <Header/>
    <Content>
      <h1 class="app-title">Famous Persons</h1>
      <div class="cards-container">
        <AddCard :selected="selectedCard === null" @create-person="handleCreatePerson"/>
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
    <Footer/>
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
import moment from "moment";
import AddCard from "./components/AddCard.vue";
import Content from "./components/layout/Content.vue";
import Footer from "./components/layout/Footer.vue";
import Header from "./components/layout/Header.vue";
import PersonCard from "./components/PersonCard.vue";

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
    handleCreatePerson(person, success, error) {
      const date = moment.utc(person.DateOfBirth);
      person.DateOfBirth = date.isValid() && date.format("x");

      axios
        .post("/persons", qs.stringify(person))
        .then(res => {
          this.persons.unshift(res.data.data);

          typeof success === "function" && success(res);
        })
        .catch(err => {
          typeof error === "function" && error(err);

          console.log(err);
          window.alert("Something broke! Please try again later.");
        });
    },
    handleDeletePerson(id, success, error) {
      axios
        .delete(`/persons/${id}`)
        .then(res => {
          this.persons = this.persons.filter(value => {
            return value.ID !== res.data.data.ID;
          });

          typeof success === "function" && success(res);
        })
        .catch(err => {
          typeof error === "function" && error(err);

          console.log(err);
          window.alert("Something broke! Please try again later.");
        });
    },
    handleUpdatePerson(person, success, error) {
      const date = moment.utc(person.DateOfBirth);
      person.DateOfBirth = date.isValid() && date.format("x");

      axios
        .put(`/persons/${person.ID}`, qs.stringify(person))
        .then(res => {
          const index = this.persons.findIndex(value => {
            return value.ID === person.ID;
          });
          this.$set(this.persons, index, res.data.data);

          typeof success === "function" && success(res);
        })
        .catch(err => {
          typeof error === "function" && error(err);

          console.log(err);
          window.alert("Something broke! Please try again later.");
        });
    },
    getPersons(success, error) {
      axios
        .get("/persons")
        .then(res => {
          this.persons = res.data.data.reverse();

          typeof success === "function" && success(res);
        })
        .catch(err => {
          typeof error === "function" && error(err);

          console.log(err);
          window.alert("Something broke! Please try again later.");
        });
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
