import PDFDocument from "pdfkit";
import { createWriteStream, mkdirSync } from "fs";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const OUT_DIR = resolve(__dirname, "../apps/dashboard/public/documents");

mkdirSync(OUT_DIR, { recursive: true });

const COMPANY = "Bobadilla Technologies";
const PRODUCT = "Requiems API";
const CONTACT = "eliaz@bobadilla.tech";
const EFFECTIVE_DATE = "February 1, 2026";
const EFFECTIVE_DATE_LEGAL = "2026-02-01";

// ─── helpers ─────────────────────────────────────────────────────────────────

function doc(filename) {
  const pdf = new PDFDocument({ margin: 72, size: "LETTER", bufferPages: true });
  const stream = createWriteStream(resolve(OUT_DIR, filename));
  stream.on("error", (err) => {
    console.error(`Failed to write ${filename}:`, err);
    process.exit(1);
  });
  pdf.pipe(stream);
  return pdf;
}

function finalize(pdf, filename) {
  footer(pdf);
  pdf.flushPages();
  pdf.end();
  console.log(`✓ ${filename}`);
}

// Like moveDown but skips if it would push past the bottom margin (prevents trailing blank pages)
function safeDown(pdf, lines) {
  const bottom = pdf.page.height - pdf.page.margins.bottom;
  const step = pdf.currentLineHeight(true) * lines;
  if (pdf.y + step < bottom) {
    pdf.moveDown(lines);
  }
}

function header(pdf, title, subtitle) {
  pdf
    .fontSize(9)
    .fillColor("#6b7280")
    .text(`${COMPANY} · ${PRODUCT} · ${CONTACT}`, { align: "right" });

  pdf.moveDown(0.5);
  pdf
    .moveTo(pdf.page.margins.left, pdf.y)
    .lineTo(pdf.page.width - pdf.page.margins.right, pdf.y)
    .strokeColor("#e5e7eb")
    .stroke();
  pdf.moveDown(1.5);

  pdf
    .fontSize(26)
    .fillColor("#111827")
    .font("Helvetica-Bold")
    .text(title);

  if (subtitle) {
    pdf.moveDown(0.4);
    pdf
      .fontSize(13)
      .fillColor("#6b7280")
      .font("Helvetica")
      .text(subtitle);
  }

  pdf.moveDown(0.5);
  pdf
    .moveTo(pdf.page.margins.left, pdf.y)
    .lineTo(pdf.page.width - pdf.page.margins.right, pdf.y)
    .strokeColor("#e5e7eb")
    .stroke();
  pdf.moveDown(1.5);
}

function section(pdf, title) {
  pdf.moveDown(0.8);
  pdf
    .fontSize(13)
    .fillColor("#1f2937")
    .font("Helvetica-Bold")
    .text(title);
  pdf.moveDown(0.4);
  pdf.font("Helvetica").fillColor("#374151").fontSize(10.5);
}

function body(pdf, text) {
  pdf
    .fontSize(10.5)
    .fillColor("#374151")
    .font("Helvetica")
    .text(text, { lineGap: 3 });
  safeDown(pdf, 0.5);
}

function bullet(pdf, items) {
  items.forEach((item) => {
    pdf
      .fontSize(10.5)
      .fillColor("#374151")
      .font("Helvetica")
      .text(`• ${item}`, { indent: 12, lineGap: 2 });
  });
  safeDown(pdf, 0.5);
}

function footer(pdf) {
  const range = pdf.bufferedPageRange();
  for (let i = range.start; i < range.start + range.count; i++) {
    pdf.switchToPage(i);
    // footer Y (height-40 = 752) is past maxY (height-margins.bottom = 720).
    // Temporarily remove the bottom margin constraint so pdf.text() doesn't
    // trigger _nextSection() and create a new overflow page for every page.
    const savedBottom = pdf.page.margins.bottom;
    pdf.page.margins.bottom = 0;
    pdf
      .fontSize(8)
      .fillColor("#9ca3af")
      .text(
        `${PRODUCT} Compliance Documentation · Effective ${EFFECTIVE_DATE} · ${CONTACT}`,
        pdf.page.margins.left,
        pdf.page.height - 40,
        { align: "center", width: pdf.page.width - pdf.page.margins.left - pdf.page.margins.right, lineBreak: false }
      );
    pdf.page.margins.bottom = savedBottom;
  }
}

// ─── 1. Security & Architecture Packet ───────────────────────────────────────

