<template>
  <v-app>
    <v-main>
    <v-tabs
    bg-color="primary"
      next-icon="mdi-arrow-right-thin"
      prev-icon="mdi-arrow-left-thin"
      show-arrows
      >
      <v-tab @click="chest = true ; forearm = false ; biceps = false ; shoulder = false ; triceps = false ; back = false ; legs = false" >Chest</v-tab>
      <v-tab @click="biceps = true ; forearm = false ; chest = false ; shoulder = false ; triceps = false ; back = false ; legs = false"  >Biceps</v-tab>
      <v-tab @click="forearm = true ; chest = false ; biceps = false ; shoulder = false ; triceps = false ; back = false ; legs = false" >Forearm</v-tab>
      <v-tab @click="triceps = true ; chest = false ; biceps = false ; forearm = false ; shoulder = false  ; back = false ; legs = false" >Triceps</v-tab>
      <v-tab @click="shoulder = true ; chest = false ; biceps = false ; forearm = false ; triceps = false ; back = false ; legs = false" >Shoulder</v-tab>
      <v-tab @click="back = true ; chest = false ; biceps = false ; forearm = false ; shoulder = false  ; triceps = false ; legs = false" >Back</v-tab>
      <v-tab @click="legs = true ; chest = false ; biceps = false ; forearm = false ; shoulder = false  ; triceps = false ; back = false" >Legs</v-tab>
    </v-tabs>
    <v-divider horizontal></v-divider>
    <div>
      <Forearm v-if="forearm" />
      <Triceps v-if="triceps" />
      <Shoulder v-if="shoulder" />
      <Legs v-if="legs" />
      <Chest v-if="chest" />
      <Biceps v-if="biceps" />
      <Back v-if="back" />
    </div>
  </v-main>
</v-app>
</template>
<script>
import Forearm from "@/views/Workout/Forearm.vue"
import Triceps from "@/views/Workout/Triceps.vue"
import Shoulder from "@/views/Workout/Shoulder.vue"
import Legs from "@/views/Workout/Legs.vue"
import Chest from "@/views/Workout/Chest.vue"
import Biceps from "@/views/Workout/Biceps.vue"
import Back from "@/views/Workout/Back.vue"


  export default {
    data () {
      return {
        forearm : false,
        triceps : false,
        shoulder : false,
        legs : false,
        chest : true,
        biceps : false,
        back : false,
      }
    },
    components:{
      Forearm,
      Triceps,
      Shoulder,
      Legs,
      Chest,
      Biceps,
      Back,
    }
  }
</script>

