<template>
<v-app class="cal">
    <h4 class="text-center">
        Calculate Protein Intake Per Day
    </h4>
    <div class="inpt">
        <v-text-field v-model="weight" clearable type="number" label="Enter Weight (In Kg)" variant="outlined"></v-text-field>
    </div>
    <div class="switch">
        <v-switch v-model="Workoutdaily" label="Do You Workout Daily ?" color="primary" inset></v-switch>
    </div>
    
    <div class="text-center">
        <v-dialog v-model="dialog" width="100%"  >
          <template v-slot:activator="{ props }">
            <v-btn @click="Calculate()" color="primary" v-bind="props"> Submit </v-btn>
          </template>
          <v-card>
            <v-card-text class="teext">
                Your recommended Protein intake would be approximately {{ proteinValue }} grams per day.
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
            Workoutdaily : '',
            proteinValue : '',
        }
      },

      methods:{
        Calculate(){
            let a = parseFloat(this.weight)
            axios.post('http://localhost:8080/GetPro', {
            Weight: a,
            Workoutdaily: this.Workoutdaily,
            })
            .then((response) => {
            console.log(response.data);
            const responseData = response.data;
            const workoutDailyObj = JSON.parse(responseData.split('\n')[0]); 
            const proteinValueObj = JSON.parse(responseData.split('\n')[1]); 
            const workproteinValueObj = JSON.parse(responseData.split('\n')[1]); 
            const workoutDaily = workoutDailyObj.Workoutdaily; 
            const proteinValue = proteinValueObj.proteinvalue;
            const workproteinValue = workproteinValueObj.workproteinvalue;
                if (this.Workoutdaily == false) {
                    this.proteinValue = proteinValue;
                } else {
                    this.proteinValue = workproteinValue;
                }
            console.log(workoutDaily, proteinValue,workproteinValue);
            })
              .catch(error => {
              console.log(error);
            })
        },
      }
    }
    
</script>

<style>
.switch{
    margin-left: 20%;
}
</style>