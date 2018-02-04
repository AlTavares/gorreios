package correios

type TrackingInfo struct {
	Versao     string   `json:"versao"`
	Quantidade string   `json:"quantidade"`
	Pesquisa   string   `json:"pesquisa"`
	Resultado  string   `json:"resultado"`
	Objeto     []Objeto `json:"objeto"`
}

type Objeto struct {
	Numero    string   `json:"numero"`
	Sigla     string   `json:"sigla"`
	Nome      string   `json:"nome"`
	Categoria string   `json:"categoria"`
	Evento    []Evento `json:"evento"`
}

type Evento struct {
	CodigoServico string     `json:"codigoServico,omitempty"`
	CepDestino    string     `json:"cepDestino"`
	DiasUteis     string     `json:"diasUteis"`
	DataPostagem  string     `json:"dataPostagem"`
	Tipo          string     `json:"tipo"`
	Status        string     `json:"status"`
	Data          string     `json:"data"`
	Hora          string     `json:"hora"`
	Criacao       string     `json:"criacao"`
	Descricao     string     `json:"descricao"`
	Unidade       Unidade    `json:"unidade"`
	DetalheOEC    DetalheOEC `json:"detalheOEC,omitempty"`
	Destino       []Destino  `json:"destino,omitempty"`
	Postagem      Postagem   `json:"postagem,omitempty"`
	PrazoGuarda   string     `json:"prazoGuarda"`
}
type Unidade struct {
	Local       string   `json:"local"`
	Codigo      string   `json:"codigo"`
	Cidade      string   `json:"cidade"`
	Uf          string   `json:"uf"`
	Sto         string   `json:"sto"`
	TipoUnidade string   `json:"tipounidade"`
	Endereco    Endereco `json:"endereco"`
}
type Endereco struct {
	Codigo     string `json:"codigo"`
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Bairro     string `json:"bairro"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

type DetalheOEC struct {
	Carteiro string `json:"carteiro"`
	Distrito string `json:"distrito"`
	Lista    string `json:"lista"`
	Unidade  string `json:"unidade"`
}
type Destino struct {
	Local    string   `json:"local"`
	Codigo   string   `json:"codigo"`
	Cidade   string   `json:"cidade"`
	Bairro   string   `json:"bairro"`
	Uf       string   `json:"uf"`
	Endereco Endereco `json:"endereco"`
}
type Postagem struct {
	CepDestino      string `json:"cepdestino"`
	Ar              string `json:"ar"`
	Mp              string `json:"mp"`
	Dh              string `json:"dh"`
	Peso            string `json:"peso"`
	Volume          string `json:"volume"`
	DataProgramada  string `json:"dataprogramada"`
	DataPostagem    string `json:"datapostagem"`
	PrazoTratamento string `json:"prazotratamento"`
	CodigoServico   string `json:"codigoservico"`
}
