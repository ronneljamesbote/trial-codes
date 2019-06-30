<template functional>
  <div class="content">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "content-component"
};
</script>

<style scoped>
.content {
  flex: 1;
  margin: 2em 0;
  padding: 1em;
  padding: 1em calc(50% - 450px);
}

@media screen and (max-width: 900px) {
  .content {
    padding: 1.5em;
  }
}
</style>

