# typed: true
# frozen_string_literal: true

require 'socket'
require 'sorbet-runtime'

require_relative './base'
require_relative '../logger'
require_relative 'mock_server/input_processor'
require_relative 'mock_server/command_line_processor'

module OrsumInflandi
  module Commands
    class Mock < Base
      extend T::Sig

      sig { params(options: T::Hash[Symbol, Integer]).void }
      def initialize(options)
        super(options)

        @threads = []
      end

      sig { void }
      def run
        trap_kill_signal { kill_threads(@threads) }

        Logger.new('Server started, waiting for connections').info_log
        loop(&method(:main_loop))
      end

      private

      sig { void }
      def main_loop
        client = server.accept
        Logger.new("Accepted connection, #{client.peeraddr.join(', ')}").info_log

        @threads.push(Thread.new { MockServer::InputProcessor.new(client).run })
        @threads.push(Thread.new { MockServer::CommandLineProcessor.new(client).run })
      end

      sig { returns(TCPServer) }
      def server
        @server ||= TCPServer.new @options[:port]
      end
    end
  end
end
