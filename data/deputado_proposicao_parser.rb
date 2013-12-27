# encoding: utf-8
class DeputadoProposicaoParser < CamaraParser
  def initialize(deputado, year = Time.now.year, sigla = 'PL')
    part_name = deputado.nome_parlamentar.split(' ').first
    @url = "http://www.camara.gov.br/SitCamaraWS/Proposicoes.asmx/ListarProposicoes?sigla=#{sigla}&numero=&ano=#{year}&datApresentacaoIni=&datApresentacaoFim=&autor=&parteNomeAutor=#{part_name}&siglaPartidoAutor=&siglaUFAutor=&generoAutor=&codEstado=&codOrgaoEstado=&emTramitacao=&idTipoAutor="
    super()
  end

  def propositions
    @parser.search('//proposicao').map do |proposition|
      date, time = (proposition/'datapresentacao').text.split(' ')
      day, month, year = date.split('/')
      hour, minute, second = (time || '0:0').split(':')
      presentations_at = Time.new(year.to_i, month.to_i, day.to_i, hour.to_i, minute.to_i, second.to_i, '-03:00')
      {
        proposition_id: (proposition/'id').first.text,
        presentations_at: presentations_at,
        year: (proposition/'ano').text,
        number: (proposition/'numero').first.text,
        name: (proposition/'nome').first.text,
        body: (proposition/'txtementa').text,
        cadastro_id: (proposition/'autor1/idecadastro').text,
        url: "http://www.camara.gov.br/proposicoesWeb/
        fichadetramitacao?idProposicao=#{(proposition/'id').first.text}"
      }
    end
  end
end
