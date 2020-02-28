export interface LoginDetails {
  identification: string;
  token: string;
  authMethodId: number;
}

export interface authPayload {
  error: string;
  tokenToStore: string;
}

export interface RegisterDetails {
  firstName: string;
  lastName: string;
  email: string;
  token: string;
  authMethodId: number;
}