<template>
  <div class="dropdown">

    <ul>
      <li>
        <router-link
          :to="{name: route.name}"
          class="nav-link"
        >{{name}}</router-link>
      </li>
    </ul>

  </div>
</template>

<script>
/* eslint-disable semi */
/* eslint-disable no-console */
// eslint-disable-next-line no-unused-vars
export default {
  name: 'menu-item',
  components: {
  },
  props: {
    route: {
      type: Object
    }
  },
  computed: {
    name () {
      return this.route.name;
    }
  }
}
</script>

<style scoped>


</style>
