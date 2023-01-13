import axios from "axios";

import { ApiResponse } from "../models/api_response";
import { handleError } from "./common_service";

const svc = axios.create({ baseURL: process.env.VUE_APP_API });

const loginUser = async (username: string, password: string) => {
  const response: ApiResponse = new ApiResponse();
  try {
    const result = await svc.post("auth/login", {
      username,
      password,
    });

    response.isSuccess = true;
    response.data = result.data;
  } catch (err) {
    return handleError(err);
  }

  return response;
};

const registerUser = async (username: string, password: string) => {
  const response: ApiResponse = new ApiResponse();
  try {
    const result = await svc.post("auth/register", {
      username,
      password,
    });

    response.isSuccess = true;
    response.data = result.data;
  } catch (err) {
    return handleError(err);
  }

  return response;
};

export { loginUser, registerUser };
