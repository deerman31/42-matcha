// src/types/api.ts として型定義を分離

export interface RegisterResponse {
  message?: string;
  error?: string;
}

export interface LoginResponse {
  is_preparation: boolean;
  access_token: string;
}

export interface ErrorResponse {
  error: string;
}
