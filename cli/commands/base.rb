# frozen_string_literal: true

module OrsumInflandi
  module Commands
    class Base
      def initialize(options)
        @options = options
      end

      protected

      def kill_threads(threads)
        puts "\n"
        Logger.new('Received KILL SIGNAL, shutting down...').info_log
        threads.each(&:kill)
        exit
      end

      def trap_kill_signal(&block)
        %w[INT TERM].each do |signal|
          Signal.trap(signal, &block) if block_given?
        end
      end
    end
  end
end
