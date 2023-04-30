<template>
<v-app>
    <v-main>
            <v-tabs
            bg-color="primary"
            grow
          >
            <v-tab @click="bmi = true ; protein = false ; water = false" > BMI </v-tab>

            <v-tab @click="protein = true ; bmi = false ; water = false"> PROTEIN </v-tab>

            <v-tab @click="water = true ; protein = false ; bmi = false"> WATER </v-tab>
          </v-tabs>
      <v-divider horizontal></v-divider>

          <div>
            <Bmi v-if="bmi" />
            <Protein v-if="protein" />  
            <Water v-if="water" />
          </div>
    </v-main>
</v-app>
    
</template>

<script>

import Bmi from "@/calculate/bmi.vue"
import Protein from "@/calculate/protein.vue"
import Water from "@/calculate/water.vue"

export default {
    name: "App",
    data() {
        return {
            bmi: true,
            protein: false,
            water: false,
        }
    },
    components: {
        Bmi,
        Protein,
        Water
    },
}
</script>

