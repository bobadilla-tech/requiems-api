#!/usr/bin/env node
/**
 * OG Image Generator — 1200×630 PNG for requiems.xyz
 *
 * Usage:
 *   npm install @napi-rs/canvas   # once
 *   node scripts/gen-og-image.mjs
 *
 * Output: apps/dashboard/public/og-image.png
 */

import { createCanvas, loadImage } from "@napi-rs/canvas";
import { writeFileSync } from "fs";
import { dirname, resolve } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = resolve(__dirname, "..");

const W = 1200;
const H = 630;
const LOGO = resolve(ROOT, "apps/dashboard/app/assets/images/logo.png");
const OUT = resolve(ROOT, "apps/dashboard/public/og-image.png");

const canvas = createCanvas(W, H);
const ctx = canvas.getContext("2d");

// Background gradient
const grad = ctx.createLinearGradient(0, 0, W, H);
grad.addColorStop(0, "#0f172a");
grad.addColorStop(1, "#1e1b4b");
ctx.fillStyle = grad;
ctx.fillRect(0, 0, W, H);

// Grid lines
ctx.strokeStyle = "rgba(99,102,241,0.08)";
ctx.lineWidth = 1;
for (let x = 0; x <= W; x += 60) {
  ctx.beginPath();
  ctx.moveTo(x, 0);
  ctx.lineTo(x, H);
  ctx.stroke();
}
for (let y = 0; y <= H; y += 60) {
  ctx.beginPath();
  ctx.moveTo(0, y);
  ctx.lineTo(W, y);
  ctx.stroke();
}

// Accent glow
const glow = ctx.createRadialGradient(
  W * 0.85,
  H * 0.2,
  0,
  W * 0.85,
  H * 0.2,
  280,
);
glow.addColorStop(0, "rgba(99,102,241,0.25)");
glow.addColorStop(1, "rgba(99,102,241,0)");
ctx.fillStyle = glow;
ctx.fillRect(0, 0, W, H);

// Logo inside white circle so eyes are visible on dark bg
const LOGO_SIZE = 72;
const CIRCLE_R = 50;
const CIRCLE_X = 80 + CIRCLE_R;
const CIRCLE_Y = 72 + CIRCLE_R;

ctx.save();
ctx.fillStyle = "#ffffff";
ctx.beginPath();
ctx.arc(CIRCLE_X, CIRCLE_Y, CIRCLE_R, 0, Math.PI * 2);
ctx.fill();
ctx.restore();

try {
  const logo = await loadImage(LOGO);
  const offset = (CIRCLE_R * 2 - LOGO_SIZE) / 2;
  ctx.drawImage(logo, 80 + offset, 72 + offset, LOGO_SIZE, LOGO_SIZE);
} catch {
  ctx.fillStyle = "#6366f1";
  ctx.beginPath();
  ctx.arc(CIRCLE_X, CIRCLE_Y, 36, 0, Math.PI * 2);
  ctx.fill();
}

// Brand name — vertically centred next to logo circle
ctx.fillStyle = "#f8fafc";
ctx.font = "bold 28px sans-serif";
ctx.fillText("Requiems API", 196, 127);

// Headline
ctx.font = "bold 68px sans-serif";
ctx.fillText("All-in-one backend", 80, 295);
ctx.fillText("for SaaS products.", 80, 375);

// Subline
ctx.fillStyle = "#94a3b8";
ctx.font = "26px sans-serif";
ctx.fillText("Authentication · Validation · Payments · Global Data", 80, 438);

// Domain badge
function roundRect(c, x, y, w, h, r) {
  c.beginPath();
  c.moveTo(x + r, y);
  c.lineTo(x + w - r, y);
  c.arcTo(x + w, y, x + w, y + r, r);
  c.lineTo(x + w, y + h - r);
  c.arcTo(x + w, y + h, x + w - r, y + h, r);
  c.lineTo(x + r, y + h);
  c.arcTo(x, y + h, x, y + h - r, r);
  c.lineTo(x, y + r);
  c.arcTo(x, y, x + r, y, r);
  c.closePath();
}

ctx.fillStyle = "rgba(99,102,241,0.2)";
roundRect(ctx, 80, 492, 290, 50, 10);
ctx.fill();

ctx.fillStyle = "#a5b4fc";
ctx.font = "20px sans-serif";
ctx.fillText("requiems.xyz", 145, 523);

ctx.fillStyle = "#6366f1";
ctx.beginPath();
ctx.arc(118, 518, 7, 0, Math.PI * 2);
ctx.fill();

writeFileSync(OUT, canvas.toBuffer("image/png"));
console.log(`OG image written: ${OUT} (${W}x${H})`);
