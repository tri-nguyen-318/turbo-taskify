"use client";

import { GoogleOAuthProvider as Provider } from "@react-oauth/google";

const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID ?? "";

export const googleOAuthEnabled = clientId.length > 0;

export function GoogleOAuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  if (!googleOAuthEnabled) {
    return <>{children}</>;
  }

  return <Provider clientId={clientId}>{children}</Provider>;
}
