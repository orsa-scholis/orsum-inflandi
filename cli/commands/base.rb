# typed: true
# frozen_string_literal: true

require 'sorbet-runtime'

module OrsumInflandi
  module Commands
    class Base
      extend T::Sig

      sig { params(options: T::Hash[Symbol, T.untyped]).void }
      def initialize(options)
        @options = options
      end

      protected

      sig { params(threads: T::Array[Thread]).returns(T.noreturn) }
      def kill_threads(threads)
        puts "\n"
        Logger.new('Received KILL SIGNAL, shutting down...').info_log
        threads.each(&:kill)
        exit
      end

      sig { params(block: T.nilable(T.proc.params(arg0: Integer).returns(BasicObject))).void }
      def trap_kill_signal(&block)
        %w[INT TERM].each do |signal|
          Signal.trap(signal, &block) if block_given?
        end
      end
    end
  end
end
