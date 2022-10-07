# AVH Booking System

Booking system based on liabilities for [AV Hütte](https://www.av-huette.de/)   

Docker image: [maseiler/avhbs](https://hub.docker.com/repository/docker/maseiler/avhbs)

## Development
Prerequisites
- at least Go version [19.2](https://go.dev/dl/)
- [NodeJS](https://nodejs.org/en/)

### Installation
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@v2.0.0

cd client
go mod tidy
go install
npm install

cd ../server
go mod tidy
go install
```

### Start Hacking
Terminal 1
```bash
cd client/frontend
npm install
npm run dev
```

Terminal 2
```bash
cd client
wails dev
```

Terminal 3
```bash
cd server
go run .
```

## Build
### Client
```bash
cd client
wails Build
```

### Server
```bash
cd server
go Build
```
