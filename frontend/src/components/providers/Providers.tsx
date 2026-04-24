"use client";

import { ThemeProvider } from "./ThemeProvider";
import { ReduxProvider } from "./ReduxProvider";
import { GoogleOAuthProvider } from "./GoogleOAuthProvider";
import { Toaster } from "@/components/ui/sonner";

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <ReduxProvider>
      <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
        <GoogleOAuthProvider>
          {children}
          <Toaster />
        </GoogleOAuthProvider>
      </ThemeProvider>
    </ReduxProvider>
  );
}
