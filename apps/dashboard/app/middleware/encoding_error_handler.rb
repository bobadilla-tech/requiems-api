# frozen_string_literal: true

class EncodingErrorHandler
  def initialize(app)
    @app = app
  end

  def call(env)
    @app.call(env)
  rescue Encoding::CompatibilityError
    [ 400, { "Content-Type" => "text/plain", "Content-Length" => "11" }, [ "Bad Request" ] ]
  end
end
