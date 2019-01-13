# GasMap
React PWA with GoLang backend. This app's purpose is to find the most economical gas station stops along a road trip. Live version of this app can be found [here](https://gas-map.now.sh)

## Client
### Development
1. Run client in dev mode: `npm start`

### Build and Serve Locally (to test PWA features)
1. Build: `npm run build`
2. Install serve: `npm i -g serve`
3. Serve build: `serve -s build`

### Testing
1. Run all tests: `npm test`

### Deployment
1. `now && now alias`

## Server
### Development
1. Run server: `go run main.go`

### Build
1. Build server: `go build`

### Testing 
1. Run all tests: `go test ./...`
2. Check test coverage: `go test -cover ./...`
3. Run Benchmarks: `go test -bench . ./...`

### Deployment
1. `go build`
2. `git add .`
3. `git commit -m 'Deploying to heroku`
2. `git subtree push --prefix server heroku master`

## Important Resources
- Wireframe: [Here](https://docs.google.com/document/d/1YeSE6cU_osruhf7CHbpamtt6iIiCj2Nl9eezKKZxYaM/edit)
- User Narratives: [Here](https://docs.google.com/document/d/1zTAgrXNFwEfVFGPkSgPuXvTbp4nFH6_5oHCADkD9NEc/edit)
- ERD: [Here](https://docs.google.com/document/d/1g5PqcSkDw_cGNFT1ZePD6GSw9poX2-ZmQVgwLyp711A/edit)
- Presentation Slides: [Here](https://docs.google.com/presentation/d/1dzx4RrsT4Db0ewEdqRW5b46ZPHRLH6IMMQISJY0aXp0/edit)