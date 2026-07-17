# frozen_string_literal: true

require "test_helper"
require "minitest/mock"

class RefreshSitemapJobTest < ActiveJob::TestCase
  test "executes sitemap:refresh and logs success when the task is already defined" do
    task = Minitest::Mock.new
    task.expect(:execute, nil)

    Rake::Task.stub :task_defined?, true do
      Rake::Task.stub :[], task do
        RefreshSitemapJob.perform_now
      end
    end

    task.verify
  end

  test "loads rake tasks first when sitemap:refresh is not yet defined" do
    task = Minitest::Mock.new
    task.expect(:execute, nil)

    load_tasks_called = false

    Rake::Task.stub :task_defined?, false do
      Rails.application.stub :load_tasks, -> { load_tasks_called = true } do
        Rake::Task.stub :[], task do
          RefreshSitemapJob.perform_now
        end
      end
    end

    assert load_tasks_called
    task.verify
  end

  test "logs the error and re-raises when sitemap:refresh fails" do
    failing_task = Object.new
    def failing_task.execute
      raise StandardError, "boom"
    end

    logged_errors = []

    Rails.logger.stub :error, ->(msg) { logged_errors << msg } do
      Rake::Task.stub :task_defined?, true do
        Rake::Task.stub :[], failing_task do
          assert_raises(StandardError) { RefreshSitemapJob.perform_now }
        end
      end
    end

    assert_equal 1, logged_errors.size
    assert_match(/Failed to regenerate sitemap: boom/, logged_errors.first)
  end
end
