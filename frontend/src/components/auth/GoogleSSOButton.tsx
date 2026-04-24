"use client";

import { GoogleLogin } from "@react-oauth/google";
import { useTranslations } from "next-intl";
import { useTheme } from "next-themes";
import { googleOAuthEnabled } from "@/components/providers/GoogleOAuthProvider";

interface GoogleSSOButtonProps {
  onSuccess: (idToken: string) => void;
  isLoading?: boolean;
  label?: string;
}

export function GoogleSSOButton({
  onSuccess,
  isLoading,
  label,
}: GoogleSSOButtonProps) {
  const t = useTranslations("auth");
  const { resolvedTheme } = useTheme();

  if (!googleOAuthEnabled) return null;

  return (
    <div className={isLoading ? "pointer-events-none opacity-50" : undefined}>
      <GoogleLogin
        onSuccess={(credentialResponse) => {
          if (credentialResponse.credential) {
            onSuccess(credentialResponse.credential);
          }
        }}
        onError={() => {
          console.error("Google login failed");
        }}
        text={label === t("signup.googleButton") ? "signup_with" : "signin_with"}
        theme={resolvedTheme === "dark" ? "filled_black" : "outline"}
        shape="rectangular"
        width="100%"
      />
    </div>
  );
}
