# encoding: utf-8
#
require './config/boot'

module SocialCamara
  class API < Grape::API
    format :json
    logger LOGGER
    helpers do
      def logger
        LOGGER
      end
    end

    resource :deputados do
      desc 'Returns all deputados'
      get do
        Deputado.allcached
      end

      get ':uri' do
        Deputado.cached_by_uri(params[:uri])
      end

      get ':uri/about' do
        Deputado.about(params[:uri])
      end

      get ':uri/propositions' do
        Deputado.propositions(params[:uri])
      end

      get ':uri/post/:id' do
        type, id = params[:id].split("-")
        data = Deputado.send(type, params[:uri], id)
        present data, with: "#{type.capitalize}Entity".constantize
      end

      get ':uri/feed' do
        deputado = Deputado.cached_by_uri(params[:uri])
        cotas = CotaOnFeedEntity.new(deputado).results
        propositions = PropositionOnFeedEntity.new(deputado).results
        Aggregator.build([cotas, propositions])
      end
    end
  end
end
