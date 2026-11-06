"use client";

import { useState, useEffect, useCallback, useRef } from "react";
import { useRouter } from "next/navigation";
import { signIn, signUp } from "@/lib/auth-client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { EmailValidationCard } from "@/components/EmailValidationCard";
import type { EmailValidationData } from "@/lib/requiems";

type ValidationState =
  | { status: "idle" }
  | { status: "loading" }
  | { status: "done"; data: EmailValidationData }
  | { status: "error"; message: string };

const DEBOUNCE_MS = 600;

type Mode = "signup" | "signin";

export function SignupForm() {
  const router = useRouter();
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const [mode, setMode] = useState<Mode>("signup");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [validation, setValidation] = useState<ValidationState>({ status: "idle" });
  const [submitting, setSubmitting] = useState(false);
  const [authError, setAuthError] = useState<string | null>(null);

  // ── Debounced real-time validation ──────────────────────────────────────────

  const runValidation = useCallback(async (value: string) => {
    if (!value.trim()) {
      setValidation({ status: "idle" });
      return;
    }

    setValidation({ status: "loading" });

    try {
      const res = await fetch("/api/validate-email", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email: value.trim() }),
      });
      const json = await res.json();

      if (!res.ok) {
        setValidation({ status: "error", message: json.error ?? "Validation failed" });
        return;
      }

      setValidation({ status: "done", data: json.data });
    } catch {
      setValidation({ status: "error", message: "Could not reach the validation service" });
    }
  }, []);

  useEffect(() => {
    if (timerRef.current) clearTimeout(timerRef.current);
    timerRef.current = setTimeout(() => runValidation(email), DEBOUNCE_MS);
    return () => {
      if (timerRef.current) clearTimeout(timerRef.current);
    };
  }, [email, runValidation]);

  // ── Accept typo suggestion ───────────────────────────────────────────────────

  const acceptSuggestion = () => {
    if (validation.status !== "done" || !validation.data.suggestion) return;
    const [localPart] = email.split("@");
    setEmail(`${localPart}@${validation.data.suggestion}`);
  };

  // ── Form submit ──────────────────────────────────────────────────────────────

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setAuthError(null);

    // Client-side guard: block if real-time check already shows invalid
    if (validation.status === "done" && !validation.data.valid) {
      setAuthError("Please fix the email address before continuing.");
      return;
    }

    setSubmitting(true);

    if (mode === "signup") {
      const { error } = await signUp.email({
        name,
        email,
        password,
      });

      if (error) {
        setAuthError(error.message ?? "Sign-up failed. Please try again.");
        setSubmitting(false);
        return;
      }
    } else {
      const { error } = await signIn.email({
        email,
        password,
      });

      if (error) {
        setAuthError(error.message ?? "Sign-in failed. Please try again.");
        setSubmitting(false);
        return;
      }
    }

    router.push("/dashboard");
  };

  // ── Derived UI state ─────────────────────────────────────────────────────────

  const inputBorderClass =
    validation.status === "done"
      ? validation.data.valid
        ? "border-emerald-400 focus-visible:ring-emerald-300"
        : "border-red-400 focus-visible:ring-red-300"
      : "";

  const isBlocked =
    submitting ||
    validation.status === "loading" ||
    (validation.status === "done" && !validation.data.valid);

  const suggestion =
    validation.status === "done" && validation.data.suggestion
      ? email.split("@")[0] + "@" + validation.data.suggestion
      : null;

  return (
    <div className="space-y-6">
      <form onSubmit={handleSubmit} className="space-y-4">

        {/* Mode toggle */}
        <div className="flex rounded-lg border border-border overflow-hidden">
          {(["signup", "signin"] as Mode[]).map((m) => (
            <button
              key={m}
              type="button"
              onClick={() => { setMode(m); setAuthError(null); }}
              className={`flex-1 py-2 text-sm font-medium transition-colors ${
                mode === m
                  ? "bg-primary text-primary-foreground"
                  : "bg-background text-muted-foreground hover:bg-muted"
              }`}
            >
              {m === "signup" ? "Create account" : "Sign in"}
            </button>
          ))}
        </div>

        {/* Name (sign-up only) */}
        {mode === "signup" && (
          <div className="space-y-2">
            <Label htmlFor="name">Full name</Label>
            <Input
              id="name"
              type="text"
              placeholder="Jane Smith"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
        )}

        {/* Email */}
        <div className="space-y-2">
          <Label htmlFor="email">Email address</Label>
          <div className="relative">
            <Input
              id="email"
              type="email"
              placeholder="you@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className={inputBorderClass}
              required
            />
            {validation.status === "loading" && (
              <div className="absolute right-3 top-1/2 -translate-y-1/2">
                <div className="h-4 w-4 animate-spin rounded-full border-2 border-muted border-t-primary" />
              </div>
            )}
          </div>

          {/* Typo suggestion click-to-correct */}
          {suggestion && (
            <p className="text-sm text-amber-700">
              Did you mean{" "}
              <button
                type="button"
                onClick={acceptSuggestion}
                className="font-semibold underline underline-offset-2 hover:text-amber-900"
              >
                {suggestion}
              </button>
              ?
            </p>
          )}

          {/* Proxy error (network/API failure) */}
          {validation.status === "error" && (
            <p className="text-sm text-red-600">{validation.message}</p>
          )}
        </div>

        {/* Password */}
        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            type="password"
            placeholder="••••••••"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            minLength={8}
            required
          />
        </div>

        {/* Auth error from Better Auth hooks (server-side enforcement) */}
        {authError && (
          <p className="text-sm text-red-600 bg-red-50 border border-red-200 rounded-md px-3 py-2">
            {authError}
          </p>
        )}

        <Button type="submit" disabled={isBlocked} className="w-full">
          {submitting
            ? mode === "signup" ? "Creating account…" : "Signing in…"
            : mode === "signup" ? "Create account" : "Sign in"}
        </Button>
      </form>

      {/* Validation card — appears once the API responds */}
      {validation.status === "done" && (
        <div className="animate-in fade-in slide-in-from-top-2 duration-300">
          <EmailValidationCard
            data={validation.data}
          />
        </div>
      )}
    </div>
  );
}
