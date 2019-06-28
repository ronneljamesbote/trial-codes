<template>
  <Card :class="{'selected': selected}">
    <PersonForm v-model="person" :reset="reset" @form-valid="getFormStatus"/>
    <Button class="create-button" :disabled="!cansave" @click.native="handleCreateClick">Create</Button>
  </Card>
</template>

<script>
import Button from "./ui/Button.vue";
import Card from "./ui/Card.vue";
import PersonForm from "./PersonForm.vue";

export default {
  name: "add-card",
  components: {
    Button,
    Card,
    PersonForm
  },
  props: {
    selected: { type: Boolean, default: false }
  },
  data() {
    return {
      reset: false,
      cansave: false,
      person: {
        Bio: "",
        DateOfBirth: "",
        Name: ""
      }
    };
  },
  watch: {
    person: {
      handler(object) {
        const isAllEmpty =
          Object.values(object)
            .join("")
            .trim().length === 0;

        if (isAllEmpty) this.reset = true;
        else this.reset = false;
      },
      deep: true
    }
  },
  methods: {
    handleCreateClick() {
      this.$emit("create-person", this.person, this.clearForm);
    },
    getFormStatus(status) {
      this.cansave = status;
    },
    clearForm() {
      this.person = {
        Bio: "",
        DateOfBirth: "",
        Name: ""
      };
    }
  }
};
</script>

<style scoped>
button {
  position: absolute;
}

.create-button {
  right: 24px;
}
</style>
