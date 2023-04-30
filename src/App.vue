<template>
  <v-app>

    <v-app-bar 
    class="flex-grow-0"
    color="primary"
    dark
    flat
    >
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer" ></v-app-bar-nav-icon>

      <v-toolbar-title >WorkOutMate</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

    </v-app-bar>

    <v-navigation-drawer
        v-model="drawer"
        location="left"
        temporary
        color="primary"
      >
      <v-list density="compact" nav>
      <v-divider horizontal></v-divider>
        <v-list-item class="texxt" prepend-icon="mdi-account" title="My Profile" value="account"></v-list-item>
      <v-divider horizontal></v-divider>
        <v-list-item to="/Home" prepend-icon="mdi-home-variant-outline" title="Home" value="home"></v-list-item>
      <v-divider horizontal></v-divider>
        <v-list-item to="/Exercise" prepend-icon="mdi-weight-lifter" title="Exercise" value="exercise"></v-list-item>
      <v-divider horizontal></v-divider>
        <v-list-item to="/Calculate" prepend-icon="mdi-calculator" title="Health Calculator" value="Calculator"></v-list-item>
      <v-divider horizontal></v-divider>

      </v-list>
      </v-navigation-drawer>
 
  <v-main>
    <router-view/>
  </v-main>

    <v-footer>
      <v-bottom-navigation
      mode="teal"
      grow
      bg-color="primary"
      dark
       >
      <v-btn to="/Home">
        <v-icon>mdi-home</v-icon>
        Home
      </v-btn>
      <v-divider vertical></v-divider>
      <v-btn to="/Exercise">
        <v-icon>mdi-weight-lifter</v-icon>
        Exercise
      </v-btn>
      <v-divider vertical></v-divider>
      <v-btn to="/Calculate">
        <v-icon>mdi-heart</v-icon>
        Calculator
      </v-btn>

    </v-bottom-navigation>
    </v-footer>
    
     
  </v-app>
  
</template>

<script>
import HelloWorld from '@/components/HelloWorld.vue'

  export default {
    name : "App",
  data: () => ({
    color: 'purple accent-4',
    drawer: false,
      group: null,
      
  }),
  theme: {
    dark: true,
    themes: {
      dark: {
        primary: '#1976D2',
        secondary: '#424242',
        accent: '#82B1FF',
        error: '#FF5252',
        info: '#2196F3',
        success: '#4CAF50',
        warning: '#FFC107'
      }
    }
  },
   components: {
    HelloWorld
  },
   watch: {
      group () {
        this.drawer = false
      },
    },

  }
</script>

<style>
.texxt{
font-size: large;
}

</style>