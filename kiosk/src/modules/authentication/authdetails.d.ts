export interface LoginDetails {
  identification: string;
  token: string;
  authMethodId: number;
}

interface authPayload {
  error: string;
  tokenToStore: string;
}
