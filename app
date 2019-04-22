#!/usr/bin/env ruby
# frozen_string_literal: true

require 'thor'
require 'colorize'

module OrsumInflandi
  BACKEND_IMAGE_NAME = 'orsa-scholis/orsum-inflandi'

  class CLI < Thor
    package_name 'OrsumInflandi'

    desc 'start', 'Start the app or a part of it'
    option :backend, type: :boolean, default: true
    option :frontend, type: :boolean, default: true
    option :force_build, type: :boolean

    def start
      require_relative 'cli/commands/start'
      Commands::Start.new(options.dup).run
    end
  end
end

OrsumInflandi::CLI.start(ARGV)
