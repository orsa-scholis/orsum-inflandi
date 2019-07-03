# typed: true
# frozen_string_literal: true

module OrsumInflandi
  module Commands
    module MockServer
      class CommandLineProcessor
        extend T::Sig

        sig { params(socket: TCPSocket).void }
        def initialize(socket)
          @socket = socket
        end

        sig { void }
        def run
          loop do
            line = STDIN.gets
            break if line.nil?

            @socket.puts line
            log line
          end
        end

        private

        sig { params(message: String).void }
        def log(message)
          Logger.new(message).server_log
        end
      end
    end
  end
end
