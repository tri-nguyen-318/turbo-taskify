"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Loader2 } from "lucide-react";
import { useGetMeQuery } from "@/store/api/authApi";

interface AuthGuardProps {
  children: React.ReactNode;
  locale: string;
}

export function AuthGuard({ children, locale }: AuthGuardProps) {
  const { isLoading, isError } = useGetMeQuery();
  const router = useRouter();

  useEffect(() => {
    if (!isLoading && isError) {
      router.replace(`/${locale}/signin`);
    }
  }, [isLoading, isError, locale, router]);

  if (isLoading || isError) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-background">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    );
  }

  return <>{children}</>;
}
