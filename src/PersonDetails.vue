<template>
  <div v-if="details">
    <p class="person-name">{{details.Name}}</p>
    <p class="person-birthdate">{{details.DateOfBirth}}</p>
    <p class="person-bio">{{details.Bio}}</p>
  </div>
</template>

<script>
export default {
  name: "person-details",
  props: {
    details: Object
  }
};
</script>

<style>
.person-name {
  font-size: 1.3em;
  font-weight: 700;
  margin-bottom: 0.5em;
}

.person-birthdate {
  font-size: 0.95em;
  margin-bottom: 1.5em;
  opacity: 0.8;
}
</style>
