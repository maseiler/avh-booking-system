# AVH Booking System

Booking system based on liabilities for [AV Hütte](https://www.av-huette.de/)   

Docker image: [maseiler/avhbs](https://hub.docker.com/repository/docker/maseiler/avhbs)

## Development Environment
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
