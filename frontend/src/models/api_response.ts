export class ApiResponse {
    statusCode: string;
    statusError: any;
    isSuccess: boolean;
    data: any;
    message: string;
  
    constructor(resp?: Partial<ApiResponse>) {
      Object.assign(this, resp);
    }
  }
  