# typed: false
# frozen_string_literal: true

require 'google/protobuf'

require_relative '../../proto/Types_pb'
require_relative './base'

module OrsumInflandi
  module Commands
    class Proto < Base
      extend T::Sig

      sig { void }
      def run
        return run_manual if @options[:manual]

        raise NotImplementedError, 'CLI mode not yet implemented'
      end

      private

      sig { returns(T.noreturn) }
      def run_manual
        exit system 'pry -r "google/protobuf" -r "./proto/Types_pb"'
      end
    end
  end
end
