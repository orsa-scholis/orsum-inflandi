# frozen_string_literal: true

module OrsumInflandi
  module Commands
    module MockServer
      class CommandLineProcessor
        def initialize(socket)
          @socket = socket
        end

        def run
          loop do
            line = STDIN.gets
            @socket.puts line
            log line
          end
        end

        private

        def log(message)
          Logger.new(message).server_log
        end
      end
    end
  end
end
