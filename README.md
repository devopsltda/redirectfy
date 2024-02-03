# Redirect

## Como configurar o ambiente de desenvolvimento

1. Edite o `php.ini` do seu sistema

- No Windows, fica em alguma dessas pastas:
    - `C:\php\`
    - `C:\xampp\php\`
    - `C:\wamp\bin\php\`
    - `C:\Program Files\PHP\`
- No Linux, fica em `/etc/php/php.ini`

2. Descomente as seguintes linhas:

```ini
;extension=pdo_sqlite
;extension=sqlite3
```

3. Execute os comandos abaixo:

```bash
git clone git@github.com:De-v-0ps/redirect.git # Clona o repositório

cp .env.example .env # Copia as configurações padrão

php artisan sail:install # Instala o sail (Docker compose do Laravel)

./vendor/bin/sail up # Inicia o serviço Docker

# ou

./vendor/bin/sail up -d # Inicia o serviço Docker em background
```

### Tá dizendo que não tenho permissão pra acessar um arquivo X

No linux, use esse comando na pasta raiz do projeto:

```bash
chmod 777 ./**/* ./**/.*
```

### Tá dizendo que não tenho o driver do PHP

No Linux, é só instalar o driver:

- Debian ou Ubuntu

```bash
sudo apt install php-sqlite3
```

- Arch

```bash
sudo pacman -S php-sqlite
```
