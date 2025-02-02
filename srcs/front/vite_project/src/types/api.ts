// src/types/api.ts として型定義を分離

export interface Response {
  message: string;
}

export interface LoginResponse {
  is_preparation: boolean;
  access_token: string;
}

export interface ErrorResponse {
  error: string;
}
