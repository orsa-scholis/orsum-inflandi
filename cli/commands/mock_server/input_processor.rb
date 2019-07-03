# typed: true
# frozen_string_literal: true

require_relative '../../logger'

module OrsumInflandi
  module Commands
    module MockServer
      class InputProcessor
        extend T::Sig

        sig { params(socket: TCPSocket).void }
        def initialize(socket)
          @socket = socket
        end

        sig { void }
        def run
          loop { log @socket.gets }
        end

        private

        sig { params(message: String).void }
        def log(message)
          Logger.new(message).client_log
        end
      end
    end
  end
end
