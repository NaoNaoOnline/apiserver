# apiserver

Golang based [Twirp] apiserver. Requires at least `go1.21.1`.



```
./apiserver -h
Golang based gRPC apiserver.

Usage:
  apiserver [flags]
  apiserver [command]

Available Commands:
  daemon      Execute the long running process exposing gRPC server handlers.
  fakeit      Fill redis with fake data.
  systemd     Launch systemd unit files on the executing host.
  version     Print version information of this command line tool.

Flags:
  -h, --help   help for apiserver

Use "apiserver [command] --help" for more information about a command.
```



Running redis stack locally.

```
docker run --rm --name redis-stack-apiserver -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



Filling the apiserver with test data.

```
./apiserver fakeit
```



Running the apiserver locally.

```
./apiserver daemon
```

```
{ "time":"2023-10-02 21:08:10", "leve":"info", "mess":"worker searching for tasks", "addr":"127.0.0.1:6379", "call":"/Users/xh3b4sd/project/NaoNaoOnline/apiserver/pkg/worker/worker.go:50" }
{ "time":"2023-10-02 21:08:10", "leve":"info", "mess":"server listening for calls", "addr":"127.0.0.1:7777", "call":"/Users/xh3b4sd/project/NaoNaoOnline/apiserver/pkg/server/server.go:69" }
```



Note that a `.env.local` file is required for the apiserver to be configured
properly.

```
APISERVER_CHAIN_CID="31337"
APISERVER_CHAIN_POL="0x5FbDB2315678afecb367f032d93F642f64180aa3"
APISERVER_CHAIN_RPC="http://127.0.0.1:8545"
APISERVER_CHAIN_SUB="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"

APISERVER_MSUBS_EVE="10"
APISERVER_MSUBS_LIN="15"

APISERVER_OAUTH_ISS="https://[REDACTED].auth0.com/"
APISERVER_OAUTH_AUD="https://apiserver"

APISERVER_UPREM_TIM="2023-11-01T00:00:00Z"
```



Calling the apiserver locally.

```
curl -s --request "POST" --header "Content-Type: application/json" --data '{"object":[{"intern":{"user":"1699746343006008"}}]}' http://127.0.0.1:7777/event.API/Search | jq '.object[0]'
```

```
{
  "extern": [
    {
      "amnt": "43",
      "kind": "link",
      "user": false
    }
  ],
  "intern": {
    "crtd": "1699799790",
    "evnt": "1699799790111846",
    "user": "1699746343006008"
  },
  "public": {
    "cate": "1699746125560859,1699746343054161,1699746006668458,",
    "dura": "5400",
    "host": "1699746006088467,1699746346734161,1699746125979691",
    "link": "https://internationalinfomediaries.com.org",
    "time": "1699804800"
  }
}
```



Given the following environment variables, some automation in combination with
the Twitter API might be enabled, e.g. creating tweets for accounts associated
to events created on NaoNao. The `API_KEY` and `ACCESS_TOKEN` set here is for
the same twitter account. Creating posts is free, meaning there is no need for a
paid subscription plan.

```
export GOTWI_ACCESS_TOKEN=$(        cat ~/.credential/apiserver-twitter-acc-key    )
export GOTWI_ACCESS_TOKEN_SECRET=$( cat ~/.credential/apiserver-twitter-acc-secret )
export GOTWI_API_KEY=$(             cat ~/.credential/apiserver-twitter-api-key    )
export GOTWI_API_KEY_SECRET=$(      cat ~/.credential/apiserver-twitter-api-secret )
```



Generating smart contract bindings using [abigen].

```
abigen --abi pkg/contract/policycontract/Policy.ABI.json --pkg policycontract --type Policy --out pkg/contract/policycontract/policy_contract.go
```

```
abigen --abi pkg/contract/subscriptioncontract/Subscription.ABI.json --pkg subscriptioncontract --type Subscription --out pkg/contract/subscriptioncontract/subscription_contract.go
```


Running conformance tests.


```
export REDIS_PORT=6382
```

```
docker run --rm -p 6382:6379 -p 8082:8001 redis/redis-stack:latest
```

```
go test ./cnf --tags redis -count 1 -race
```



[abigen]: https://geth.ethereum.org/docs/tools/abigen
[Twirp]: https://github.com/twitchtv/twirp
