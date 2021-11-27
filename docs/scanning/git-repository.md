# Git Repository

Scan your remote git repository

```
$ trivy repo https://github.com/knqyf263/trivy-ci-test
```

Only public repositories are supported.

<details>
<summary>Result</summary>

```
2021-03-09T15:04:19.003+0200    INFO    Detecting cargo vulnerabilities...
2021-03-09T15:04:19.005+0200    INFO    Detecting pipenv vulnerabilities...

Cargo.lock
==========
Total: 7 (UNKNOWN: 7, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)

+----------+-------------------+----------+-------------------+------------------------------+---------------------------------------------+
| LIBRARY  | VULNERABILITY ID  | SEVERITY | INSTALLED VERSION |        FIXED VERSION         |                    TITLE                    |
+----------+-------------------+----------+-------------------+------------------------------+---------------------------------------------+
| ammonia  | RUSTSEC-2019-0001 | UNKNOWN  | 1.9.0             | >= 2.1.0                     | Uncontrolled recursion leads                |
|          |                   |          |                   |                              | to abort in HTML serialization              |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2019-0001 |
+----------+-------------------+          +-------------------+------------------------------+---------------------------------------------+
| openssl  | RUSTSEC-2016-0001 |          | 0.8.3             | >= 0.9.0                     | SSL/TLS MitM vulnerability                  |
|          |                   |          |                   |                              | due to insecure defaults                    |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2016-0001 |
+----------+-------------------+          +-------------------+------------------------------+---------------------------------------------+
| smallvec | RUSTSEC-2018-0018 |          | 0.6.9             | >= 0.6.13                    | smallvec creates uninitialized              |
|          |                   |          |                   |                              | value of any type                           |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2018-0018 |
+          +-------------------+          +                   +------------------------------+---------------------------------------------+
|          | RUSTSEC-2019-0009 |          |                   | >= 0.6.10                    | Double-free and use-after-free              |
|          |                   |          |                   |                              | in SmallVec::grow()                         |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2019-0009 |
+          +-------------------+          +                   +                              +---------------------------------------------+
|          | RUSTSEC-2019-0012 |          |                   |                              | Memory corruption in SmallVec::grow()       |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2019-0012 |
+          +-------------------+          +                   +------------------------------+---------------------------------------------+
|          | RUSTSEC-2021-0003 |          |                   | >= 0.6.14, < 1.0.0, >= 1.6.1 | Buffer overflow in SmallVec::insert_many    |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2021-0003 |
+----------+-------------------+          +-------------------+------------------------------+---------------------------------------------+
| tempdir  | RUSTSEC-2018-0017 |          | 0.3.7             |                              | `tempdir` crate has been                    |
|          |                   |          |                   |                              | deprecated; use `tempfile` instead          |
|          |                   |          |                   |                              | -->rustsec.org/advisories/RUSTSEC-2018-0017 |
+----------+-------------------+----------+-------------------+------------------------------+---------------------------------------------+

Pipfile.lock
============
Total: 20 (UNKNOWN: 3, LOW: 0, MEDIUM: 7, HIGH: 5, CRITICAL: 5)

+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
|       LIBRARY       | VULNERABILITY ID | SEVERITY | INSTALLED VERSION |     FIXED VERSION      |                 TITLE                 |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| django              | CVE-2019-19844   | CRITICAL | 2.0.9             | 3.0.1, 2.2.9, 1.11.27  | Django: crafted email address         |
|                     |                  |          |                   |                        | allows account takeover               |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-19844 |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-7471    |          |                   | 3.0.3, 2.2.10, 1.11.28 | django: potential SQL injection       |
|                     |                  |          |                   |                        | via StringAgg(delimiter)              |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-7471  |
+                     +------------------+----------+                   +------------------------+---------------------------------------+
|                     | CVE-2019-6975    | HIGH     |                   | 2.1.6, 2.0.11, 1.11.19 | python-django: memory exhaustion in   |
|                     |                  |          |                   |                        | django.utils.numberformat.format()    |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-6975  |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-9402    |          |                   | 3.0.4, 2.2.11, 1.11.29 | django: potential SQL injection       |
|                     |                  |          |                   |                        | via "tolerance" parameter in          |
|                     |                  |          |                   |                        | GIS functions and aggregates...       |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-9402  |
+                     +------------------+----------+                   +------------------------+---------------------------------------+
|                     | CVE-2019-3498    | MEDIUM   |                   | 2.1.5, 2.0.10, 1.11.18 | python-django: Content spoofing       |
|                     |                  |          |                   |                        | via URL path in default 404 page      |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-3498  |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-13254   |          |                   | 3.0.7, 2.2.13          | django: potential data leakage        |
|                     |                  |          |                   |                        | via malformed memcached keys          |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-13254 |
+                     +------------------+          +                   +                        +---------------------------------------+
|                     | CVE-2020-13596   |          |                   |                        | django: possible XSS via              |
|                     |                  |          |                   |                        | admin ForeignKeyRawIdWidget           |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-13596 |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| django-cors-headers | pyup.io-37132    | UNKNOWN  | 2.5.2             | 3.0.0                  | In django-cors-headers                |
|                     |                  |          |                   |                        | version 3.0.0,                        |
|                     |                  |          |                   |                        | ``CORS_ORIGIN_WHITELIST``             |
|                     |                  |          |                   |                        | requires URI schemes, and             |
|                     |                  |          |                   |                        | optionally ports. This...             |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| djangorestframework | CVE-2020-25626   | MEDIUM   | 3.9.2             | 3.11.2                 | django-rest-framework: XSS            |
|                     |                  |          |                   |                        | Vulnerability in API viewer           |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-25626 |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| httplib2            | CVE-2021-21240   | HIGH     | 0.12.1            | 0.19.0                 | python-httplib2: Regular              |
|                     |                  |          |                   |                        | expression denial of                  |
|                     |                  |          |                   |                        | service via malicious header          |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2021-21240 |
+                     +------------------+----------+                   +------------------------+---------------------------------------+
|                     | CVE-2020-11078   | MEDIUM   |                   | 0.18.0                 | python-httplib2: CRLF injection       |
|                     |                  |          |                   |                        | via an attacker controlled            |
|                     |                  |          |                   |                        | unescaped part of uri for...          |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-11078 |
+                     +------------------+----------+                   +                        +---------------------------------------+
|                     | pyup.io-38303    | UNKNOWN  |                   |                        | Httplib2 0.18.0 is an                 |
|                     |                  |          |                   |                        | important security update to          |
|                     |                  |          |                   |                        | patch a CWE-93 CRLF...                |
+---------------------+------------------+          +-------------------+------------------------+---------------------------------------+
| jinja2              | pyup.io-39525    |          | 2.10.1            | 2.11.3                 | This affects the package              |
|                     |                  |          |                   |                        | jinja2 from 0.0.0 and before          |
|                     |                  |          |                   |                        | 2.11.3. The ReDOS...                  |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| py                  | CVE-2020-29651   | HIGH     | 1.8.0             |                        | python-py: ReDoS in the py.path.svnwc |
|                     |                  |          |                   |                        | component via mailicious input        |
|                     |                  |          |                   |                        | to blame functionality...             |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-29651 |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| pyyaml              | CVE-2019-20477   | CRITICAL |               5.1 |                        | PyYAML: command execution             |
|                     |                  |          |                   |                        | through python/object/apply           |
|                     |                  |          |                   |                        | constructor in FullLoader             |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-20477 |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-14343   |          |                   |                    5.4 | PyYAML: incomplete                    |
|                     |                  |          |                   |                        | fix for CVE-2020-1747                 |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-14343 |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-1747    |          |                   | 5.3.1                  | PyYAML: arbitrary command             |
|                     |                  |          |                   |                        | execution through python/object/new   |
|                     |                  |          |                   |                        | when FullLoader is used               |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-1747  |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
| urllib3             | CVE-2019-11324   | HIGH     | 1.24.1            | 1.24.2                 | python-urllib3: Certification         |
|                     |                  |          |                   |                        | mishandle when error should be thrown |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-11324 |
+                     +------------------+----------+                   +------------------------+---------------------------------------+
|                     | CVE-2019-11236   | MEDIUM   |                   |                        | python-urllib3: CRLF injection        |
|                     |                  |          |                   |                        | due to not encoding the               |
|                     |                  |          |                   |                        | '\r\n' sequence leading to...         |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2019-11236 |
+                     +------------------+          +                   +------------------------+---------------------------------------+
|                     | CVE-2020-26137   |          |                   | 1.25.9                 | python-urllib3: CRLF injection        |
|                     |                  |          |                   |                        | via HTTP request method               |
|                     |                  |          |                   |                        | -->avd.aquasec.com/nvd/cve-2020-26137 |
+---------------------+------------------+----------+-------------------+------------------------+---------------------------------------+
```

</details>
