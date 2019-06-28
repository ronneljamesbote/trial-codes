<template>
  <div class="card-form">
    <input
      class="person-name-input"
      :class="{'invalid': touched && !valid.Name}"
      placeholder="Fullname"
      type="text"
      v-model="value.Name"
    >
    <input
      class="person-birthdate-input"
      :class="{'invalid': touched && !valid.DateOfBirth}"
      placeholder="Date of birth"
      title="Enter a valid date"
      type="text"
      v-model="value.DateOfBirth"
    >
    <textarea
      class="person-bio-input"
      :class="{'invalid': touched && !valid.Bio}"
      placeholder="Bio"
      v-model="value.Bio"
    />
  </div>
</template>

<script>
import moment from "moment";

export default {
  name: "person-form",
  props: {
    reset: { type: Boolean, default: false },
    value: { type: Object, required: true }
  },
  data() {
    return {
      touched: false,
      valid: {
        Bio: true,
        DateOfBirth: true,
        Name: true
      }
    };
  },
  watch: {
    value: {
      handler(object) {
        this.$emit("input", this.value);
        this.touched = true;

        if (!this.isFormValid(object)) {
          this.$emit("form-valid", false);
          return;
        }

        this.$emit("form-valid", true);
      },
      deep: true
    },
    reset(value) {
      if (value === true) {
        this.resetForm();
      }
    }
  },
  destroyed() {
    this.$emit("form-destroyed");
    this.$emit("form-valid", true);
  },
  methods: {
    isFormValid(object) {
      if (object.Name.trim().length < 5) {
        this.valid.Name = false;
      } else {
        this.valid.Name = true;
      }

      if (object.Bio.trim().length < 5) {
        this.valid.Bio = false;
      } else {
        this.valid.Bio = true;
      }
      this.touched = true;

      const date = moment.utc(object.DateOfBirth);
      if (
        object.DateOfBirth.trim().length < 8 ||
        !date.isValid() ||
        date.year() < 1000 ||
        date.year() > 5000
      ) {
        this.valid.DateOfBirth = false;
      } else {
        this.valid.DateOfBirth = true;
      }

      return Object.values(this.valid).indexOf(false) === -1;
    },
    resetForm() {
      this.touched = false;
      this.valid = {
        Bio: true,
        DateOfBirth: true,
        Name: true
      };
    }
  }
};
</script>

<style scoped>
.person-name-input,
.person-birthdate-input,
.person-bio-input {
  border-radius: 0;
  border: 1px solid #a9a9a9;
  display: block;
  padding: 0.1em;
  width: 95%;
}

.person-name-input {
  margin-bottom: 0.5em;
  margin-top: 1em;
  font-size: 1.3em;
  font-weight: 700;
  margin-bottom: 0.5em;
}

.person-birthdate-input {
  margin-bottom: 1em;
}

.person-bio-input {
  height: 6em;
  resize: none;
}

.invalid {
  border: 1px solid #dc3545;
}

.invalid:focus,
.invalid:active {
  outline: none;
}
</style>
