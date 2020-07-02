package main

func main() {
	m := make(map[string]interface{})

	// decode(m) что то там.

	level := m["level"]
	time := m["time"]
	message := m["message"]
	build_commit := m["build_commit"]
	config_hash := m["config_hash"]
	host := m["host"]
	source := m["source"]
	namespace := m["namespace"]
	action := m["action"]

	var (
		params_string_keys   = []string{}
		params_string_values = []string{}
		params_float_keys    = []string{}
		params_float_values  = []float64{}
	)

	for key, value := range m {
		if isDefault(key) {
			continue
		}

		
	}
}

func isDefault(key string) bool {
	m := map[string]struct{}{
		"level": struct{}{},
		"time": struct{}{},
		"message": struct{}{},
		"build_commit": struct{}{},
		"config_hash": struct{}{},
		"host": struct{}{},
		"source": struct{}{},
		"namespace": struct{}{},
		"action": struct{}{},
	}

	_, ok := m[key]

	return ok
}

/*
    `params_string.keys` Array(String),
    `params_string.values` Array(String),

{
   "level":"ERROR",                                           ---
   "time":"2020-06-28T17:46:18+03:00",                        ---
   "caller":"app1/main.go:125",                               +++
   "message":"update order failed: some message 4",           ---
   "build_commit":"db957a22b3c1d6e508c0828917a5e14c572fb007", ---
   "config_hash":"130362a6fd10cf2f939dd0cfc0ab222cee6a99ec",  ---
   "host":"127.100.0.1:50000",                                ---
   "source":"app_1",                                          ---
   "namespace":"prod",                                        ---
   "action":"1c75sek3g3jn6pl",                                ---
   "some_type": {                                             +++
	   "key1": "aaaa",
	   "key2": [1, 2, 3, 4],
	   "key3": {
		   "key3_key1": "AAAAAAAAAAAAA!!!!!!"
	   }
   }
}

params = {ALL JSON}

params_string.keys   = [caller, key1, key3_key1]
params_string.values = [app1/main.go:125, "aaaa", "AAAAAAAAAAAAA!!!!!!"]

params_float.keys   = [key2, key2, key2, key2]
params_float.values = [1, 2 ,3, 4]

*/

