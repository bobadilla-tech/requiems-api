#!/usr/bin/env python3
"""Fill divisions.es.yml and divisions.fr.yml from en/divisions.en.yml with cached translation."""

from __future__ import annotations

import json
import re
import sys
import time
from pathlib import Path

import yaml
from deep_translator import GoogleTranslator

ROOT = Path(__file__).resolve().parents[1]
EN_FILE = ROOT / "config/locales/en/divisions.en.yml"
CACHE_DIR = ROOT / "tmp/locale_translation_cache"

PLACEHOLDER_RE = re.compile(r"%\{[^}]+\}")


def should_skip(text: str) -> bool:
    stripped = text.strip()
    return (
        not stripped
        or stripped.startswith("{")
        or stripped.startswith("[")
        or stripped.startswith("POST ")
        or stripped.startswith("GET ")
        or stripped.startswith("/v1/")
    )


def mask(text: str) -> tuple[str, list[str]]:
    tokens: list[str] = []

    def repl(match: re.Match[str]) -> str:
        tokens.append(match.group(0))
        return f"__PH{len(tokens) - 1}__"

    return PLACEHOLDER_RE.sub(repl, text), tokens


def unmask(text: str, tokens: list[str]) -> str:
    for index, token in enumerate(tokens):
        text = text.replace(f"__PH{index}__", token)
    return text


def collect_strings(node, out: set[str]) -> None:
    if isinstance(node, dict):
        for value in node.values():
            collect_strings(value, out)
    elif isinstance(node, list):
        for item in node:
            collect_strings(item, out)
    elif isinstance(node, str) and not should_skip(node):
        out.add(node)


def translate_strings(strings: set[str], locale: str, cache_path: Path) -> dict[str, str]:
    cache: dict[str, str] = {}
    if cache_path.exists():
        cache = json.loads(cache_path.read_text(encoding="utf-8"))

    translator = GoogleTranslator(source="en", target=locale)
    pending = sorted(s for s in strings if s not in cache)
    print(f"Translating {len(pending)} unique strings to {locale} ({len(cache)} cached)")

    for index, text in enumerate(pending, start=1):
        masked, tokens = mask(text)
        try:
            translated = translator.translate(masked)
        except Exception as error:
            print(f"  skip ({error}): {text[:60]}...")
            translated = text
        cache[text] = unmask(translated, tokens)
        if index % 25 == 0:
            print(f"  {index}/{len(pending)}")
            cache_path.write_text(json.dumps(cache, ensure_ascii=False, indent=2), encoding="utf-8")
        time.sleep(0.08)

    cache_path.write_text(json.dumps(cache, ensure_ascii=False, indent=2), encoding="utf-8")
    return cache


def polish_brand(text: str) -> str:
    return (
        text.replace("Réquiems", "Requiems")
        .replace("réquiems", "requiems")
        .replace("API de Réquiems", "Requiems API")
        .replace("API de Requiems", "Requiems API")
    )


def apply_translations(node, mapping: dict[str, str]):
    if isinstance(node, dict):
        return {key: apply_translations(value, mapping) for key, value in node.items()}
    if isinstance(node, list):
        return [apply_translations(item, mapping) for item in node]
    if isinstance(node, str):
        if should_skip(node):
            return node
        return polish_brand(mapping.get(node, node))
    return node


def write_locale(locale: str, tree, mapping: dict[str, str]) -> None:
    translated = apply_translations(tree, mapping)
    output = ROOT / "config/locales" / locale / f"divisions.{locale}.yml"
    body = yaml.dump(
        {locale: {"divisions": translated}},
        allow_unicode=True,
        sort_keys=False,
        width=10000,
    )
    output.write_text(f"---\n{body}", encoding="utf-8")
    print(f"Written: {output}")


def main() -> None:
    locales = sys.argv[1:] if len(sys.argv) > 1 else ["es", "fr"]
    CACHE_DIR.mkdir(parents=True, exist_ok=True)
    data = yaml.safe_load(EN_FILE.read_text(encoding="utf-8"))
    tree = data["en"]["divisions"]
    strings: set[str] = set()
    collect_strings(tree, strings)

    for locale in locales:
        cache_path = CACHE_DIR / f"divisions_{locale}.json"
        mapping = translate_strings(strings, locale, cache_path)
        write_locale(locale, tree, mapping)


if __name__ == "__main__":
    main()