function generateSecurityPacket() {
  const pdf = doc("security-packet.pdf");

  header(
    pdf,
    `${PRODUCT} — Security & Architecture Packet`,
    "For enterprise security reviews and procurement teams"
  );

  section(pdf, "1. Core Security Model — Stateless Architecture");
  body(
    pdf,
    `${PRODUCT} is engineered as a zero-trust, stateless processing engine. Every API ` +
    `request is processed entirely in volatile server memory (RAM). The moment your HTTP ` +
    `response is delivered, the payload is permanently purged. We have no mechanism to ` +
    `recover submitted data because it was never written to disk, a database, or any ` +
    `persistent medium.`
  );
  bullet(pdf, [
    "In-memory execution — payloads never touch disk or any persistent layer",
    "Zero data retention — no database of user API inputs or outputs exists",
    "No request/response cache or logging of payload content",
    "Anonymized telemetry only — HTTP status codes and timestamps, never payload content",
  ]);

  section(pdf, "2. Transparency — The Counter API Exception");
  body(
    pdf,
    `Our Counter API endpoint (/v1/technology/counter/{namespace}) is the only endpoint ` +
    `that persists data beyond a single request. It stores a named integer count value ` +
    `(e.g., namespace "page-views" maps to an integer). No request payload, no user ` +
    `content, and no identifying data is stored — only the counter value itself. ` +
    `This endpoint is rarely used and can be avoided entirely if zero-persistence is ` +
    `required by your organization.`
  );

  section(pdf, "3. Network Perimeter — Cloudflare");
  body(pdf, "All inbound traffic is routed through Cloudflare before reaching our origin servers.");
  bullet(pdf, [
    "Enterprise-grade DDoS mitigation at the edge",
    "Web Application Firewall (WAF) filtering on all requests",
    "TLS 1.3 enforced — no downgrade to older protocol versions",
    "Rate limiting and abuse prevention at the network layer",
    "IP reputation filtering",
  ]);

  section(pdf, "4. Physical Infrastructure — Hetzner (EU)");
  body(
    pdf,
    `Our custom datasets and processing code are hosted on servers in European Union ` +
    `data centers operated by Hetzner Online GmbH, an ISO/IEC 27001-certified provider.`
  );
  bullet(pdf, [
    "All data centers located in the European Union",
    "ISO/IEC 27001 certified facilities (physical security, environmental controls)",
    "Strict physical access controls and 24/7 CCTV monitoring",
    "Redundant power systems and network connectivity",
    "Hetzner's ISO 27001 certificate available upon request",
  ]);

  section(pdf, "5. API Authentication & Access Control");
  bullet(pdf, [
    "High-entropy API keys issued per account — cryptographically random, not guessable",
    "Instant key revocation available via user dashboard",
    "TLS 1.3 required on every API request — keys transmitted only over encrypted channels",
    "Rate limiting enforced at the Cloudflare edge layer",
    "X-Backend-Secret header validation at the origin for Cloudflare-bypass protection",
  ]);

  section(pdf, "6. Compliance Scope Alignment");
  body(
    pdf,
    `Because ${PRODUCT} maintains a 100% stateless architecture for user payloads, a ` +
    `traditional SOC 2 data-retention audit does not apply to our operational footprint. ` +
    `We satisfy the core Security and Confidentiality principles of SOC 2 by design — ` +
    `never retaining data at rest. Physical infrastructure compliance is inherited via ` +
    `Hetzner's ISO 27001 certification. Network compliance is inherited via Cloudflare's ` +
    `enterprise security controls.`
  );
  bullet(pdf, [
    "GDPR: No user input stored or profiled. EU-only infrastructure. Data minimization is structural.",
    "CCPA: No sale or sharing of personal data. No user input database exists to share.",
    "SOC 2 Security/Confidentiality: Satisfied by never retaining data at rest.",
    "Physical Compliance: Inherited from Hetzner ISO 27001 certification.",
  ]);

  section(pdf, "7. Incident Response");
  body(
    pdf,
    `In the event of a security incident, we commit to notifying affected clients within ` +
    `72 hours of discovery. Contact: ${CONTACT}`
  );

  finalize(pdf, "security-packet.pdf");
}

// ─── 2. Data Processing Agreement ────────────────────────────────────────────

