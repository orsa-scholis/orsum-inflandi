#!/usr/bin/env ruby
# frozen_string_literal: true

require 'thor'
require 'colorize'

BACKEND_IMAGE_NAME = 'orsa-scholis/orsum-inflandi'

module OrsumInflandi
  class CLI < Thor
    desc 'start', 'Start the app or a part of it'
    option :backend, type: :boolean, default: true
    option :frontend, type: :boolean, default: true
    option :force_build, type: :boolean

    def start
      threads = []
      threads << Thread.new(Dir.pwd, &method(:start_backend)) if options[:backend]
      threads << Thread.new(&method(:start_frontend)) if options[:frontend]

      %w[INT TERM].each do |signal|
        Signal.trap(signal) { kill_threads(threads) }
      end

      threads.each(&:join)
    end

    private

    def kill_threads(threads)
      puts "\nReceived KILL SIGNAL, shutting down..."
      threads.each(&:kill)
      exit
    end

    def execute_command(command, &block)
      IO.popen(command, err: %i[child out]).each(&block)
    end

    def build_backend(directory)
      backend_log('Building image')
      execute_command(%W[docker build -t #{BACKEND_IMAGE_NAME} #{directory}/backend], &method(:backend_log))
    end

    def start_backend(directory)
      build_backend(directory) unless backend_image_exists?
      execute_command(
        %w[docker run --rm -p 4560:4560 orsa-scholis/orsum-inflandi --verbose],
        &method(:backend_log)
      )
    end

    def start_frontend
      Dir.chdir('frontend') { execute_command(%w[yarn run start], &method(:frontend_log)) }
    end

    def backend_image_exists?
      system "[ \"$(docker images -q #{BACKEND_IMAGE_NAME} 2>/dev/null)\" != \"\" ]"
    end

    def frontend_log(log)
      puts "#{add_padding('Frontend').blue} #{log}"
    end

    def backend_log(log)
      puts "#{add_padding('Backend').green} #{log}"
    end

    def add_padding(str)
      "#{str.ljust(10, ' ')}|"
    end
  end
end

OrsumInflandi::CLI.start(ARGV)
