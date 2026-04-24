import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { AuthUser } from "@/lib/api";

interface AuthState {
  user: AuthUser | null;
}

const initialState: AuthState = {
  user: null,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setUser(state, action: PayloadAction<AuthUser>) {
      state.user = action.payload;
    },
    clearAuth(state) {
      state.user = null;
    },
  },
});

export const { setUser, clearAuth } = authSlice.actions;
export default authSlice.reducer;
