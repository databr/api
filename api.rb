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
        Deputado.cached(params[:uri])
      end

      get ':uri/about' do
        Deputado.about(params[:uri])
      end

      get ':uri/propositions' do
        deputado = Deputado.cached(params[:uri])
        Proposition.order('presentations_at DESC').where(cadastro_id: deputado.cadastro_id)
      end

      get ':uri/feed' do
        deputado = Deputado.cached(params[:uri])
        cotas = CotaEntity.new(deputado).results
        propositions = PropositionEntity.new(deputado).results
        Aggregator.build([cotas, propositions])
      end
    end
  end
end
