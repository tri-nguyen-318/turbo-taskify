"use client";

import { Loader2, LogOut } from "lucide-react";
import { useTranslations } from "next-intl";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { useSignOutMutation } from "@/store/api/authApi";

interface SignOutButtonProps {
  locale: string;
  variant?: "default" | "ghost" | "outline";
}

export function SignOutButton({ locale, variant = "ghost" }: SignOutButtonProps) {
  const t = useTranslations("auth");
  const router = useRouter();
  const [signOut, { isLoading }] = useSignOutMutation();

  const handleSignOut = async () => {
    await signOut();
    router.push(`/${locale}/signin`);
  };

  return (
    <Button
      variant={variant}
      size="sm"
      onClick={handleSignOut}
      disabled={isLoading}
      className="gap-2"
    >
      {isLoading ? (
        <Loader2 className="h-4 w-4 animate-spin" />
      ) : (
        <LogOut className="h-4 w-4" />
      )}
      {t("signout.button")}
    </Button>
  );
}
