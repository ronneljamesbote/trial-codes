<template>
  <div class="card">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "card-component"
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
  padding-bottom: 2em;
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

/*  For medium screen sizes */
@media (min-width: 576px) {
  /* Show cards by 2 */
  .card {
    width: 30%;
  }
}

/* For big screen sizes */
@media (min-width: 768px) {
  /* Show card by 3 */
  .card {
    margin-left: 0;
    margin-right: 3%;
    width: 26%;
  }
}
</style>
