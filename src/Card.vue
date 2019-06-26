<template>
  <div
    class="card"
    :class="{'selected': editting}"
    @mouseenter="handleHover"
    @mouseleave="handleHover"
  >
    <div v-if="!this.editting">
      <p class="person-name">{{person.name}}</p>
      <p class="person-birthdate">{{person.birthdate}}</p>
      <p class="person-bio">{{person.bio}}</p>
    </div>

    <div v-else>
      <input class="person-name-input" type="text" v-model="tempPerson.name">
      <input class="person-birthdate-input" type="text" v-model="tempPerson.birthdate">
      <textarea class="person-bio-input" v-model="tempPerson.bio"/>
    </div>

    <button class="cancel-button" v-show="editting" @click="handleCancelClick">Cancel</button>
    <button
      class="edit-button"
      :disabled="!cansave && editting"
      v-show="hovering"
      @click="handleSaveClick"
    >{{ buttonLabel }}</button>
  </div>
</template>

<script>
export default {
  name: "cards",
  props: {
    person: Object
  },
  data() {
    return {
      cansave: true,
      editting: false,
      hovering: false,
      tempPerson: {
        name: "",
        birthdate: "",
        bio: ""
      }
    };
  },
  computed: {
    buttonLabel() {
      return this.editting ? "Save" : "Edit";
    }
  },
  watch: {
    tempPerson: {
      handler(object) {
        if (
          object.name.trim() === "" ||
          object.birthdate.trim() === "" ||
          object.bio.trim() === ""
        ) {
          this.cansave = false;
          return;
        }

        this.cansave = true;
      },
      deep: true
    }
  },
  methods: {
    handleCancelClick() {
      this.tempPerson = {
        name: "",
        birthdate: "",
        bio: ""
      };
      this.editting = false;
    },
    handleSaveClick() {
      if (!this.editting) {
        this.editting = true;
        this.tempPerson = Object.assign({}, this.person);
        return;
      }

      // Put API request here //

      this.editting = false;

      window.alert("Work in progress. Not implemented yet.");
    },
    handleHover(event) {
      if (this.editting) return;

      switch (event.type) {
        case "mouseleave":
          this.hovering = false;
          break;
        default:
          this.hovering = true;
          break;
      }
    }
  }
};
</script>

<style scoped>
.card {
  background: #fff;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  margin: 0 1.5em;
  margin-bottom: 3em;
  min-height: 15em;
  padding: 1em;
  position: relative;
  transition: box-shadow 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  width: 100%;
}

.card.selected {
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
}

.card:hover {
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
}

.person-name,
.person-name-input {
  font-size: 1.3em;
  font-weight: 700;
  margin-bottom: 0.5em;
}

.person-birthdate {
  font-size: 0.95em;
  margin-bottom: 1.5em;
  opacity: 0.8;
}

.person-name-input,
.person-birthdate-input,
.person-bio-input {
  display: block;
  padding: 0.1em;
  width: 95%;
}

.person-name-input {
  margin-bottom: 0.5em;
  margin-top: 1em;
}

.person-birthdate-input {
  margin-bottom: 1em;
}

.person-bio-input {
  height: 6em;
  resize: none;
}

.cancel-button,
.edit-button {
  background: none;
  border: 0;
  bottom: 24px;
  color: #2c3e50;
  cursor: pointer;
  outline: 0;
  padding: 0;
  position: absolute;
}

button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.cancel-button {
  right: 70px;
}

.edit-button {
  right: 24px;
}

/* Show cards by 2 - for medium screen sizes */
@media screen and (min-width: 576px) {
  .card {
    width: 30%;
  }
}

/* Show card by 3 - for big screen sizes */
@media screen and (min-width: 769px) {
  .card {
    margin-left: 0;
    margin-right: 0;
    width: 25%;
  }
}
</style>

