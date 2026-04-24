import { redirect } from "next/navigation";

export default async function HomePage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  // AuthGuard on /dashboard handles the logged-in check;
  // unauthenticated users land on signin.
  redirect(`/${locale}/dashboard`);
}
