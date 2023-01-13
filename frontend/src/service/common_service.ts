import { ApiResponse } from "../models/api_response";
import { HttpErrorCode } from "../enums/error.enum";

export const handleError = (err: any): ApiResponse => {
  const response: ApiResponse = new ApiResponse();
  response.isSuccess = false;

  try {
    if (err.response) {
      response.statusCode = err.response.status;
      response.statusError = err.response.statusText;
      response.message = err.response.data;

      if (
        response.statusCode == HttpErrorCode.NOT_FOUND_404 ||
        response.statusCode == HttpErrorCode.SERVER_ERROR_500
      ) {
        response.message =
          "Unable to reach the server. Please refresh your browser or contact the administrator.";
      }
    } else {
      response.statusCode = HttpErrorCode.SERVER_ERROR_500;
      response.message = "Server error. Please contact the administrator.";
    }

    console.log(err);
  } catch (err) {
    console.log(err);
    response.statusCode = HttpErrorCode.SERVER_ERROR_500;
    response.message = "Server error. Please contact the administrator.";
  }

  return response;
};
