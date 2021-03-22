# web-scraping

## Instruções para uso da aplicação:

#### Requisitos:
  - É necessário uma conexão com a internet;
  - É necessário ter o Git instalado em sua máquina;
  - É necessário ter o Go instalado em sua máquina, caso não tenha visite https://golang.org/dl/ para fazer o download.
  
## Como iniciar a aplicação:
  - Faça o clone do respositório em sua máquina (o diretório é de sua preferência);
  - Usando um terminal (windows com gitbash ou linux com qualquer terminal shell) abra o diretório do projeto;
  - Utilize o comando `go get -d -v ./src/` para fazer o download das dependências necessárias;
  
  <span color="orange">Obs</span>: Os próximos passos devem ser feitos até o momento em que o problema da dependência google-search esteja resolvido (https://github.com/rocketlaunchr/google-search/pull/7/files)
  
  - Procure pelo Go Workspace digitando o comando `go env GOPATH`
  - Vá até o diretório mostrado no terminal e procure pela dependência citada acima, exemplo:
  `$PATH/go/pkg/mod/github.com/rocketlaunchr/google-search@v1.1.1`
  - Usando um editor de texto abra o arquivo `search.go`
  - Nele você deve procurar o seguinte bloco de código:
  ```
  // https://www.w3schools.com/cssref/css_selectors.asp
	c.OnHTML("div.g", func(e *colly.HTMLElement) {

		sel := e.DOM

		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("div > div > a > h3 > span").Text()) <---- irá alterar esta linha

		descText := strings.TrimSpace(sel.Find("div > div > div > span > span").Text())

		if linkText != "" && linkText != "#" {
			result := Result{
				Rank:        rank,
				URL:         linkText,
				Title:       titleText,
				Description: descText,
			}
			results = append(results, result)
			rank += 1
		}
	})
  ```
  - Remova o elemento 'span' do método `sel.Find("div > div > a > h3 > span")` como neste pull request: https://github.com/rocketlaunchr/google-search/pull/7/files;
  - Salve o arquivo;
  
<span color="orange">Obs</span>: Neste momento não é preciso seguir os passos acima caso o problema esteja resolvido;

  - Vá até o diretório do projeto e digite o comando `go install ./src/`
  - Rode a aplicação com o comando `go run ./src/`
  - O servidor estará rodando em http://localhost:8080/home/
