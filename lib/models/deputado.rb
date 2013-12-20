class Deputado < ActiveRecord::Base
  validates :nome_parlamentar, presence: true
  validates :cadastro_id, uniqueness: true, presence: true

  has_many :cotas
  has_many :videos

  before_save :set_uri

  def self.cached(uri)
    cached_deputado = REDIS.hget('dep', uri) if ENV['USE_CACHE'] == 'true'
    deputado = Oj.load(cached_deputado) if cached_deputado
    unless cached_deputado
      deputado = find_by_uri(uri)
      REDIS.hset('dep', deputado.uri, Oj.dump(deputado.attributes) )
    end
    deputado
  end

  def self.allcached
    if ENV['USE_CACHE'] == 'true' && deputados = REDIS.hgetall('dep')
      deputados.map{|i, d| Oj.load(d) }
    else
      Deputado.all.map do |deputado|
        REDIS.hset('dep', deputado.uri, Oj.dump(deputado.attributes) )
        deputado
      end
    end
  end

  private
  def set_uri
    self.uri = self.nome_parlamentar.parameterize
  end
end
