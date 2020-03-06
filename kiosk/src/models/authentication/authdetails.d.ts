export interface LoginDetails {
  identification: string;
  token: string;
  authMethodId: number;
}

export interface authPayload {
  error: string;
  tokenToStore: string;
}

export interface IRegisterDetails {
  firstName: string;
  lastName: string;
  email: string;
  token: string;
  authMethodId: number;
}

export interface IRegisterResponse {
  success: boolean;
  message: string;
}