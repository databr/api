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
      deputado = find_by_uri(uri).attributes
      REDIS.hset('dep', deputado['uri'], Oj.dump(deputado) )
    end
    deputado
  end

  def self.about(uri)
    deputado = cached(uri)
    about_cached = CACHE.get("a:#{deputado['cadastro_id']}") if ENV['USE_CACHE'] == 'true'
    abouts = Oj.load(about_cached) if about_cached
    unless about_cached
      abouts = About.where(cadastro_id: deputado['cadastro_id'])
      binding.pry
      CACHE.set("a:#{deputado['cadastro_id']}", Oj.dump(abouts.all), ((60)*60)*3 )
    end
    abouts
  end

  def self.allcached
    if ENV['USE_CACHE'] == 'true' && deputados = REDIS.hgetall('dep')
      deputados.map{|i, d| Oj.load(d) }
    else
      Deputado.order("nome_parlamentar").all.map do |deputado|
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
