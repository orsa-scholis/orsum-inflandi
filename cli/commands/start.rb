# frozen_string_literal: true

require 'etc'
require_relative '../logger'

module OrsumInflandi
  module Commands
    class Start
      BACKEND_IMAGE_NAME = 'orsa-scholis/orsum-inflandi-ii'
      BACKEND_DEV_IMAGE_NAME = 'orsa-scholis/orsum-inflandi-ii-dev'

      def initialize(options)
        @options = options
      end

      def run
        threads = []
        threads << backend_thread if @options[:backend]
        threads << Thread.new(&method(:start_frontend)) if @options[:frontend]

        %w[INT TERM].each do |signal|
          Signal.trap(signal) { kill_threads(threads) }
        end

        threads.each(&:join)
      end

      private

      def backend_thread
        Thread.new(Dir.pwd) { |directory| @options[:dev] ? start_dev_backend(directory) : start_backend(directory) }
      end

      def kill_threads(threads)
        puts "\n"
        Logger.new('Received KILL SIGNAL, shutting down...').info_log
        threads.each(&:kill)
        exit
      end

      def execute_command(command, &block)
        IO.popen(command, err: %i[child out]).each(&block)
      end

      def build_backend(directory)
        backend_log('Building image')
        execute_command(
          %W[docker build -f #{directory}/backend/prod.Dockerfile -t #{BACKEND_IMAGE_NAME} #{directory}/backend],
          &method(:backend_log)
        )
      end

      def build_dev_backend(directory)
        backend_log('Building dev image')
        execute_command(
          %W[docker build -f #{directory}/backend/dev.Dockerfile -t #{BACKEND_DEV_IMAGE_NAME} #{directory}/backend],
          &method(:backend_log)
        )
      end

      def start_backend(directory)
        build_backend(directory) unless backend_image_exists?
        execute_command(%W[docker run --rm -p 4560:4560 #{BACKEND_IMAGE_NAME} --verbose], &method(:backend_log))
      end

      def start_dev_backend(directory)
        build_dev_backend(directory) unless backend_dev_image_exists?
        container_intern_path = '/go/src/github.com/orsa-scholis/orsum-inflandi-II/backend'
        execute_command(
          %W[docker run --rm -p 4560:4560 -v #{directory}/backend:#{container_intern_path} #{BACKEND_DEV_IMAGE_NAME}],
          &method(:backend_log)
        )
      end

      def start_frontend
        Dir.chdir('frontend') { execute_command(%w[yarn run start], &method(:frontend_log)) }
      end

      def backend_image_exists?
        system "[ \"$(docker images -q #{BACKEND_IMAGE_NAME} 2>/dev/null)\" != \"\" ]"
      end

      def backend_dev_image_exists?
        system "[ \"$(docker images -q #{BACKEND_DEV_IMAGE_NAME} 2>/dev/null)\" != \"\" ]"
      end

      def frontend_log(log)
        Logger.new(log).frontend_log
      end

      def backend_log(log)
        Logger.new(log).backend_log
      end
    end
  end
end
