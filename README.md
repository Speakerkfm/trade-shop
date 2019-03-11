## Trade-shop

## Installation
1. Download project
```bash
$ cd ~/go/src
$ git clone https://github.com/Speakerkfm/trade-shop
```
2. Create database
```bash
$ make migrations
```

3. Install dependencies
```bash
$ dep ensure
```

4. Configure .env
```bash
$ cp .env.dist .env
```

5. Run tests
```bash
$ make clearqueue
$ make test
```

6. Run project
```bash
$ make run
```

## Documentation
```bash
$ make doc
```