"use client";

import { Moon, Sun, Zap } from "lucide-react";
import { useTranslations } from "next-intl";
import { useTheme } from "next-themes";
import { AuthGuard } from "@/components/common/AuthGuard";
import { SignOutButton } from "@/components/auth/SignOutButton";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useAppSelector } from "@/store/hooks";

interface DashboardViewProps {
  locale: string;
}

export function DashboardView({ locale }: DashboardViewProps) {
  const t = useTranslations("dashboard");
  const tAuth = useTranslations("auth");
  const { user } = useAppSelector((s) => s.auth);
  const { resolvedTheme, setTheme } = useTheme();

  return (
    <AuthGuard locale={locale}>
      <div className="min-h-screen flex flex-col bg-background">
        {/* Navbar */}
        <header className="sticky top-0 z-10 border-b border-border bg-background/80 backdrop-blur-sm">
          <div className="max-w-5xl mx-auto px-4 sm:px-6 h-14 flex items-center justify-between">
            <div className="flex items-center gap-2 text-primary font-bold text-lg">
              <Zap className="h-5 w-5" />
              {tAuth("appName")}
            </div>

            <div className="flex items-center gap-2">
              <Button
                variant="ghost"
                size="icon"
                aria-label="Toggle theme"
                onClick={() =>
                  setTheme(resolvedTheme === "dark" ? "light" : "dark")
                }
              >
                {resolvedTheme === "dark" ? (
                  <Sun className="h-4 w-4" />
                ) : (
                  <Moon className="h-4 w-4" />
                )}
              </Button>

              <SignOutButton locale={locale} />
            </div>
          </div>
        </header>

        {/* Main content */}
        <main className="flex-1 max-w-5xl mx-auto w-full px-4 sm:px-6 py-10 space-y-8">
          {/* Greeting */}
          <div>
            <h1 className="text-3xl font-bold text-foreground">
              {t("greeting", { name: user?.username ?? "…" })}
            </h1>
            <p className="mt-1 text-muted-foreground">{t("tagline")}</p>
          </div>

          {/* Account card */}
          <Card className="max-w-sm">
            <CardHeader>
              <CardTitle className="text-base">{t("accountInfo")}</CardTitle>
              <CardDescription>{user?.email}</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              <Row label={t("username")} value={user?.username} />
              <Row label={t("email")} value={user?.email} />
              <Row label={t("theme")} value={user?.theme} />
              <Row label={t("language")} value={user?.language} />
            </CardContent>
          </Card>
        </main>
      </div>
    </AuthGuard>
  );
}

function Row({
  label,
  value,
}: {
  label: string;
  value: string | undefined;
}) {
  return (
    <div className="flex items-center justify-between gap-4">
      <span className="text-muted-foreground">{label}</span>
      <span className="font-medium text-right truncate">{value ?? "—"}</span>
    </div>
  );
}
