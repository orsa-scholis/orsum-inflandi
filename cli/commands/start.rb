# frozen_string_literal: true

require_relative '../logger'

module OrsumInflandi
  module Commands
    class Start
      def initialize(options)
        @options = options
      end

      def run
        threads = []
        threads << Thread.new(Dir.pwd, &method(:start_backend)) if @options[:backend]
        threads << Thread.new(&method(:start_frontend)) if @options[:frontend]

        %w[INT TERM].each do |signal|
          Signal.trap(signal) { kill_threads(threads) }
        end

        threads.each(&:join)
      end

      private

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
        execute_command(%W[docker build -t #{BACKEND_IMAGE_NAME} #{directory}/backend], &method(:backend_log))
      end

      def start_backend(directory)
        build_backend(directory) unless backend_image_exists?
        execute_command(%w[docker run --rm -p 4560:4560 orsa-scholis/orsum-inflandi --verbose], &method(:backend_log))
      end

      def start_frontend
        Dir.chdir('frontend') { execute_command(%w[yarn run start], &method(:frontend_log)) }
      end

      def backend_image_exists?
        system "[ \"$(docker images -q #{BACKEND_IMAGE_NAME} 2>/dev/null)\" != \"\" ]"
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
