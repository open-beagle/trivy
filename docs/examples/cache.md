# Cache

## Clear Caches
The `--clear-cache` option removes caches.

**The scan is not performed.**

```
$ trivy image --clear-cache
```

<details>
<summary>Result</summary>

```
2019-11-15T15:13:26.209+0200    INFO    Reopening vulnerability DB
2019-11-15T15:13:26.209+0200    INFO    Removing image caches...
```

</details>

## Cache Directory
Specify where the cache is stored with `--cache-dir`.

```
$ trivy --cache-dir /tmp/trivy/ image python:3.4-alpine3.9
```

## Cache Backend
[EXPERIMENTAL] This feature might change without preserving backwards compatibility.

Trivy supports local filesystem and Redis as the cache backend. This option is useful especially for client/server mode.

Two options:
- `fs`
    - the cache path can be specified by `--cache-dir`
- `redis://`
    - `redis://[HOST]:[PORT]`

```
$ trivy server --cache-backend redis://localhost:6379
```
