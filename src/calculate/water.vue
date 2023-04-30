<template>

    <v-app class="cal">
    <h4 class="text-center">
        Calculate Water Intake Per Day
    </h4>
    <div class="inpt">
        <v-text-field v-model="weight" clearable type="number" label="Enter Weight (In Kg)" variant="outlined"></v-text-field>
    </div>
    
    <div class="text-center">
        <v-dialog v-model="dialog" width="100%"  >
          <template v-slot:activator="{ props }">
            <v-btn @click="Calculate()" color="primary" v-bind="props"> Submit </v-btn>
          </template>
    
          <v-card>
            <v-card-text class="teext">
                Your recommended water intake would be approximately {{ Waterintake }} liters per day
            </v-card-text>
            <v-card-actions>
              <v-btn color="primary" block @click="dialog = false">Ok</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
    </v-app>
    
    </template>
    
    <script>
        import axios from 'axios';

    export default {
      name: 'App',
      data: function(){
        return{
            dialog : false,
            weight :'',
            Waterintake : '',
        }
    },

      methods:{
        Calculate(){
            let a = parseFloat(this.weight)
            axios.post('http://localhost:8080/GetWat', {
            Weight: a,
            })
            .then((response) => {
            console.log(response.data);
            const responseData = response.data;
            const proteinValueObj = JSON.parse(responseData.split('\n')[1]); 
            const waterValue = proteinValueObj.watervalue;
            this.Waterintake = waterValue
            console.log(this.Waterintake);
            })
            .catch(error => {
              console.log(error);
            })
        },
    },
}
    
</script>

<style>

</style>