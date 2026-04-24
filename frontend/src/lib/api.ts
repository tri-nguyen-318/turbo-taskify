import axios from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api",
  withCredentials: true,
});


export interface AuthUser {
  id: string;
  email: string;
  username: string;
  language: string;
  theme: string;
}

export interface AuthResponse {
  success: boolean;
  user: AuthUser;
  accessToken: string;
  refreshToken?: string;
  isNewUser?: boolean;
}

export interface SignInPayload {
  email: string;
  password: string;
}

export interface SignUpPayload {
  email: string;
  password: string;
  username: string;
}

export interface GoogleSignInPayload {
  idToken: string;
}
