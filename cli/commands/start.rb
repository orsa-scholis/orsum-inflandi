# typed: true
# frozen_string_literal: true

require 'etc'
require_relative '../logger'
require_relative './base'

module OrsumInflandi
  module Commands
    class Start < Base
      extend T::Sig

      BACKEND_IMAGE_NAME = 'orsa-scholis/orsum-inflandi-ii'
      BACKEND_DEV_IMAGE_NAME = 'orsa-scholis/orsum-inflandi-ii-dev'

      sig { void }
      def run
        initialize_threads

        trap_kill_signal { kill_threads(@threads) }

        @threads.each(&:join)
      end

      private

      sig { void }
      def initialize_threads
        @threads = []
        @threads << backend_thread if @options[:backend]
        @threads << Thread.new(&method(:start_frontend)) if @options[:frontend]
        @threads << Thread.new(&method(:start_frontend)) if @options[:frontend] && @options[:dual_frontend]
      end

      sig { returns(Thread) }
      def backend_thread
        Thread.new(Dir.pwd) { |directory| @options[:dev] ? start_dev_backend(directory) : start_backend(directory) }
      end

      sig { params(command: T::Array[String], block: T.proc.void).void }
      def execute_command(command, &block)
        IO.popen(command, err: %i[child out]).each(&block)
      end

      sig { params(directory: String).void }
      def build_backend(directory)
        backend_log('Building image')
        execute_command(
          %W[docker build -f #{directory}/backend/prod.Dockerfile -t #{BACKEND_IMAGE_NAME} #{directory}/backend],
          &method(:backend_log)
        )
      end

      sig { params(directory: String).void }
      def build_dev_backend(directory)
        backend_log('Building dev image')
        execute_command(
          %W[docker build -f #{directory}/backend/dev.Dockerfile -t #{BACKEND_DEV_IMAGE_NAME} #{directory}/backend],
          &method(:backend_log)
        )
      end

      sig { params(directory: String).void }
      def start_backend(directory)
        build_backend(directory) unless backend_image_exists?
        execute_command(%W[docker run --rm -p 4560:4560 #{BACKEND_IMAGE_NAME} --verbose], &method(:backend_log))
      end

      sig { params(directory: String).void }
      def start_dev_backend(directory)
        build_dev_backend(directory) unless backend_dev_image_exists?
        container_intern_path = '/go/src/github.com/orsa-scholis/orsum-inflandi-II/backend'
        execute_command(
          %W[docker run --rm -p 4560:4560 -v #{directory}/backend:#{container_intern_path} #{BACKEND_DEV_IMAGE_NAME}],
          &method(:backend_log)
        )
      end

      sig { void }
      def start_frontend
        execute_command(%w[yarn --cwd frontend run start], &method(:frontend_log))
      end

      sig { returns(T::Boolean) }
      def backend_image_exists?
        image_exists? BACKEND_IMAGE_NAME
      end

      sig { returns(T::Boolean) }
      def backend_dev_image_exists?
        image_exists? BACKEND_DEV_IMAGE_NAME
      end

      sig { params(name: String).returns(T::Boolean) }
      def image_exists?(name)
        return_code = system "[ \"$(docker images -q #{name} 2>/dev/null)\" != \"\" ]"
        !return_code.nil? && return_code
      end

      sig { params(log: String).void }
      def frontend_log(log)
        Logger.new(log).frontend_log
      end

      sig { params(log: String).void }
      def backend_log(log)
        Logger.new(log).backend_log
      end
    end
  end
end
