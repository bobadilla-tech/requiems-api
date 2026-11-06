$LOAD_PATH.unshift File.expand_path('../../clients/ruby/lib', __dir__)
require 'openapi_client'

api = OpenapiClient::AdviceApi.new
begin
  api.v1_text_advice_get
rescue StandardError => e
  if e.respond_to?(:code) && e.code == 401
    puts "✅ Ruby Client E2E OK (Expected 401)"
    exit 0
  end
  puts "❌ Failed: #{e}"
  exit 1
end
