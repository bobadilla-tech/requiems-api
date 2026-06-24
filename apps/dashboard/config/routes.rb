# frozen_string_literal: true

require_relative "../lib/division_slugs.rb"
require_relative "../lib/comparison_slugs.rb"
require_relative "../lib/industry_slugs.rb"
require "sidekiq/web"
require "sidekiq/cron/web"

Rails.application.routes.draw do
  # Non-locale routes: health check, sitemaps, webhooks, API proxy, dev tools
  get "up" => "rails/health#show", as: :rails_health_check
  get "manifest" => "pwa#manifest", as: :pwa_manifest, defaults: { format: :json }
  get "llms.txt",      to: "sitemap#llms",      defaults: { format: :text }
  get "llms-full.txt", to: "sitemap#llms_full", defaults: { format: :text }
  get "apis/:id/index.md", to: "sitemap#api_doc", defaults: { format: :text }
  post "api/proxy", to: "api_proxy#create"
  post "locale", to: "locale#update", as: :switch_locale

  # Tool demo form submissions — server-side API calls that render Turbo Frame responses.
  # Outside locale scope so they share the same CSRF protection as api/proxy.
  post "tools/demos/unit-conversion",   to: "tool_demos#unit_conversion",   as: :tool_demo_unit_conversion
  post "tools/demos/sentiment-analysis", to: "tool_demos#sentiment_analysis", as: :tool_demo_sentiment_analysis
  post "tools/demos/email-validator",    to: "tool_demos#email_validator",    as: :tool_demo_email_validator
  post "tools/demos/email-normalizer",  to: "tool_demos#email_normalizer",  as: :tool_demo_email_normalizer
  post "tools/demos/domain-checker",   to: "tool_demos#domain_checker",   as: :tool_demo_domain_checker
  post "tools/demos/phone-validator",  to: "tool_demos#phone_validator",  as: :tool_demo_phone_validator

  namespace :webhooks do
    post "lemonsqueezy", to: "lemonsqueezy#create"
  end

  authenticate :user, ->(u) { u.admin? } do
    mount Sidekiq::Web => "/sidekiq"
  end

  if Rails.env.development?
    mount LetterOpenerWeb::Engine, at: "/letter_opener"
  end

  locale_route_pattern = Regexp.union(*Rails.application.config.i18n.available_locales.map(&:to_s))
  # All user-facing routes scoped under optional locale prefix (/en/, /es/, /fr/, ...)
  scope "(:locale)", locale: locale_route_pattern do
    devise_for :users, controllers: {
      registrations: "users/registrations",
      sessions: "users/sessions",
      confirmations: "users/confirmations"
    }

    root "home#index"

    namespace :dashboard do
      root "overview#index"

      resources :api_keys do
        member do
          post :regenerate
          delete :revoke
        end
      end

      resource :usage, only: [ :show ], controller: "usage" do
        collection do
          get :by_endpoint
          get :by_date
          get :export
        end
      end

      resource :billing, only: [ :show, :update ], controller: "billing"
      post "billing/checkout", to: "billing#checkout", as: :checkout_billing
      post "billing/portal", to: "billing#portal", as: :portal_billing
      delete "billing/cancel_subscription", to: "billing#cancel_subscription", as: :cancel_subscription_billing

      resource :referral, only: [ :show ], controller: "referrals"

      resources :invoices, only: [ :index, :show ]

      get "quick_start", to: "quick_start#index"

      resource :settings, only: [ :show, :update ] do
        member do
          post   :request_deletion
          get    :confirm_deletion
          delete :execute_deletion
        end
      end
    end

    namespace :admin do
      root "dashboard#index"

      authenticate :user, ->(u) { u.admin? } do
        resources :users do
          member do
            post :suspend
            post :unsuspend
            post :ban
            post :make_admin
            post :remove_admin
          end

          resource :promotion, only: [ :create, :destroy ],
                               controller: "promotions"
          resources :credit_adjustments, only: [ :new, :create ]
        end

        resources :api_keys, only: [ :index, :show ] do
          member do
            delete :revoke
          end
        end

        resources :abuse_reports do
          member do
            post :resolve
            post :investigate
          end
        end

        resources :private_deployments, only: [ :index, :show ] do
          member do
            patch :activate
            patch :cancel
          end
        end

        namespace :analytics do
          get :usage
          get :revenue
          get :system_health
        end
      end
    end

    get "docs", to: "home#docs"
    get "pricing", to: "home#pricing"
    get "about", to: "home#about"
    get "team", to: "home#team"
    get "privacy", to: "home#privacy"
    get "terms", to: "home#terms"
    get "contact", to: "home#contact"
    get "api_reference", to: "home#api_reference"
    get "changelog", to: "home#changelog"

    get "blog", to: "home#blog"
    get "status", to: "home#status"
    get "glossary", to: "home#glossary"
    get "error_codes", to: "home#error_codes"
    get "faq", to: "home#faq"
    get "for-llms", to: "home#for_llms"
    get "domain-checker", to: "home#domain_checker"

    get "suggest-an-api", to: "suggestions#new", as: "suggest_api"
    post "suggest-an-api", to: "suggestions#create"
    get "talk-to-sales", to: "sales_inquiries#new", as: "talk_to_sales"
    post "talk-to-sales", to: "sales_inquiries#create"
    get "private-deployment", to: "private_deployments#new", as: "new_private_deployment"
    post "private-deployment", to: "private_deployments#create", as: "private_deployments"

    get "examples", to: "examples#index"
    get "apis/:id/index.md", to: "sitemap#api_doc", defaults: { format: :text }
    resources :apis, only: [ :index, :show ]
    resources :tools, only: [ :index, :show ]
    resources :categories, only: [ :show ]
    resources :examples, only: [ :show ]

    get "divisions", to: "divisions#index", as: :divisions
    get "systems", to: "systems#index", as: :systems
    get "systems/:system_slug", to: "systems#show", as: :system,
        constraints: { system_slug: /identity-risk|payments-intelligence|global-data|data-integrity|developer-utilities/ }

    get "case-studies", to: "case_studies#index", as: :case_studies
    get "case-studies/:slug", to: "case_studies#show", as: :case_study

    get "compare", to: "comparisons#index", as: :comparisons
    get "compare/:slug", to: "comparisons#show", as: :comparison,
        constraints: { slug: Regexp.union(*ComparisonSlugs::ALL) }

    get "industries", to: "industries#index", as: :industries
    get "industries/:industry_slug", to: "industries#show", as: :industry,
        constraints: { industry_slug: Regexp.union(*IndustrySlugs::ALL) }

    get ":division_slug", to: "divisions#show",
                         constraints: { division_slug: Regexp.union(*DivisionSlugs::ALL) },
                         as: :division
  end
end
