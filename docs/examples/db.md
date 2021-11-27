# Vulnerability DB

## Skip update of vulnerability DB
`Trivy` downloads its vulnerability database every 12 hours when it starts operating.
This is usually fast, as the size of the DB is only 10~30MB.
But if you want to skip even that, use the `--skip-update` option.

```
$ trivy image --skip-update python:3.4-alpine3.9
```

<details>
<summary>Result</summary>

```
2019-05-16T12:48:08.703+0900    INFO    Detecting Alpine vulnerabilities...

python:3.4-alpine3.9 (alpine 3.9.2)
===================================
Total: 1 (UNKNOWN: 0, LOW: 0, MEDIUM: 1, HIGH: 0, CRITICAL: 0)

+---------+------------------+----------+-------------------+---------------+--------------------------------+
| LIBRARY | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION |             TITLE              |
+---------+------------------+----------+-------------------+---------------+--------------------------------+
| openssl | CVE-2019-1543    | MEDIUM   | 1.1.1a-r1         | 1.1.1b-r1     | openssl: ChaCha20-Poly1305     |
|         |                  |          |                   |               | with long nonces               |
+---------+------------------+----------+-------------------+---------------+--------------------------------+
```

</details>

## Only download vulnerability database
You can also ask `Trivy` to simply retrieve the vulnerability database.
This is useful to initialize workers in Continuous Integration systems.

```
$ trivy image --download-db-only
```

## Lightweight DB
The lightweight DB doesn't contain vulnerability detail such as descriptions and references. Because of that, the size of the DB is smaller and the download is faster.

This option is useful when you don't need vulnerability details and is suitable for CI/CD.
To find the additional information, you can search vulnerability details on the NVD website.
https://nvd.nist.gov/vuln/search

```
$ trivy image --light alpine:3.10
```

`--light` option doesn't display titles like the following example.

<details>
<summary>Result</summary>

```
2019-11-14T10:21:01.553+0200    INFO    Reopening vulnerability DB
2019-11-14T10:21:02.574+0200    INFO    Detecting Alpine vulnerabilities...

alpine:3.10 (alpine 3.10.2)
===========================
Total: 3 (UNKNOWN: 0, LOW: 1, MEDIUM: 2, HIGH: 0, CRITICAL: 0)

+---------+------------------+----------+-------------------+---------------+
| LIBRARY | VULNERABILITY ID | SEVERITY | INSTALLED VERSION | FIXED VERSION |
+---------+------------------+----------+-------------------+---------------+
| openssl | CVE-2019-1549    | MEDIUM   | 1.1.1c-r0         | 1.1.1d-r0     |
+         +------------------+          +                   +               +
|         | CVE-2019-1563    |          |                   |               |
+         +------------------+----------+                   +               +
|         | CVE-2019-1547    | LOW      |                   |               |
+---------+------------------+----------+-------------------+---------------+
```
</details>
