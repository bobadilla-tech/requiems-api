# frozen_string_literal: true

namespace :i18n do
  desc "Report missing translations for one or all non-English locales. Usage: rake i18n:report[es] or rake i18n:report"
  task :report, [:locale] do |_, args|
    locales = args[:locale] ? [args[:locale]] : %w[es fr]
    locales.each do |locale|
      puts "\n#{"=" * 60}"
      puts "Missing translations for #{locale.upcase}"
      puts "=" * 60
      system("bundle exec i18n-tasks missing -l #{locale}")
    end
  end

  desc "List all TODO: translate placeholders across all locale files"
  task :todos do
    locale_dir = Rails.root.join("config/locales")
    todos = []

    Dir.glob("#{locale_dir}/**/*.yml").sort.each do |file|
      content = File.read(file)
      content.each_line.with_index(1) do |line, lineno|
        todos << "#{file.sub(locale_dir.to_s + '/', '')}:#{lineno}: #{line.strip}" if line.include?("TODO: translate")
      end
    end

    if todos.empty?
      puts "No TODO: translate placeholders found — all locales are complete!"
    else
      puts "Found #{todos.size} TODO: translate placeholders:\n\n"
      todos.each { |t| puts t }
    end
  end
end
