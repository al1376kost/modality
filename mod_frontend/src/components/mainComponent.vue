<template>
    <div class="wrapper">
    <div id="app-private">

      <header>
        <div class="wrap">

          <div class="menu-comp">
            <menu-item v-for="(r,i) in routes" :key="i" :route="r"></menu-item>
          </div>

        </div>
      </header>

      <keep-alive>
        <router-view class="view"></router-view>
      </keep-alive>

    </div>
    </div>
</template>

<script>
    export default {
        name: "mainComponent",
        components: {
            menuItem: () => import(/* webpackChunkName: "menu-item" */ './menu.vue'),
        },
        computed: {
            routes() {
                let routes = this.$router.options.routes.find(item => item.name === 'mainComponent').children;
                return routes;
            }
        }
        // beforeCreate() {
        //     this.$router.push({
        //       name: 'modalityComponent'
        //     });
        // }
    }
</script>

<style scoped>
    .wrapper {
        width: 100%;
    }
</style>