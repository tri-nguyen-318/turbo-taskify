import { Separator } from "@/components/ui/separator";

interface AuthDividerProps {
  label: string;
}

export function AuthDivider({ label }: AuthDividerProps) {
  return (
    <div className="relative flex items-center gap-3">
      <Separator className="flex-1" />
      <span className="text-xs text-muted-foreground whitespace-nowrap">
        {label}
      </span>
      <Separator className="flex-1" />
    </div>
  );
}
