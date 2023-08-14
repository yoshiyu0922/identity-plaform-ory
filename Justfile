run:
    docker compose -f db/compose.yml -f api/compose.yml -f kratos/compose.yml -f hydra/compose.yml up --build --force-recreate
