require 'rubygems'
require 'bundler/setup'
require 'rake'
require 'prmd/rake_tasks/combine'
require 'prmd/rake_tasks/verify'
require 'prmd/rake_tasks/doc'
require 'prmd/link'
require 'redcarpet'

namespace :schema do
  Prmd::RakeTasks::Combine.new do |t|
    t.options[:meta] = 'schema/meta.json'
    t.paths << 'schema/schemata'
    t.output_file = 'schema/api.json'
  end

  Prmd::RakeTasks::Verify.new do |t|
    t.files << 'schema/api.json'
  end

  Prmd::RakeTasks::Doc.new do |t|
    t.files = { 'schema/api.json' => 'schema/api.md' }
  end

  task :html do
    markdown = Redcarpet::Markdown.new(Redcarpet::Render::HTML, autolink: true, tables: true, fenced_code_blocks: true)
    html = '<!DOCTYPE html><html><head><meta charset="utf-8"><title>API Spec</title></head><body>'
    html += markdown.render(File.read('schema/api.md'))
    html += '</body></html>'
    File.write("schema/api.html", html)
  end
end
 
task default: ['schema:combine', 'schema:verify', 'schema:doc', 'schema:html']
