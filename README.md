# Telegram_bot
This project main target is use docker deploy project

## Setup
1. `docker compose up`

## Docker-compose
1. Telegram_bot_main
   + DATABASE_HOST
   + TELEGRAM_BOT_MODE
2. Telegram_bot_postgres 
   + PGDATA
3. Telegram_bot_pgadmin

## ENV(Default)
1. DATABASE_HOST=telegram-bot-postgres
2. TELEGRAM_BOT_MODE=prod
3. PGDATA=/var/lib/postgresql/data
4. PGADMIN_DEFAULT_EMAIL=root@gmail.com
5. PGADMIN_DEFAULT_PASSWORD=root