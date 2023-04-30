<template>
<v-app class="cal">
<h4 class="text-center">
    Calculate Body Mass Index (BMI)
</h4>

<div class="inpt">
    <v-text-field v-model="weight" clearable type="number" label="Enter Weight (In Kg)" variant="outlined"></v-text-field>
    <v-text-field v-model="height" clearable type="number" label="Enter Height (In Cm)" variant="outlined"></v-text-field>
</div>

<div class="text-center">
    <v-dialog v-model="dialog" width="100%"  >
      <template v-slot:activator="{ props }">
        <v-btn @click="Calculate()" color="primary" v-bind="props"> Submit </v-btn>
      </template>

      <v-card>
        <v-card-text class="teext">
            Your Body Mass Index (BMI) is {{ getbmivalue }}
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" block @click="dialog = false">Ok</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
  <br>
  <v-divider horizontal></v-divider>

  <div>
    <v-table>
        <thead>
          <tr>
            <th class="text-center">
                BMI Range
            </th>
            <th class="text-center">
                Health Category
            </th>
          </tr>
        </thead>
        <tbody class="text-center">
          <tr
            v-for="item in bmiresult"
            :key="item.bmi"
          >
            <td>{{ item.bmi }}</td>
            <td>{{ item.catagory }}</td>
          </tr>
        </tbody>
      </v-table>
  </div>
  <v-divider horizontal></v-divider>
</v-app>

</template>

<script>
import axios from 'axios';

export default {
name: 'App',
  data: function() {
    return {
        dialog: false,
        weight:'',
        height:'',
        getbmivalue:'',
        bmiresult: [
          {
            bmi: 'Less than 18.5',
            catagory: "Under Weight",
          },
          {
            bmi: '18.5 to 24.9',
            catagory: "Normal Weight",
          },
          {
            bmi: '25 to 29.9',
            catagory: "Over Weight",
          },
          {
            bmi: '30 to 34.9',
            catagory: "Obese Class I",
          },
          {
            bmi: '35 to 39.9',
            catagory: "Obese Class II	",
          },
          {
            bmi: '40 or greater',
            catagory: "Obese Class III	",
          }
        ],
    } 
  },

  methods:{
    Calculate(){
        let a = parseFloat(this.weight)
        let b = parseFloat(this.height)
        axios.post('http://localhost:8080/GetBmi', {
            Weight: a,
            Height: b,
            })
              .then((response) => {
              console.log(response.data);
              const responseData = response.data;
              const bmivalueObj = JSON.parse(responseData.split('\n')[1]);
              const bmivalue = bmivalueObj.bmivalue;
              this.getbmivalue = bmivalue;
            })
              .catch(error => {
              console.log(error);
            })
            
        },
    }
}

</script>

<style>
.cal{
    margin-top: 10%;
}
.inpt{
    width: 80%;
    margin-left: 35px;
    margin-top: 20px;
}
.subbtm{
    margin-left: 36%;
}
.teext{
    text-align: center;
}
</style>