function generateDPA() {
  const pdf = doc("dpa-template.pdf");

  header(
    pdf,
    `Data Processing Agreement`,
    `${PRODUCT} — Template · Effective ${EFFECTIVE_DATE}`
  );

  body(
    pdf,
    `This Data Processing Agreement ("DPA") is entered into between:\n\n` +
    `Data Controller ("Client"): [CLIENT COMPANY NAME], [CLIENT ADDRESS]\n\n` +
    `Data Processor ("Processor"): ${COMPANY}, operated by Eliaz Bobadilla, ` +
    `contact: ${CONTACT}, providing the ${PRODUCT} service.\n\n` +
    `This DPA supplements and forms part of the Master Subscription Agreement or ` +
    `Terms of Service between the parties. In the event of conflict, this DPA prevails ` +
    `with respect to data processing matters.`
  );

  section(pdf, "1. Definitions");
  body(pdf,
    '"Personal Data" means any information relating to an identified or identifiable natural person.\n\n' +
    '"Processing" means any operation performed on Personal Data.\n\n' +
    '"Controller" means the entity that determines the purposes and means of processing Personal Data.\n\n' +
    '"Processor" means the entity that processes Personal Data on behalf of the Controller.\n\n' +
    '"Sub-processor" means any third party engaged by the Processor to process Personal Data.'
  );

  section(pdf, "2. Subject Matter and Duration");
  body(pdf,
    `The Processor provides API-based data processing services as described in the ` +
    `service agreement. This DPA remains in effect for the duration of the service ` +
    `agreement and terminates automatically upon its expiration or termination.`
  );

  section(pdf, "3. Nature and Purpose of Processing");
  bullet(pdf, [
    "Purpose: Delivery of API responses as requested by the Controller",
    "Type of Personal Data: Temporary data payloads submitted via API endpoint",
    "Categories of Data Subjects: End users of the Controller's applications, if applicable",
    "Retention Period: 0 seconds — data is purged instantly from volatile server memory upon HTTP response delivery",
  ]);

  body(pdf,
    `IMPORTANT: The Processor's architecture is stateless. Submitted payloads are ` +
    `processed entirely in volatile server memory (RAM) and are permanently purged ` +
    `the moment the HTTP response is delivered. The Processor maintains no database, ` +
    `cache, or log of submitted payload content. The Processor cannot recover, produce, ` +
    `or disclose submitted data because it does not persist beyond the request lifecycle.`
  );

  section(pdf, "4. Processor Obligations");
  bullet(pdf, [
    "Process Personal Data only on documented instructions from the Controller",
    "Ensure that persons authorized to process data are bound by confidentiality",
    "Implement appropriate technical and organizational security measures",
    "Assist the Controller in ensuring compliance with GDPR Articles 32-36",
    "Delete or return all Personal Data upon termination of services",
    "Provide all information necessary to demonstrate compliance with this DPA",
  ]);

  section(pdf, "5. Sub-processors");
  body(pdf, "The Processor engages the following sub-processors in delivering the service:");
  bullet(pdf, [
    "Hetzner Online GmbH (Germany/EU) — Compute infrastructure. ISO/IEC 27001 certified.",
    "Cloudflare, Inc. (USA, EU data routing) — Network perimeter, TLS termination, DDoS protection.",
  ]);
  body(pdf,
    `The Processor will notify the Controller of any intended changes to sub-processors ` +
    `with reasonable advance notice, providing the Controller the opportunity to object.`
  );

  section(pdf, "6. Security Measures");
  bullet(pdf, [
    "TLS 1.3 encryption for all data in transit",
    "Zero data retention for API payloads — volatile memory processing only",
    "High-entropy API key authentication per account",
    "DDoS mitigation and WAF protection via Cloudflare",
    "Physical security via Hetzner ISO 27001 certified facilities",
    "Access controls limiting server access to authorized personnel only",
  ]);

  section(pdf, "7. Data Breach Notification");
  body(pdf,
    `In the event of a personal data breach, the Processor shall notify the Controller ` +
    `without undue delay and, where feasible, within 72 hours of becoming aware of the ` +
    `breach. Notification shall be sent to the Controller's designated contact. ` +
    `Security contact: ${CONTACT}`
  );

  section(pdf, "8. Data Subject Rights");
  body(pdf,
    `Given the stateless architecture, the Processor does not retain submitted payload data. ` +
    `The Processor cannot fulfill data subject access, erasure, or portability requests ` +
    `for submitted payload content because such data does not persist beyond the request ` +
    `lifecycle. The Controller is responsible for exercising data subject rights with ` +
    `respect to data they control.`
  );

  section(pdf, "9. Termination and Data Deletion");
  body(pdf,
    `Upon termination of the service agreement, the Processor confirms that no submitted ` +
    `payload data persists to delete, given the stateless architecture. Account metadata ` +
    `(email, API keys, usage counts) will be deleted within 30 days of written request.`
  );

  section(pdf, "10. Governing Law");
  body(pdf,
    `This DPA shall be governed by [JURISDICTION — e.g., the laws of the Republic of ` +
    `Peru / the European Union / the State of Delaware]. Any disputes shall be resolved ` +
    `in the courts of [JURISDICTION].`
  );

  section(pdf, "Signatures");
  body(pdf,
    `For the Data Controller:\n\n` +
    `Name: _______________________________\n` +
    `Title: _______________________________\n` +
    `Date:  _______________________________\n` +
    `Signature: ___________________________\n\n` +
    `For ${COMPANY} (Data Processor):\n\n` +
    `Name: Eliaz Bobadilla\n` +
    `Title: Founder\n` +
    `Date:  _______________________________\n` +
    `Signature: ___________________________`
  );

  finalize(pdf, "dpa-template.pdf");
}

