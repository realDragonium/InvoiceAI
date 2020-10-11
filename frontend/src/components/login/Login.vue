<template>
  <div>
    <v-text-field v-model="user.username" label="Username"></v-text-field>
    <v-text-field v-model="user.email" label="Email"></v-text-field>
    <v-text-field v-model="user.password" label="Password" type="password"></v-text-field>
    <v-btn elevation="2" @click="sendLogin">login</v-btn>
    <v-btn elevation="2" @click="sendCreateAccount">create account</v-btn>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Login",
  data() {
    return {
      user: {
        username: '',
        email: '',
        password:''
      }
    }
  },
  methods: {
    async sendLogin(){
      axios.post('http://localhost:3000/api/v1/user', this.user)
          .then(resp => {console.log(resp)})
    },
    async sendCreateAccount () {
      axios.post('http://localhost:3000/api/v1/user', this.user)
      .then(resp => {console.log(resp)})
    }
  }
}
</script>

<style scoped>

</style>