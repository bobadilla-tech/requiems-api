import { auth } from "@/lib/auth";
import { headers } from "next/headers";
import { redirect } from "next/navigation";
import { SignupForm } from "@/components/SignupForm";

export const metadata = { title: "Sign Up — Email Validation Demo" };

export default async function SignupPage() {
  // Already signed in → go to dashboard
  const session = await auth.api.getSession({ headers: await headers() });
  if (session) redirect("/dashboard");

  return (
    <div className="min-h-screen flex">

      {/* Left panel — branding / feature list (large screens only) */}
      <div className="hidden lg:flex lg:w-1/2 bg-slate-900 text-white flex-col justify-center px-16 py-12">
        <div className="max-w-md">
          <span className="inline-block bg-blue-500 text-white text-xs font-bold px-3 py-1 rounded-full uppercase tracking-wide mb-6">
            Tutorial Demo
          </span>
          <h1 className="text-4xl font-bold leading-tight mb-4">
            Email Validation<br />Done Right
          </h1>
          <p className="text-slate-400 text-lg leading-relaxed mb-8">
            Watch the Requiems API validate emails in real time — catching typos,
            blocking disposable addresses, and verifying DNS records before your
            users even hit submit.
          </p>

          <ul className="space-y-3">
            {[
              ["Syntax check",     "RFC 5322 format validation"],
              ["MX record lookup", "Live DNS — can this domain receive mail?"],
              ["Disposable block", "90,000+ blocked temporary domains"],
              ["Typo suggestion",  "gmial.com → gmail.com"],
              ["Normalization",    "User+Tag@Gmail.COM → user@gmail.com"],
            ].map(([title, desc]) => (
              <li key={title} className="flex items-start gap-3">
                <div className="mt-1 h-4 w-4 flex-shrink-0 rounded-full bg-emerald-500 flex items-center justify-center">
                  <svg className="h-2.5 w-2.5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={3}>
                    <path strokeLinecap="round" strokeLinejoin="round" d="M4.5 12.75l6 6 9-13.5" />
                  </svg>
                </div>
                <div>
                  <span className="text-sm font-semibold text-white">{title}</span>
                  <span className="text-sm text-slate-400"> — {desc}</span>
                </div>
              </li>
            ))}
          </ul>

          <div className="mt-10 pt-8 border-t border-slate-700">
            <p className="text-sm text-slate-500">Powered by</p>
            <p className="text-lg font-semibold text-white mt-1">Requiems API</p>
            <p className="text-sm text-slate-400">requiems.xyz</p>
          </div>
        </div>
      </div>

      {/* Right panel — form */}
      <div className="flex-1 flex items-center justify-center px-6 py-12">
        <div className="w-full max-w-md">
          <div className="mb-8">
            <h2 className="text-2xl font-bold text-foreground">Get started</h2>
            <p className="text-muted-foreground mt-1">
              Type an email — watch it validate in real time as you type.
            </p>
          </div>

          <SignupForm />

          {/* Demo hint for tutorial viewers */}
          <div className="mt-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
            <p className="text-xs font-semibold text-blue-700 uppercase tracking-wide mb-2">
              Try these test emails
            </p>
            <div className="space-y-1 text-xs text-blue-800">
              <p><code className="font-mono">user@gmial.com</code> → typo suggestion</p>
              <p><code className="font-mono">test@mailinator.com</code> → disposable blocked</p>
              <p><code className="font-mono">not-an-email</code> → syntax invalid</p>
              <p><code className="font-mono">User+Tag@Googlemail.COM</code> → normalization demo</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
