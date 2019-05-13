# frozen_string_literal: true

require_relative '../../logger'

module OrsumInflandi
  module Commands
    module MockServer
      class InputProcessor
        def initialize(socket)
          @socket = socket
        end

        def run
          loop { log @socket.gets }
        end

        private

        def log(message)
          Logger.new(message).client_log
        end
      end
    end
  end
end
