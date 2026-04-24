"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Eye, EyeOff, Loader2 } from "lucide-react";
import { useTranslations } from "next-intl";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { GoogleSSOButton } from "./GoogleSSOButton";
import { AuthDivider } from "./AuthDivider";
import { useSignUpMutation, useGoogleSignInMutation } from "@/store/api/authApi";

const createSignUpSchema = (t: ReturnType<typeof useTranslations<"auth">>) =>
  z.object({
    username: z
      .string()
      .min(2, t("validation.usernameMinLength"))
      .max(50, t("validation.usernameMaxLength"))
      .regex(/^[a-zA-Z0-9_]+$/, t("validation.usernamePattern")),
    email: z
      .string()
      .min(1, t("validation.emailRequired"))
      .email(t("validation.emailInvalid")),
    password: z
      .string()
      .min(8, t("validation.passwordMinLength"))
      .regex(
        /^(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?])/,
        t("validation.passwordPattern")
      ),
  });

type SignUpValues = z.infer<ReturnType<typeof createSignUpSchema>>;

interface SignUpFormProps {
  locale: string;
}

export function SignUpForm({ locale }: SignUpFormProps) {
  const t = useTranslations("auth");
  const tCommon = useTranslations("common");
  const router = useRouter();
  const [showPassword, setShowPassword] = useState(false);

  const [signUp, { isLoading: isSigningUp }] = useSignUpMutation();
  const [googleSignIn, { isLoading: isGoogleSigningIn }] = useGoogleSignInMutation();
  const isLoading = isSigningUp || isGoogleSigningIn;

  const schema = createSignUpSchema(t);
  const form = useForm<SignUpValues>({
    resolver: zodResolver(schema),
    defaultValues: { username: "", email: "", password: "" },
  });

  const getErrorMessage = (errorKey: string) => {
    const knownErrors = [
      "email_already_exists",
      "username_taken",
      "invalid_email",
      "weak_password",
      "generic",
    ] as const;
    const key = knownErrors.find((k) => k === errorKey);
    return key ? t(`errors.${key}`) : t("errors.generic");
  };

  const onSubmit = async (values: SignUpValues) => {
    try {
      await signUp(values).unwrap();
      toast.success(t("signup.title"));
      router.push(`/${locale}/dashboard`);
    } catch (err) {
      const errKey =
        (err as { data?: { error?: string } }).data?.error ?? "generic";
      toast.error(getErrorMessage(errKey));
    }
  };

  const handleGoogleSuccess = async (accessToken: string) => {
    try {
      await googleSignIn({ idToken: accessToken }).unwrap();
      router.push(`/${locale}/dashboard`);
    } catch (err) {
      const errKey =
        (err as { data?: { error?: string } }).data?.error ?? "generic";
      toast.error(getErrorMessage(errKey));
    }
  };

  return (
    <div className="space-y-5">
      <GoogleSSOButton
        onSuccess={handleGoogleSuccess}
        isLoading={isLoading}
        label={t("signup.googleButton")}
      />

      <AuthDivider label={t("divider")} />

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("signup.username")}</FormLabel>
                <FormControl>
                  <Input
                    placeholder={t("signup.usernamePlaceholder")}
                    autoComplete="username"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("signup.email")}</FormLabel>
                <FormControl>
                  <Input
                    placeholder={t("signup.emailPlaceholder")}
                    type="email"
                    autoComplete="email"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("signup.password")}</FormLabel>
                <FormControl>
                  <div className="relative">
                    <Input
                      placeholder={t("signup.passwordPlaceholder")}
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                      disabled={isLoading}
                      className="pr-10"
                      {...field}
                    />
                    <button
                      type="button"
                      onClick={() => setShowPassword((v) => !v)}
                      className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground transition-colors"
                      aria-label={
                        showPassword
                          ? tCommon("hidePassword")
                          : tCommon("showPassword")
                      }
                    >
                      {showPassword ? (
                        <EyeOff className="h-4 w-4" />
                      ) : (
                        <Eye className="h-4 w-4" />
                      )}
                    </button>
                  </div>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button type="submit" className="w-full" disabled={isLoading}>
            {isLoading ? (
              <>
                <Loader2 className="h-4 w-4 animate-spin" />
                {tCommon("loading")}
              </>
            ) : (
              t("signup.submit")
            )}
          </Button>
        </form>
      </Form>
    </div>
  );
}
