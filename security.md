# Security Policy

## Supported Versions

Requiems API only supports the latest released version. We do not provide
security patches, backports, or maintenance updates for older versions.

| Version        | Supported          |
| -------------- | ------------------ |
| Latest release | :white_check_mark: |
| Older releases | :x:                |

If you are running an older release, please upgrade to the latest version before
reporting or requesting support for a vulnerability.

## Deployment and Support Scope

Requiems API is primarily intended to be operated by us as a hosted service on
our own infrastructure. The core source code is open source, and you may
self-deploy it, but self-hosted deployments are not an officially supported
product offering.

Some backend functionality depends on proprietary files, private datasets, and
operational infrastructure that are not included in this repository and are not
distributed publicly. Because of that, self-hosted installations may require
their own replacement data, configuration, or service integrations, and may not
match the behavior of the hosted Requiems API service.

Security reports are welcome for the open source code in this repository and for
the hosted Requiems API service. We generally do not provide security support
for custom self-hosted deployments, local infrastructure, private forks, or
modified versions outside our control.

## Reporting a Vulnerability

We take the security of Requiems API seriously. If you discover a security
vulnerability, please report it responsibly.

### Where to Report

**DO NOT** create a public GitHub issue for security vulnerabilities.

Instead, please report security vulnerabilities to:

**Email:** eliaz@bobadilla.tech

### What to Include

Please include the following information in your report:

- **Description** of the vulnerability
- **Steps to reproduce** the issue
- **Potential impact** of the vulnerability
- **Suggested fix** (if you have one)
- **Your contact information** for follow-up

### What to Expect

- **Acknowledgment:** We'll acknowledge receipt of your report within 48 hours
- **Updates:** We'll keep you informed about our progress as we investigate
- **Timeline:** We aim to validate and patch critical vulnerabilities within 7
  days
- **Credit:** We'll credit you in our security advisories (unless you prefer to
  remain anonymous)

## Security Best Practices

### For Developers and Contributors

1. **Secure Your API Keys**
   - Store API keys in environment variables, never in code
   - Use different keys for different environments (test vs live)
   - Rotate keys regularly
   - Revoke compromised keys immediately

2. **Code Security**
   - Never commit `.env` files or secrets to version control
   - Review code for common vulnerabilities (SQL injection, XSS, etc.)
   - Keep dependencies updated
   - Run security scans before submitting PRs

3. **Access Control**
   - Use strong passwords and enable 2FA on your GitHub account
   - Follow the principle of least privilege
   - Review and approve third-party access carefully

## Scope

This security policy covers:

- ✅ The latest released version of the Requiems API source code in this
  repository
- ✅ The hosted Requiems API service operated by us
- ✅ Official deployment configurations used by the hosted service
- ✅ Dependencies we directly control
- ❌ Older releases
- ❌ Private forks, modified versions, or custom self-hosted deployments
- ❌ Third-party services or integrations
- ❌ Infrastructure outside of our control

## Questions?

For general security questions or concerns, please contact: eliaz@bobadilla.tech

Thank you for helping keep Requiems API and our users secure! 🛡️
