<template>
  <button>
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: "button-component"
};
</script>

<style scoped>
button {
  background: none;
  border: 0;
  bottom: 24px;
  color: #2c3e50;
  cursor: pointer;
  outline: 0;
  padding: 0;
}

button:disabled {
  cursor: not-allowed;
  opacity: 0.7;
}
</style>
