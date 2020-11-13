## Rabbitmq Cluster setup

* Run cluster using docker compose
```bash
docker-compose up -d
```

* Check cluster status using cli or using rabbitmq admin ui
```bash
docker exec -it rabbitmq-01 rabbitmqctl cluster_status
```

* Enable classic queue mirroring & auto node sync
```
docker exec -it rabbitmq-01 rabbitmqctl set_policy ha-fed \
    ".*" '{"ha-mode":"all", "federation-upstream-set":"all", "ha-sync-mode":"automatic", }' \
    --priority 1 \
    --apply-to queues
```

### Create Admin user
```bash
# add user
rabbitmqctl add_user admin admin
# set user as admin
rabbitmqctl set_user_tags admin administrator
# set all permission
rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"
```

