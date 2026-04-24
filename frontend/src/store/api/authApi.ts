import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import type { AuthResponse, AuthUser, SignInPayload, SignUpPayload, GoogleSignInPayload } from "@/lib/api";
import { setUser, clearAuth } from "@/store/slices/authSlice";

export const authApi = createApi({
  reducerPath: "authApi",
  baseQuery: fetchBaseQuery({
    baseUrl: process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080/api",
    credentials: "include",
  }),
  endpoints: (builder) => ({
    signIn: builder.mutation<AuthResponse, SignInPayload>({
      query: (payload) => ({ url: "/auth/login", method: "POST", body: payload }),
      onQueryStarted: async (_, { dispatch, queryFulfilled }) => {
        try {
          const { data } = await queryFulfilled;
          dispatch(setUser(data.user));
        } catch {}
      },
    }),

    signUp: builder.mutation<AuthResponse, SignUpPayload>({
      query: (payload) => ({ url: "/auth/signup", method: "POST", body: payload }),
      onQueryStarted: async (_, { dispatch, queryFulfilled }) => {
        try {
          const { data } = await queryFulfilled;
          dispatch(setUser(data.user));
        } catch {}
      },
    }),

    googleSignIn: builder.mutation<AuthResponse, GoogleSignInPayload>({
      query: (payload) => ({ url: "/auth/google", method: "POST", body: payload }),
      onQueryStarted: async (_, { dispatch, queryFulfilled }) => {
        try {
          const { data } = await queryFulfilled;
          dispatch(setUser(data.user));
        } catch {}
      },
    }),

    signOut: builder.mutation<void, void>({
      query: () => ({ url: "/auth/logout", method: "POST" }),
      onQueryStarted: async (_, { dispatch, queryFulfilled }) => {
        try {
          await queryFulfilled;
        } finally {
          dispatch(clearAuth());
        }
      },
    }),

    getMe: builder.query<AuthUser, void>({
      query: () => "/auth/me",
      transformResponse: (res: { success: boolean; user: AuthUser }) => res.user,
      onQueryStarted: async (_, { dispatch, queryFulfilled }) => {
        try {
          const { data } = await queryFulfilled;
          dispatch(setUser(data));
        } catch {}
      },
    }),
  }),
});

export const {
  useSignInMutation,
  useSignUpMutation,
  useGoogleSignInMutation,
  useSignOutMutation,
  useGetMeQuery,
} = authApi;
