package main

import (
    "domestic_system/libs/wallet"

)

func main() {

    wallet.List()

    /*values := map[string]string{"jsonrpc":"2.0","method":"eth_accounts","id":"1"}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://127.0.0.1:8545", "application/json",
        bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    aInterface := res["result"].([]interface{})
    aString := make([]string, len(aInterface))
    for i, v := range aInterface {
        aString[i] = v.(string)
    }

    p := tea.NewProgram(initialModel(aString))
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }*/
}
