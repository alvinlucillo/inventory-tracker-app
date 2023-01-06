<template>
  <v-app>
    <v-main>
      <v-container>
        <div :style="divStyle">
          <v-text-field
            label="Username"
            required
            v-model="username"
          ></v-text-field>
          <v-text-field
            label="Password"
            required
            v-model="password"
          ></v-text-field>
          <v-btn @click="login"> Login </v-btn>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";

const svc = axios.create({ baseURL: process.env.VUE_APP_API });

export default {
  name: "LoginView",
  data() {
    return {
      username: "",
      password: "",
      divStyle: {
        marginLeft: "30%",
        marginRight: "30%",
      },
    };
  },
  methods: {
    async login() {
      const { username, password } = this;

      const result = await svc.post("auth/login", {
        username,
        password,
      });

      console.log(result);

      console.log(username, password);
    },
  },
};
</script>
