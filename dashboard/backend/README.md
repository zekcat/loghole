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
      ARRAY JOIN `params_string.keys` AS key_string,     -- только если есть json поля
      arrayEnumerate(`params_string.keys`) AS idx_string -- только если есть json поля
      ARRAY JOIN `params_float.keys` AS key_float,     -- только если есть json поля
      arrayEnumerate(`params_float.keys`) AS idx_float -- только если есть json поля

    WHERE (key_string='n2' AND arrayElement(`params_string.values`, idx_string) LIKE '%v3%') AND
          key_float='n3' AND arrayElement(`params_float.values`, idx_float) > 10
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

