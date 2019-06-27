<template>
  <Card
    :class="{'selected': isSelected}"
    @mouseenter.native="handleHover"
    @mouseleave.native="handleHover"
  >
    <PersonDetails v-if="!isSelected" :details="tempPerson"/>
    <PersonForm
      v-else
      v-model="tempPerson"
      @form-destroyed="handleFormDestroyed"
      @form-valid="getFormStatus"
    />

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
    person: Object,
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
  computed: {
    strDateOfBirth() {
      return moment.utc(this.person.DateOfBirth).format("MMM DD YYYY");
    },
    isSelected() {
      return this.editting && this.selected;
    }
  },
  created() {
    if (this.person) {
      this.resetTempValue();
    }
  },
  methods: {
    handleEditClick() {
      this.editting = true;

      this.$emit("selected-changed", this.person.ID);
    },
    handleDeleteClick() {
      if (!window.confirm("You are deleting a person. Continue?")) return;

      this.$emit("delete-person", this.person.ID);
    },
    handleSaveClick() {
      const vm = this;
      this.editting = false;

      this.$emit("selected-changed", null);
      this.$emit("update-person", this.tempPerson, function(updatedPerson) {
        vm.tempPerson = {
          ID: updatedPerson.ID,
          Name: updatedPerson.Name,
          DateOfBirth: moment
            .utc(updatedPerson.DateOfBirth)
            .format("MMM DD YYYY"),
          Bio: updatedPerson.Bio
        };
      });
    },
    handleCancelClick() {
      this.editting = false;

      this.$emit("selected-changed", null);
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
    handleFormDestroyed() {
      this.resetTempValue();
    },
    getFormStatus(status) {
      this.cansave = status;
    },
    resetTempValue() {
      this.tempPerson = {
        ID: this.person.ID,
        Name: this.person.Name,
        DateOfBirth: this.strDateOfBirth,
        Bio: this.person.Bio
      };
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
