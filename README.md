# Delivery Bot API for Telegram

Periodically check supermarket onlin delivery slots for availability, and notify users on telegram

## Table of Contents

- [Usage](#usage)
- [Maintainers](#maintainers)
- [License](#license)

## Usage

1. Create your own .env using .sample.env

2. Start docker containers

```
make up
```

3. View logs

```
make logs
```

4. Visit `localhost:4000/some` to check if API is responding

5. Generate docs from swagger comments

```
make generate-docs
```

6. Visit `localhost:4000/docs` for documentation if `DOCS=true` in .env

7. Stop docker containers

```
make down
```

## Test

```
make test
```

## Maintainers

[@gpng](https://github.com/gpng)

## License

MIT © 2018 Gerald Png
