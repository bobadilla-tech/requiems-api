"use client";

import { signOut } from "@/lib/auth-client";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";

export function SignOut() {
  const router = useRouter();

  return (
    <Button
      variant="outline"
      className="w-full"
      onClick={async () => {
        await signOut();
        router.push("/signup");
      }}
    >
      Sign out
    </Button>
  );
}
