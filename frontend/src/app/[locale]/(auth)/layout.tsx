import { Zap } from "lucide-react";
import { getTranslations } from "next-intl/server";
import Link from "next/link";

export default async function AuthLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: "auth" });

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="px-6 py-4">
        <Link
          href={`/${locale}`}
          className="inline-flex items-center gap-2 text-primary font-bold text-lg hover:opacity-80 transition-opacity"
        >
          <Zap className="h-5 w-5" />
          {t("appName")}
        </Link>
      </header>

      <main className="flex-1 flex items-center justify-center px-4 py-12">
        <div className="w-full max-w-md">{children}</div>
      </main>

      <footer className="px-6 py-4 text-center text-xs text-muted-foreground">
        © {new Date().getFullYear()} {t("appName")}
      </footer>
    </div>
  );
}
