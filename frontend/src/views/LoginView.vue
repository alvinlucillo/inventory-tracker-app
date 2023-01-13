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
          <v-alert
            dense
            text
            type="success"
            v-if="isLoginSuccessful"
            @click="hide('success-alert')"
          >
            You have successfully logged in
          </v-alert>
          <v-alert
            dense
            outlined
            type="error"
            v-if="isLoginError"
            @click="hide('fail-alert')"
          >
            {{ loginErrorMsg }}
          </v-alert>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { loginUser } from "../service/user_service";

export default {
  name: "LoginView",
  data() {
    return {
      username: "",
      password: "",
      isLoginError: false,
      loginErrorMsg: "",
      isLoginSuccessful: false,
      divStyle: {
        marginLeft: "30%",
        marginRight: "30%",
      },
    };
  },
  methods: {
    async login() {
      this.isLoginSuccessful = false;
      this.isLoginError = false;
      this.loginErrorMsg = "";

      const { username, password } = this;

      const result = await loginUser(username, password);

      if (result.isSuccess) {
        this.isLoginSuccessful = true;
      } else {
        this.isLoginError = true;
        this.loginErrorMsg = result.message;

        console.log(result);
      }
    },
    hide(alertType) {
      if (alertType == "success-alert") {
        this.isLoginSuccessful = false;
      }
      if (alertType == "fail-alert") {
        this.isLoginError = false;
      }
    },
  },
};
</script>
