# frozen_string_literal: true

require "test_helper"

class DashboardOverviewHelperTest < ActionView::TestCase
  include Dashboard::OverviewHelper

  Log = Struct.new(:status_code, :credits_used)

  def kind(code) = activity_status_kind(Log.new(code, 1))

  test "status kinds at every boundary" do
    {
      nil => :success, 0 => :success, 1 => :unknown, 199 => :unknown, 200 => :success, 299 => :success,
      300 => :redirect, 399 => :redirect, 400 => :client_error, 499 => :client_error,
      500 => :server_error, 599 => :server_error, 600 => :server_error, 999 => :server_error
    }.each do |code, expected|
      assert_equal expected, kind(code), "status #{code.inspect}"
    end
  end

  test "only client and server errors are errors" do
    assert activity_error?(Log.new(404, 0))
    assert activity_error?(Log.new(503, 0))
    assert activity_error?(Log.new(999, 0))
    assert_not activity_error?(Log.new(200, 1))
    assert_not activity_error?(Log.new(302, 1))
    assert_not activity_error?(Log.new(nil, 1))
    assert_not activity_error?(Log.new(101, 1))
  end

  test "every status kind has a chip class" do
    %i[success redirect client_error server_error unknown].each do |k|
      assert Dashboard::OverviewHelper::STATUS_CHIPS.key?(k)
    end
  end

  test "credits labels: success rows, charged error rows, free and unrecorded" do
    assert_equal "1 credit", activity_credits_label(Log.new(200, 1))
    assert_equal "2 credits", activity_credits_label(Log.new(200, 2))
    assert_equal "1 credit (charged)", activity_credits_label(Log.new(404, 1))
    assert_equal "3 credits (charged)", activity_credits_label(Log.new(500, 3))
    assert_equal "not charged", activity_credits_label(Log.new(404, 0))
    assert_equal "not charged", activity_credits_label(Log.new(200, 0))
    assert_equal "not recorded", activity_credits_label(Log.new(404, nil))
    assert_equal "not recorded", activity_credits_label(Log.new(200, nil))
  end

  test "every status kind has a translation in every locale" do
    %i[en es fr].each do |locale|
      %w[success redirect client_error server_error unknown legacy_ok].each do |k|
        assert I18n.exists?("dashboard.overview.activity.status.#{k}", locale), "#{locale}: #{k}"
      end
    end
  end

  test "active filter chips are visually distinct" do
    assert_includes activity_filter_classes(true), "bg-blue-600"
    assert_not_includes activity_filter_classes(false), "bg-blue-600"
  end
end
