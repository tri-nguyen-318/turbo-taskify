import { getTranslations } from "next-intl/server";
import Link from "next/link";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { SignUpForm } from "@/components/auth/SignUpForm";

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: "auth.signup" });
  return { title: `${t("title")} — Turbo Taskify` };
}

export default async function SignUpPage({
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
          {t("signup.title")}
        </CardTitle>
        <CardDescription className="text-center">
          {t("signup.subtitle")}
        </CardDescription>
      </CardHeader>

      <CardContent className="space-y-5">
        <SignUpForm locale={locale} />

        <p className="text-center text-sm text-muted-foreground">
          {t("signup.hasAccount")}{" "}
          <Link
            href={`/${locale}/signin`}
            className="font-medium text-primary hover:underline underline-offset-4"
          >
            {t("signup.signinLink")}
          </Link>
        </p>
      </CardContent>
    </Card>
  );
}
