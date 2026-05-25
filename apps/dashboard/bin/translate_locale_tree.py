#!/usr/bin/env python3
"""Translate a nested locale YAML tree from English to target locale."""

from __future__ import annotations

import re
import sys
import time
from pathlib import Path

import yaml

try:
    from deep_translator import GoogleTranslator
except ImportError:
    print("Install: pip install deep-translator pyyaml", file=sys.stderr)
    sys.exit(1)

PLACEHOLDER_RE = re.compile(r"%\{[^}]+\}")
HTML_TAG_RE = re.compile(r"<[^>]+>")


def mask_placeholders(text: str) -> tuple[str, list[str]]:
    tokens: list[str] = []

    def repl(match: re.Match[str]) -> str:
        tokens.append(match.group(0))
        return f"__PH{len(tokens) - 1}__"

    masked = PLACEHOLDER_RE.sub(repl, text)
    return masked, tokens


def unmask_placeholders(text: str, tokens: list[str]) -> str:
    for index, token in enumerate(tokens):
        text = text.replace(f"__PH{index}__", token)
    return text


def should_skip_translation(text: str) -> bool:
    stripped = text.strip()
    if not stripped:
        return True
    if stripped.startswith("{") or stripped.startswith("["):
        return True
    if stripped.startswith("POST ") or stripped.startswith("GET "):
        return True
    if stripped.startswith("/v1/"):
        return True
    if text == "TODO: translate":
        return True
    return False


def translate_string(text: str, translator: GoogleTranslator) -> str:
    if should_skip_translation(text):
        return text

    masked, tokens = mask_placeholders(text)
    # Preserve simple HTML wrappers used in docs_notice_html
    parts = HTML_TAG_RE.split(masked)
    tags = HTML_TAG_RE.findall(masked)
    translated_parts = []
    for part in parts:
        if not part.strip():
            translated_parts.append(part)
            continue
        try:
            translated_parts.append(translator.translate(part))
        except Exception:
            translated_parts.append(part)
        time.sleep(0.05)

    rebuilt = ""
    tag_index = 0
    for index, part in enumerate(translated_parts):
        rebuilt += part
        if index < len(tags):
            rebuilt += tags[tag_index]
            tag_index += 1

    return unmask_placeholders(rebuilt, tokens)


def translate_node(node, translator: GoogleTranslator):
    if isinstance(node, dict):
        return {key: translate_node(value, translator) for key, value in node.items()}
    if isinstance(node, list):
        return [translate_node(item, translator) for item in node]
    if isinstance(node, str):
        return translate_string(node, translator)
    return node


def main() -> None:
    if len(sys.argv) != 4:
        print(f"Usage: {sys.argv[0]} <en.yml> <target_locale> <output.yml>", file=sys.stderr)
        sys.exit(1)

    source = Path(sys.argv[1])
    locale = sys.argv[2]
    output = Path(sys.argv[3])

    data = yaml.safe_load(source.read_text(encoding="utf-8"))
    if "en" not in data:
        print("Expected top-level 'en' key", file=sys.stderr)
        sys.exit(1)

    target_code = "es" if locale == "es" else "fr"
    translator = GoogleTranslator(source="en", target=target_code)
    translated_root = translate_node(data["en"], translator)

    output.parent.mkdir(parents=True, exist_ok=True)
    output.write_text(
        yaml.dump({locale: translated_root}, allow_unicode=True, sort_keys=False, width=10000),
        encoding="utf-8",
    )
    print(f"Written: {output}")


if __name__ == "__main__":
    main()
