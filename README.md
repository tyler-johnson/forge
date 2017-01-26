
### internal databases

- accounts - users that control apps
  - apps - an overall single CRUD app
    - configs - versioned config files
    - services - running api at config version with dedicated resources (db, nodes)
    - instances - specific instance of running service on a node

### forge parts
- smart proxy - balances requests to nodes running the desired apps
- leader - ochestrates and balances apps across connected agents
- agent - starts and runs nodes for service
- node - http server running config business logic
- admin - gui for managing apps and exploring data

### external parts
- cockroachdb - normal, small data
- consul - leader election, fault tolerance, node detection
- seaweedfs - file storage

### workflow
- make repo with `.fg` config files
- when ready run `forge publish my-forge-app`
- use cli or gui to create running services for app