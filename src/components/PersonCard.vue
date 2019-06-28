<template>
  <Card
    :class="{'selected': isSelected}"
    @mouseenter.native="handleHover"
    @mouseleave.native="handleHover"
  >
    <PersonDetails v-if="!isSelected" :details="tempPerson"/>
    <PersonForm v-else v-model="tempPerson" @form-valid="getFormStatus"/>

    <div v-if="mouseon && !editting">
      <Button class="delete-button" @click.native="handleDeleteClick">Delete</Button>
      <Button class="edit-button" @click.native="handleEditClick">Edit</Button>
    </div>
    <div v-if="isSelected">
      <Button class="cancel-button" @click.native="handleCancelClick">Cancel</Button>
      <Button class="save-button" :disabled="!cansave" @click.native="handleSaveClick">Save</Button>
    </div>
  </Card>
</template>

<script>
import moment from "moment";
import Button from "./ui/Button.vue";
import Card from "./ui/Card.vue";
import PersonDetails from "./PersonDetails.vue";
import PersonForm from "./PersonForm.vue";

export default {
  name: "person-card",
  components: {
    Button,
    Card,
    PersonDetails,
    PersonForm
  },
  props: {
    person: { type: Object, required: true },
    selected: { type: Boolean, default: false }
  },
  data() {
    return {
      cansave: true,
      editting: false,
      mouseon: false,
      tempPerson: {
        ID: 0,
        Bio: "",
        DateOfBirth: "",
        Name: ""
      }
    };
  },
  watch: {
    person() {
      this.initTempPerson();
    },
    selected(status) {
      if (!status) this.initTempPerson();
    }
  },
  computed: {
    compDateOfBirth() {
      return moment.utc(this.person.DateOfBirth).format("MMM DD YYYY");
    },
    isSelected() {
      return this.editting && this.selected;
    }
  },
  created() {
    this.initTempPerson();
  },
  methods: {
    handleEditClick() {
      this.editting = true;

      this.emitSelectedChanged(this.person.ID);
    },
    handleDeleteClick() {
      if (!window.confirm("You are deleting a person. Continue?")) return;

      this.$emit("delete-person", this.person.ID);
    },
    handleSaveClick() {
      this.editting = false;

      this.emitSelectedChanged();
      this.$emit("update-person", this.tempPerson);
    },
    handleCancelClick() {
      this.editting = false;

      this.initTempPerson();
      this.emitSelectedChanged();
    },
    handleHover(event) {
      if (this.isSelected) return;

      switch (event.type) {
        case "mouseleave":
          this.mouseon = false;
          break;
        default:
          this.mouseon = true;
          break;
      }

      this.editting = false;
    },
    getFormStatus(status) {
      this.cansave = status;
    },
    initTempPerson() {
      this.tempPerson = {
        ID: this.person.ID,
        Name: this.person.Name,
        DateOfBirth: this.compDateOfBirth,
        Bio: this.person.Bio
      };
    },
    emitSelectedChanged(selected) {
      this.$emit("selected-changed", selected === undefined ? null : selected);
    }
  }
};
</script>

<style scoped>
button {
  position: absolute;
}

.cancel-button,
.delete-button {
  right: 70px;
}

.edit-button,
.save-button {
  right: 24px;
}
</style>
