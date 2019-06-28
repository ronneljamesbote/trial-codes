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
import AddCard from "./AddCard.vue";
import Content from "./layout/Content.vue";
import Footer from "./layout/Footer.vue";
import Header from "./layout/Header.vue";
import PersonCard from "./PersonCard.vue";

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
    const vm = this;

    axios
      .get("/persons")
      .then(res => {
        vm.persons = res.data.data;
      })
      .catch(err => {
        window.alert("Something broke! Please try again later.");
        console.log(err);
      });
  },
  methods: {
    handleCreatePerson(person, successCb) {
      const vm = this;
      const date = moment.utc(person.DateOfBirth);
      person.DateOfBirth = date.format("x");

      axios
        .post("/persons", qs.stringify(person))
        .then(res => {
          vm.persons.push(res.data.data);

          typeof successCb === "function" && successCb();
        })
        .catch(err => {
          window.alert("Something broke! Please try again later.");
          console.log(err);
        });
    },
    handleDeletePerson(id) {
      const vm = this;

      axios
        .delete(`/persons/${id}`)
        .then(res => {
          vm.persons = vm.persons.filter(person => {
            return person.ID !== id;
          });
        })
        .catch(err => {
          window.alert("Something broke! Please try again later.");
          console.log(err);
        });
    },
    handleUpdatePerson(person, successCb) {
      const vm = this;
      const date = moment.utc(person.DateOfBirth);
      person.DateOfBirth = date.format("x");

      axios
        .put(`/persons/${person.ID}`, qs.stringify(person))
        .then(res => {
          const index = vm.persons.findIndex(value => {
            return value.ID === person.ID;
          });

          vm.$set(vm.persons, index, res.data.data);
          typeof successCb === "function" && successCb(res.data.data);
        })
        .catch(err => {
          window.alert("Something broke! Please try again later.");
          console.log(err);
        });
    },
    handleSelectedChanged(id) {
      this.selectedCard = id;
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
