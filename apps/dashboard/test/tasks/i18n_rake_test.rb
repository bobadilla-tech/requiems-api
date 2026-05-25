# frozen_string_literal: true

require "test_helper"
require "fileutils"

module SystemCapture
  def system(*args)
    if (calls = Thread.current[:captured_system_calls])
      calls << args.first
      true
    else
      super
    end
  end
end

class I18nRakeTest < ActiveSupport::TestCase
  parallelize(workers: 1)

  Object.prepend(SystemCapture) unless Object.ancestors.include?(SystemCapture)

  setup do
    Rails.application.load_tasks
    @system_calls = []
    Thread.current[:captured_system_calls] = @system_calls
  end

  teardown do
    Thread.current[:captured_system_calls] = nil
    Rake::Task["i18n:report"].reenable
    Rake::Task["i18n:todos"].reenable
  end

  test "i18n:report runs for es and fr by default" do
    capture_io { Rake::Task["i18n:report"].invoke }
    assert_equal(
      [ "bundle exec i18n-tasks missing -l es", "bundle exec i18n-tasks missing -l fr" ],
      @system_calls
    )
  end

  test "i18n:report runs for single locale when specified" do
    capture_io { Rake::Task["i18n:report"].invoke("es") }
    assert_equal [ "bundle exec i18n-tasks missing -l es" ], @system_calls
  end

  test "i18n:report outputs locale header" do
    out, _ = capture_io { Rake::Task["i18n:report"].invoke("fr") }
    assert_includes out, "Missing translations for FR"
    assert_includes out, "=" * 60
  end

  test "i18n:todos reports completion when no TODOs found" do
    with_locale_dir do |dir|
      File.write("#{dir}/en.yml", "en:\n  greeting: hello\n")
      out, _ = capture_io { Rake::Task["i18n:todos"].invoke }
      assert_includes out, "all locales are complete"
    end
  end

  test "i18n:todos lists TODO placeholders with filename" do
    with_locale_dir do |dir|
      File.write("#{dir}/es.yml", "es:\n  greeting: 'TODO: translate hello'\n")
      out, _ = capture_io { Rake::Task["i18n:todos"].invoke }
      assert_includes out, "es.yml"
      assert_includes out, "TODO: translate"
    end
  end

  test "i18n:todos counts multiple placeholders correctly" do
    with_locale_dir do |dir|
      File.write("#{dir}/es.yml", <<~YAML)
        es:
          a: "TODO: translate a"
          b: "TODO: translate b"
          c: "TODO: translate c"
      YAML
      out, _ = capture_io { Rake::Task["i18n:todos"].invoke }
      assert_includes out, "Found 3 TODO"
    end
  end

  test "i18n:todos includes correct line numbers" do
    with_locale_dir do |dir|
      File.write("#{dir}/es.yml", <<~YAML)
        es:
          greeting: hello
          farewell: "TODO: translate goodbye"
      YAML
      out, _ = capture_io { Rake::Task["i18n:todos"].invoke }
      assert_includes out, "es.yml:3:"
    end
  end

  test "i18n:todos scans nested locale subdirectories" do
    with_locale_dir do |dir|
      FileUtils.mkdir_p("#{dir}/es")
      File.write("#{dir}/es/common.yml", "es:\n  key: 'TODO: translate common'\n")
      out, _ = capture_io { Rake::Task["i18n:todos"].invoke }
      assert_includes out, "es/common.yml"
    end
  end

  private

  def with_locale_dir
    Dir.mktmpdir do |tmpdir|
      locale_dir = File.join(tmpdir, "config", "locales")
      FileUtils.mkdir_p(locale_dir)
      stub_rails_root(tmpdir) { yield locale_dir }
    end
  end

  def stub_rails_root(path)
    root_path = Pathname.new(path)
    Rails.singleton_class.alias_method(:__orig_root__, :root)
    Rails.define_singleton_method(:root) { root_path }
    yield
  ensure
    Rails.singleton_class.alias_method(:root, :__orig_root__)
    Rails.singleton_class.remove_method(:__orig_root__)
  end
end
