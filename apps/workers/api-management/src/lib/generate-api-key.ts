import { customAlphabet } from "nanoid";

const ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
const KEY_LENGTH = 24;
const nanoid = customAlphabet(ALPHABET, KEY_LENGTH);

export function generateApiKey(): string {
  const prefix = "requiem";
  const randomPart = nanoid();
  return `${prefix}_${randomPart}`;
}
