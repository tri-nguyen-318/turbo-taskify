import { getTranslations } from "next-intl/server";
import Link from "next/link";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { SignInForm } from "@/components/auth/SignInForm";

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: "auth.signin" });
  return { title: `${t("title")} — Turbo Taskify` };
}

export default async function SignInPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: "auth" });

  return (
    <Card className="shadow-lg border-border/60">
      <CardHeader className="space-y-1 pb-4">
        <CardTitle className="text-2xl font-bold text-center">
          {t("signin.title")}
        </CardTitle>
        <CardDescription className="text-center">
          {t("signin.subtitle")}
        </CardDescription>
      </CardHeader>

      <CardContent className="space-y-5">
        <SignInForm locale={locale} />

        <p className="text-center text-sm text-muted-foreground">
          {t("signin.noAccount")}{" "}
          <Link
            href={`/${locale}/signup`}
            className="font-medium text-primary hover:underline underline-offset-4"
          >
            {t("signin.signupLink")}
          </Link>
        </p>
      </CardContent>
    </Card>
  );
}
