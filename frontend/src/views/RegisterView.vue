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
          <v-text-field
            label="Reenter password"
            required
            v-model="reenterPassword"
          ></v-text-field>
          <v-btn @click="register"> Register </v-btn>

          <v-alert
            dense
            text
            type="success"
            v-if="isRegisterSuccessful"
            @click="hide('success-alert')"
          >
            You have successfully registered
          </v-alert>
          <v-alert
            dense
            outlined
            type="error"
            v-if="isRegisterError"
            @click="hide('fail-alert')"
          >
            {{ registerErrorMsg }}
          </v-alert>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { registerUser } from "../service/user_service";

export default {
  name: "RegisterView",
  data() {
    return {
      username: "",
      password: "",
      reenterPassword: "",
      isRegisterError: false,
      registerErrorMsg: "",
      isRegisterSuccessful: false,
      divStyle: {
        marginLeft: "30%",
        marginRight: "30%",
      },
    };
  },
  methods: {
    async register() {
      this.isRegisterSuccessful = false;
      this.isRegisterError = false;
      this.registerErrorMsg = "";

      const { username, password, reenterPassword } = this;

      if (password != reenterPassword) {
        this.isRegisterError = true;
        this.registerErrorMsg = "Passwords must match.";
        return;
      }

      const result = await registerUser(username, password);

      if (result.isSuccess) {
        this.isRegisterSuccessful = true;
      } else {
        this.isRegisterError = true;
        this.registerErrorMsg = result.message;

        console.log(result);
      }
    },
    hide(alertType) {
      if (alertType == "success-alert") {
        this.isRegisterSuccessful = false;
      }
      if (alertType == "fail-alert") {
        this.isRegisterError = false;
      }
    },
  },
};
</script>
