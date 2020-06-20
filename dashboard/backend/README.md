Проксируем запросы к базе <br />
Проверяем запросы к базе <br />
Если что то не так - шлем нахрен

---
Request: 
```json
{
    "params": [
        {
            "type": "key/column",
            "join": "OR/AND",
            "key": "p1",
            "value": "120",
            "operator": ">="
        }
    ],
    "limit": 1000,
    "offset": 5000
}
```

---
Пример получения данных из базы:
```sql
SELECT * FROM internal_logs 
      ARRAY JOIN `params.keys` AS key,     -- только если есть json поля
      arrayEnumerate(`params.keys`) AS idx -- только если есть json поля
    WHERE key='n2' AND arrayElement(`params.values`, idx) LIKE '%v3%'
    ORDER BY time DESC
    LIMIT 5000, 1000;
```

---
Response:
```json
{
    "errors": [
        {
            "code": 123,
            "detail": "details"
        }
    ],
    "data": [
        {
            "time": "",
            "nsec": "",
            "namespace": "",
            "source": "",
            "host": "",
            "trace_id": "",
            "message": "",
            "params": "",
            "build_commit": "",
            "config_hash": ""
        }
    ]
}
```

