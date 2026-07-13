# frozen_string_literal: true

require "test_helper"

class CreditAdjustmentTest < ActiveSupport::TestCase
  def setup
    @user = create_user(email: "test@example.com")
  end

  test "valid with required fields" do
    adjustment = CreditAdjustment.new(user: @user, amount: 10, adjustment_type: "bonus")
    assert adjustment.valid?
  end

  test "requires user_id" do
    adjustment = CreditAdjustment.new(amount: 10, adjustment_type: "bonus")
    I18n.with_locale(:en) do
      assert_not adjustment.valid?
      assert_includes adjustment.errors[:user], "must exist"
    end
  end

  test "requires amount" do
    adjustment = CreditAdjustment.new(user: @user, adjustment_type: "bonus")
    I18n.with_locale(:en) do
      assert_not adjustment.valid?
      assert_includes adjustment.errors[:amount], "can't be blank"
    end
  end

  test "requires adjustment_type" do
    adjustment = CreditAdjustment.new(user: @user, amount: 10)
    I18n.with_locale(:en) do
      assert_not adjustment.valid?
      assert_includes adjustment.errors[:adjustment_type], "can't be blank"
    end
  end

  test "amount must be numeric" do
    adjustment = CreditAdjustment.new(user: @user, amount: "abc", adjustment_type: "bonus")
    I18n.with_locale(:en) do
      assert_not adjustment.valid?
      assert_includes adjustment.errors[:amount], "is not a number"
    end
  end
end
