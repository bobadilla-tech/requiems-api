# AppConfig Service

The `AppConfig` service centralizes all environment variable access in the Rails
dashboard with validation and fail-fast behavior.

## Features

- ✅ **Fail-fast validation**: App won't start if required config is missing
- ✅ **Type safety**: URLs, numeric IDs, and other types are validated
- ✅ **Self-documenting**: All config vars listed in one place
- ✅ **Easy testing**: Single place to stub config in tests
- ✅ **No typos**: Compile-time detection of config issues

## Location

- Service class: `apps/dashboard/app/lib/app_config.rb`
- Initializer: `apps/dashboard/config/initializers/app_config.rb`
- Example config: `apps/dashboard/.env.example`

## Usage

### Basic Usage

```ruby
# Access config values
api_key = AppConfig.api_management_api_key
store_slug = AppConfig.lemonsqueezy_store_slug
api_url = AppConfig.api_base_url
```

### Helper Methods

```ruby
# Get variant ID for a plan and billing cycle
variant_id = AppConfig.instance.variant_id_for(
  plan: "developer",
  billing_cycle: "monthly"
)

# Check if SMTP is configured
if AppConfig.instance.smtp_configured?
  # Send email
end
```

### In Tests

For testing, you can stub AppConfig values:

```ruby
# Stub a single value
allow(AppConfig).to receive(:api_base_url).and_return("http://test.example.com")

# Or stub the instance
config = instance_double(AppConfig, api_base_url: "http://test.example.com")
allow(AppConfig).to receive(:instance).and_return(config)
```

## Required Environment Variables

### Always Required

These are required in all environments:

- `API_MANAGEMENT_API_KEY` - API key for the API Management Cloudflare Worker
- `LEMONSQUEEZY_STORE_ID` - Your LemonSqueezy store ID (numeric)
- `LEMONSQUEEZY_SIGNING_SECRET` - Webhook signature secret
- `LEMONSQUEEZY_DEVELOPER_MONTHLY_VARIANT_ID` - Developer plan monthly variant
- `LEMONSQUEEZY_DEVELOPER_YEARLY_VARIANT_ID` - Developer plan yearly variant
- `LEMONSQUEEZY_BUSINESS_MONTHLY_VARIANT_ID` - Business plan monthly variant
- `LEMONSQUEEZY_BUSINESS_YEARLY_VARIANT_ID` - Business plan yearly variant
- `LEMONSQUEEZY_PROFESSIONAL_MONTHLY_VARIANT_ID` - Professional plan monthly
  variant
- `LEMONSQUEEZY_PROFESSIONAL_YEARLY_VARIANT_ID` - Professional plan yearly
  variant

### Optional (with defaults)

- `API_MANAGEMENT_URL` - Default: `https://api-management.requiems.xyz`
- `API_BASE_URL` - Default: `https://api.requiems.xyz`
- `PLAYGROUND_API_KEY` - Default: `requiem_notprovisioned0000000000`
  (format-valid but non-functional; must be set to a real provisioned key
  for the public playground/demo forms to work)
- `LEMONSQUEEZY_STORE_SLUG` - Default: `requiems`

### Production Only

These are only needed in production for email functionality:

- `SMTP_ADDRESS` - SMTP server address
- `SMTP_PORT` - Default: `587`
- `SMTP_DOMAIN` - Your domain
- `SMTP_USERNAME` - SMTP authentication username
- `SMTP_PASSWORD` - SMTP authentication password
- `MAILER_HOST` - Default: `requiems.xyz`

## CI/CD Setup

The initializer automatically skips validation for database and asset tasks:

- `db:create`, `db:migrate`, `db:rollback`, etc.
- `assets:precompile`, `assets:clean`, etc.
- `tmp:clear`, `log:clear`, etc.

### Test Environment

For test environment, AppConfig automatically provides safe defaults for all
required environment variables. **You do NOT need to set any environment
variables to run tests.**

This means CI pipelines work out of the box without managing test secrets:

```yaml
# No environment variables needed for tests!
- name: Run tests
  run: bin/rails test
```

### Production/Staging Environment

For production or staging deployments, you **must** provide all required
environment variables:

```yaml
env:
  API_MANAGEMENT_API_KEY: ${{ secrets.API_MANAGEMENT_API_KEY }}
  LEMONSQUEEZY_STORE_ID: ${{ secrets.LEMONSQUEEZY_STORE_ID }}
  LEMONSQUEEZY_SIGNING_SECRET: ${{ secrets.LEMONSQUEEZY_SIGNING_SECRET }}
  # ... all other required vars
```

## Error Handling

### Missing Required Variable

```
AppConfig::MissingConfigError: Missing required environment variable: API_MANAGEMENT_API_KEY
```

**Solution**: Add the missing environment variable to your `.env` file or
environment.

### Invalid Configuration

```
AppConfig::InvalidConfigError: API_BASE_URL must be a valid HTTP/HTTPS URL
```

**Solution**: Check the format of the environment variable value.

### Behavior by Environment

- **Production**: App fails to start if config is missing or invalid
- **Test**: Uses safe test defaults if environment variables are not set (tests
  run without secrets)
- **Development**: App starts with a warning (allows local dev without all
  secrets)

#### Test Defaults

In test environment, AppConfig provides these defaults automatically:

- `API_MANAGEMENT_API_KEY`: `"test_api_management_key"`
- `LEMONSQUEEZY_STORE_ID`: `"12345"`
- `LEMONSQUEEZY_SIGNING_SECRET`: `"test_signing_secret"`
- All variant IDs: `"123456"`, `"123457"`, etc.

This means tests can run without any environment variables set.

## Adding New Configuration

To add a new environment variable:

1. Add an `attr_reader` to `AppConfig` class
2. Load it in the `load_config` method using `require_env` or `optional_env`
3. Add validation in `validate_config` if needed
4. Update `.env.example` with the new variable
5. Document it in this file

Example:

```ruby
# In app/lib/app_config.rb
class AppConfig
  attr_reader :my_new_config

  def load_config
    @my_new_config = require_env("MY_NEW_CONFIG")
    # or
    @my_new_config = optional_env("MY_NEW_CONFIG", default: "default_value")
  end
end
```

## Migration from Direct ENV Access

Before (bad):

```ruby
api_key = ENV["API_MANAGEMENT_API_KEY"]
store_id = ENV["LEMONSQUEEZY_STORE_ID"]
```

After (good):

```ruby
api_key = AppConfig.api_management_api_key
store_id = AppConfig.lemonsqueezy_store_id
```

## Related Files

- Service implementation: `app/lib/app_config.rb`
- Initializer: `config/initializers/app_config.rb`
- Example environment: `.env.example`
- This documentation: `docs/APP_CONFIG.md`
