import { auth } from "@/lib/auth";
import { headers } from "next/headers";
import { redirect } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { SignOut } from "@/components/SignOut";

export const metadata = { title: "Dashboard — Email Validation Demo" };

export default async function DashboardPage() {
  const session = await auth.api.getSession({ headers: await headers() });
  if (!session) redirect("/signup");

  return (
    <div className="min-h-screen bg-muted/40 flex items-center justify-center p-6">
      <div className="max-w-lg w-full space-y-6">

        {/* Success banner */}
        <Card className="border-emerald-200 bg-emerald-50">
          <CardHeader>
            <CardTitle className="text-emerald-800 flex items-center gap-2 text-base">
              <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
                <path strokeLinecap="round" strokeLinejoin="round" d="M4.5 12.75l6 6 9-13.5" />
              </svg>
              Sign-in successful!
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-emerald-700 text-sm">
              The Requiems API validated your email server-side inside the Better Auth{" "}
              <code className="bg-emerald-100 px-1 rounded font-mono text-xs">hooks.before</code>{" "}
              handler before this session was created. Disposable addresses and bad
              domains are rejected even if someone bypasses the client-side check.
            </p>
          </CardContent>
        </Card>

        {/* Session info */}
        <Card>
          <CardHeader>
            <CardTitle>Your session</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <p className="text-xs text-muted-foreground uppercase tracking-wide font-semibold mb-1">Name</p>
              <p className="text-foreground font-medium">{session.user.name}</p>
            </div>
            <div>
              <p className="text-xs text-muted-foreground uppercase tracking-wide font-semibold mb-1">Email</p>
              <p className="text-foreground font-mono text-sm">{session.user.email}</p>
            </div>
            <div>
              <p className="text-xs text-muted-foreground uppercase tracking-wide font-semibold mb-1">User ID</p>
              <p className="text-muted-foreground font-mono text-xs">{session.user.id}</p>
            </div>
          </CardContent>
        </Card>

        <SignOut />
      </div>
    </div>
  );
}
