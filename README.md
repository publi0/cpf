# Consulte CPF info via terminal

Consulte CPF pelo terminal

## Instalação

### Go install

```bash
go install github.com/publi0/cpf@latest
```

### From source

```bash
go build -o cpf main.go
```

### From release

<https://github.com/publi0/cpf/releases>

## Uso

```bash
./cpf 00000000000000
```

## Saída

### Por padrão o output é o yaml

```yaml
cpf: 00000
nome: Jose Silva
sexo: M
nascimento: 1973-08-31
mae: Maria Silva
idade: 50 anos
```

## Flags

### --ouput -o

Retorna o output em json ou yaml

```bash
./cpf 00000000000000 --output json
```

### --raw -r

Retorna o output em texto puro

```bash
./cpf 00000000000000 --raw
```

## Fonte de dados

URL base: <https://consultanacional.org/>
