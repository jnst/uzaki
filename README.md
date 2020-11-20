## uzaki

A tool to check the incoming status of a product.

## Install

### Node.js

```bash
brew install nodenv
nodenv install -l
nodenv install 15.1.0
nodenv global 15.1.0
nodenv versions
node -v
```

### Serverless

```bash
npm install -g serverless
serverless -v
```

### Docker

Required to perform `serverless invoke local`.

* [Get Started with Docker | Docker](https://www.docker.com/get-started)

## Run

### Local environment

```bash
make local
```

### AWS Lambda

#### Configure aws credentials

```bash
brew install awscli
aws configure 
```

### Deploy to Lambda

```bash
make deploy
```

It will be run periodically at the interval specified in the `rate` in `serverless.yml`.  
Note that the first one will not run immediately and will run after a specified `rate`.
