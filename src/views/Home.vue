<template>
 <v-app dark class="hmbgcol">
    <div class="bgcoll">
        <img class="homeimg" src="@/images/dumbbell.png" >
        <h2 class="headtext">
            WorkOutMate
        </h2>
        <p class="fontsty">Transform your health and fitness routine with </p>
            <p  class="fontstyle">with this app. Track, calculate, and improve.</p>
            <v-btn to="/Calculate" size="large"  class="roundbtn" color="primary">
               Get Started  <v-icon>mdi-play</v-icon> 
            </v-btn>
    </div>
    
 </v-app>
</template>

<script>

export default {
name: "App",
data() {
       
}
}
</script>
<style>
.hmbgcol{
    margin: 0;
    padding: 0;
}
.bgcoll {
    height: 48%;
    width: 100%;
    background: hsla(307, 93%, 84%, 1);
    background: linear-gradient(90deg, hsla(307, 93%, 84%, 1) 0%, hsla(217, 67%, 47%, 1) 100%);
    background: -moz-linear-gradient(90deg, hsla(307, 93%, 84%, 1) 0%, hsla(217, 67%, 47%, 1) 100%);
    background: -webkit-linear-gradient(90deg, hsla(307, 93%, 84%, 1) 0%, hsla(217, 67%, 47%, 1) 100%);
    float: left;
    border-bottom-left-radius: 40px;
    border-bottom-right-radius: 40px;
  }
  .homeimg{
    position: relative;
    right: 15px;
    top: 45px;
    width: 420px;
}
.fontsty{
    text-align: center;
    font-family:'Arial, Helvetica, sans-serif';
    font-size: 18px;
    margin-top: 8px;
}
.fontstyle{
    text-align: center;
    font-family:'Arial, Helvetica, sans-serif';
    font-size: 18px;
}
.roundbtn{
    border-radius: 50px;
    margin-left: 26%;
    margin-top: 6%;
}
.headtext{
    text-align: center;
    font-family:'Arial, Helvetica, sans-serif';
    font-size: 30px;
    margin-top: 60px;
}
  

</style>