// ─── 3. API Authentication & Security Policy ─────────────────────────────────

function generateApiAuthPolicy() {
  const pdf = doc("api-auth-policy.pdf");

  header(
    pdf,
    `${PRODUCT} — API Authentication & Security Policy`,
    `Effective ${EFFECTIVE_DATE} · ${COMPANY}`
  );

  section(pdf, "1. Purpose and Scope");
  body(pdf,
    `This document defines the authentication mechanisms, access controls, and security ` +
    `procedures governing the ${PRODUCT} API platform. It applies to all API consumers ` +
    `and governs ${PRODUCT}'s obligations with respect to API access security.`
  );

  section(pdf, "2. API Key Issuance");
  bullet(pdf, [
    "Each registered account is issued a unique, cryptographically random API key",
    "Keys are generated using a high-entropy random source — not derived from account data",
    "Keys are transmitted to the user once at creation time over TLS 1.3 encrypted channels",
    "Keys are stored in our system as hashed values — the plaintext key is never stored",
    "Key format: high-entropy alphanumeric string, minimum 32 characters",
  ]);

  section(pdf, "3. API Key Usage");
  bullet(pdf, [
    "All API requests must include the API key in the Authorization header: Bearer <key>",
    "TLS 1.3 is required on every request — HTTP connections are rejected",
    "Keys must not be embedded in client-side code, public repositories, or URLs",
    "Keys should be stored in environment variables or secrets management systems",
  ]);

  section(pdf, "4. Key Revocation");
  bullet(pdf, [
    "API keys can be revoked instantly via the user dashboard",
    "Revocation takes effect within seconds at the authentication layer",
    "Users should revoke keys immediately upon suspicion of compromise",
    "Revoked keys cannot be recovered — a new key must be generated",
    "Multiple keys per account are not currently supported; users must regenerate to rotate",
  ]);

  section(pdf, "5. Transport Security");
  bullet(pdf, [
    "TLS 1.3 enforced on all API endpoints — no TLS 1.2 or lower accepted",
    "Cipher suites: TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256",
    "TLS termination handled by Cloudflare before reaching origin servers",
    "HSTS (HTTP Strict Transport Security) enforced",
    "Certificate validity monitored and auto-renewed via Cloudflare",
  ]);

  section(pdf, "6. Rate Limiting & Abuse Prevention");
  bullet(pdf, [
    "Rate limits enforced at the Cloudflare edge layer before requests reach origin",
    "Limits are applied per API key — not per IP, preventing shared-IP false positives",
    "Exceeding rate limits returns HTTP 429 with Retry-After header",
    "Cloudflare WAF rules block malformed requests, known attack patterns, and injection attempts",
    "DDoS mitigation active at all times — volumetric attacks absorbed before reaching API",
  ]);

  section(pdf, "7. Origin Security");
  bullet(pdf, [
    "Origin servers are not exposed directly to the internet",
    "An X-Backend-Secret header mechanism prevents Cloudflare bypass attacks",
    "Server firewall rules whitelist only Cloudflare IP ranges for inbound traffic",
    "SSH access to servers is key-authenticated only — no password authentication",
  ]);

  section(pdf, "8. Data Handling During Authentication");
  body(pdf,
    `Authentication is validated at the middleware layer before any request payload is ` +
    `processed. If authentication fails, the request is rejected immediately and no ` +
    `payload processing occurs. Successful authentication does not result in any logging ` +
    `of the request payload — only the API key hash, endpoint, response code, and ` +
    `timestamp are recorded for usage metering.`
  );

  section(pdf, "9. Incident Response Procedures");
  bullet(pdf, [
    "Security incidents are triaged within 4 business hours of detection",
    "Affected API keys are revoked immediately upon confirmed compromise",
    "Affected customers notified within 72 hours of confirmed incident",
    "Post-incident review conducted within 7 days",
    `Security contact: ${CONTACT}`,
  ]);

  section(pdf, "10. Policy Review");
  body(pdf,
    `This policy is reviewed at minimum annually and updated as required when the ` +
    `authentication infrastructure changes. The current effective date is ${EFFECTIVE_DATE}. ` +
    `For the latest version, visit requiems.xyz/en/security`
  );

  finalize(pdf, "api-auth-policy.pdf");
}

// ─── run all ──────────────────────────────────────────────────────────────────

console.log(`Generating compliance documents → ${OUT_DIR}\n`);
generateSecurityPacket();
generateDPA();
generateApiAuthPolicy();
console.log("\nDone. Add to git if you want to ship them, or run before deploy.");
