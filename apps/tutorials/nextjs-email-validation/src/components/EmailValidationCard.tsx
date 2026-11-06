"use client";

import type { EmailValidationData } from "@/lib/requiems";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Alert, AlertDescription } from "@/components/ui/alert";

// ── Inline SVG icons (no icon library dependency) ───────────────────────────

function CheckIcon() {
  return (
    <svg className="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M4.5 12.75l6 6 9-13.5" />
    </svg>
  );
}

function CrossIcon() {
  return (
    <svg className="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
      <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
    </svg>
  );
}

function WarnIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
      <path strokeLinecap="round" strokeLinejoin="round"
        d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z" />
    </svg>
  );
}

// ── Boolean field row ────────────────────────────────────────────────────────

type BoolRowProps = {
  label: string;
  value: boolean;
  description: string;
  /** When true, the boolean TRUE is the BAD case (e.g. disposable) */
  invertColor?: boolean;
};

function BoolRow({ label, value, description, invertColor }: BoolRowProps) {
  const isGood = invertColor ? !value : value;

  return (
    <div className="flex items-start gap-3 py-2 border-b border-border last:border-0">
      <div
        className={`mt-0.5 h-5 w-5 flex-shrink-0 rounded-full flex items-center justify-center ${
          isGood ? "bg-emerald-100 text-emerald-600" : "bg-red-100 text-red-500"
        }`}
      >
        {isGood ? <CheckIcon /> : <CrossIcon />}
      </div>
      <div className="flex-1 min-w-0">
        <div className="flex items-center gap-2 flex-wrap">
          <code className="text-xs font-mono font-medium text-foreground">{label}</code>
          <Badge
            variant={isGood ? "default" : "destructive"}
            className={`text-xs font-mono ${isGood ? "bg-emerald-500 hover:bg-emerald-500" : ""}`}
          >
            {String(value)}
          </Badge>
        </div>
        <p className="text-xs text-muted-foreground mt-0.5">{description}</p>
      </div>
    </div>
  );
}

// ── String field row ─────────────────────────────────────────────────────────

type StringRowProps = {
  label: string;
  value: string | null;
  description: string;
};

function StringRow({ label, value, description }: StringRowProps) {
  const hasValue = value !== null && value !== "";
  return (
    <div className="flex items-start gap-3 py-2 border-b border-border last:border-0">
      <div
        className={`mt-0.5 h-5 w-5 flex-shrink-0 rounded-full flex items-center justify-center text-[9px] font-bold ${
          hasValue ? "bg-blue-100 text-blue-600" : "bg-muted text-muted-foreground"
        }`}
      >
        Ab
      </div>
      <div className="flex-1 min-w-0">
        <div className="flex items-center gap-2 flex-wrap">
          <code className="text-xs font-mono font-medium text-foreground">{label}</code>
          <code
            className={`text-xs px-1.5 py-0.5 rounded font-mono ${
              hasValue ? "bg-blue-50 text-blue-700" : "bg-muted text-muted-foreground"
            }`}
          >
            {hasValue ? value : "null"}
          </code>
        </div>
        <p className="text-xs text-muted-foreground mt-0.5">{description}</p>
      </div>
    </div>
  );
}

// ── Main card ────────────────────────────────────────────────────────────────

type Props = {
  data: EmailValidationData;
  isLoading?: boolean;
};

export function EmailValidationCard({ data, isLoading }: Props) {
  // Reconstruct the suggested full email from domain suggestion
  const suggestedEmail = data.suggestion
    ? data.email.split("@")[0] + "@" + data.suggestion
    : null;

  return (
    <Card className={`w-full transition-opacity duration-200 ${isLoading ? "opacity-50" : "opacity-100"}`}>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <CardTitle className="text-base font-semibold">Validation Result</CardTitle>
          <Badge
            variant={data.valid ? "default" : "destructive"}
            className={`text-sm px-3 py-1 ${data.valid ? "bg-emerald-500 hover:bg-emerald-500" : ""}`}
          >
            {data.valid ? "✓ Valid" : "✗ Invalid"}
          </Badge>
        </div>
      </CardHeader>

      <CardContent className="space-y-1">
        {/* Typo suggestion — prominent amber alert */}
        {suggestedEmail && (
          <Alert className="mb-4 border-amber-200 bg-amber-50">
            <WarnIcon className="h-4 w-4 text-amber-600 flex-shrink-0" />
            <AlertDescription className="text-amber-800 text-sm ml-2">
              Did you mean <strong>{suggestedEmail}</strong>?
            </AlertDescription>
          </Alert>
        )}

        {/* Disposable email warning */}
        {data.disposable && (
          <Alert className="mb-4 border-red-200 bg-red-50">
            <WarnIcon className="h-4 w-4 text-red-600 flex-shrink-0" />
            <AlertDescription className="text-red-800 text-sm ml-2">
              This looks like a temporary / disposable email address.
            </AlertDescription>
          </Alert>
        )}

        {/* Boolean fields */}
        <BoolRow
          label="syntax_valid"
          value={data.syntax_valid}
          description="RFC 5322 email format check"
        />
        <BoolRow
          label="mx_valid"
          value={data.mx_valid}
          description="Domain has live MX DNS records (can receive mail)"
        />
        <BoolRow
          label="disposable"
          value={data.disposable}
          invertColor
          description="Domain found in 90,000+ disposable address blocklist"
        />
        <BoolRow
          label="valid"
          value={data.valid}
          description="Overall: true only when syntax_valid AND mx_valid are both true"
        />

        {/* String fields */}
        <StringRow
          label="normalized"
          value={data.normalized}
          description="Canonical address (lowercase, tags stripped, dots removed)"
        />
        <StringRow
          label="domain"
          value={data.domain}
          description="Domain portion extracted from the address"
        />
        <StringRow
          label="suggestion"
          value={data.suggestion}
          description="Closest known provider if input looks like a typo (Levenshtein ≤ 2)"
        />
        <StringRow
          label="email"
          value={data.email}
          description="Exact input you supplied — never modified"
        />
      </CardContent>
    </Card>
  );
}
