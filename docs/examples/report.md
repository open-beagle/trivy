# Report Formats

## Table (Default)

```
$ trivy image -f table golang:1.12-alpine
```

## JSON

```
$ trivy image -f json -o results.json golang:1.12-alpine
```

<details>
<summary>Result</summary>

```
2019-05-16T01:46:31.777+0900    INFO    Updating vulnerability database...
2019-05-16T01:47:03.007+0900    INFO    Detecting Alpine vulnerabilities...
```

</details>

<details>
<summary>JSON</summary>

```
[
  {
    "Target": "php-app/composer.lock",
    "Vulnerabilities": null
  },
  {
    "Target": "node-app/package-lock.json",
    "Vulnerabilities": [
      {
        "VulnerabilityID": "CVE-2018-16487",
        "PkgName": "lodash",
        "InstalledVersion": "4.17.4",
        "FixedVersion": "\u003e=4.17.11",
        "Title": "lodash: Prototype pollution in utilities function",
        "Description": "A prototype pollution vulnerability was found in lodash \u003c4.17.11 where the functions merge, mergeWith, and defaultsDeep can be tricked into adding or modifying properties of Object.prototype.",
        "Severity": "HIGH",
        "References": [
          "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16487",
        ]
      }
    ]
  },
  {
    "Target": "trivy-ci-test (alpine 3.7.1)",
    "Vulnerabilities": [
      {
        "VulnerabilityID": "CVE-2018-16840",
        "PkgName": "curl",
        "InstalledVersion": "7.61.0-r0",
        "FixedVersion": "7.61.1-r1",
        "Title": "curl: Use-after-free when closing \"easy\" handle in Curl_close()",
        "Description": "A heap use-after-free flaw was found in curl versions from 7.59.0 through 7.61.1 in the code related to closing an easy handle. ",
        "Severity": "HIGH",
        "References": [
          "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16840",
        ]
      },
      {
        "VulnerabilityID": "CVE-2019-3822",
        "PkgName": "curl",
        "InstalledVersion": "7.61.0-r0",
        "FixedVersion": "7.61.1-r2",
        "Title": "curl: NTLMv2 type-3 header stack buffer overflow",
        "Description": "libcurl versions from 7.36.0 to before 7.64.0 are vulnerable to a stack-based buffer overflow. ",
        "Severity": "HIGH",
        "References": [
          "https://curl.haxx.se/docs/CVE-2019-3822.html",
          "https://lists.apache.org/thread.html/8338a0f605bdbb3a6098bb76f666a95fc2b2f53f37fa1ecc89f1146f@%3Cdevnull.infra.apache.org%3E"
        ]
      },
      {
        "VulnerabilityID": "CVE-2018-16839",
        "PkgName": "curl",
        "InstalledVersion": "7.61.0-r0",
        "FixedVersion": "7.61.1-r1",
        "Title": "curl: Integer overflow leading to heap-based buffer overflow in Curl_sasl_create_plain_message()",
        "Description": "Curl versions 7.33.0 through 7.61.1 are vulnerable to a buffer overrun in the SASL authentication code that may lead to denial of service.",
        "Severity": "HIGH",
        "References": [
          "https://github.com/curl/curl/commit/f3a24d7916b9173c69a3e0ee790102993833d6c5",
        ]
      },
      {
        "VulnerabilityID": "CVE-2018-19486",
        "PkgName": "git",
        "InstalledVersion": "2.15.2-r0",
        "FixedVersion": "2.15.3-r0",
        "Title": "git: Improper handling of PATH allows for commands to be executed from the current directory",
        "Description": "Git before 2.19.2 on Linux and UNIX executes commands from the current working directory (as if '.' were at the end of $PATH) in certain cases involving the run_command() API and run-command.c, because there was a dangerous change from execvp to execv during 2017.",
        "Severity": "HIGH",
        "References": [
          "https://usn.ubuntu.com/3829-1/",
        ]
      },
      {
        "VulnerabilityID": "CVE-2018-17456",
        "PkgName": "git",
        "InstalledVersion": "2.15.2-r0",
        "FixedVersion": "2.15.3-r0",
        "Title": "git: arbitrary code execution via .gitmodules",
        "Description": "Git before 2.14.5, 2.15.x before 2.15.3, 2.16.x before 2.16.5, 2.17.x before 2.17.2, 2.18.x before 2.18.1, and 2.19.x before 2.19.1 allows remote code execution during processing of a recursive \"git clone\" of a superproject if a .gitmodules file has a URL field beginning with a '-' character.",
        "Severity": "HIGH",
        "References": [
          "http://www.securitytracker.com/id/1041811",
        ]
      }
    ]
  },
  {
    "Target": "python-app/Pipfile.lock",
    "Vulnerabilities": null
  },
  {
    "Target": "ruby-app/Gemfile.lock",
    "Vulnerabilities": null
  },
  {
    "Target": "rust-app/Cargo.lock",
    "Vulnerabilities": null
  }
]
```

</details>

`VulnerabilityID`, `PkgName`, `InstalledVersion`, and `Severity` in `Vulnerabilities` are always filled with values, but other fields might be empty.

## Template

### Custom Template

{% raw %}
```
$ trivy image --format template --template "{{ range . }} {{ .Target }} {{ end }}" golang:1.12-alpine
```
{% endraw %}

<details>
<summary>Result</summary>

```
2020-01-02T18:02:32.856+0100    INFO    Detecting Alpine vulnerabilities...
 golang:1.12-alpine (alpine 3.10.2)
```
</details>

You can compute different figures within the template using [sprig][sprig] functions.
As an example you can summarize the different classes of issues:


{% raw %}
```
$ trivy image --format template --template '{{- $critical := 0 }}{{- $high := 0 }}{{- range . }}{{- range .Vulnerabilities }}{{- if  eq .Severity "CRITICAL" }}{{- $critical = add $critical 1 }}{{- end }}{{- if  eq .Severity "HIGH" }}{{- $high = add $high 1 }}{{- end }}{{- end }}{{- end }}Critical: {{ $critical }}, High: {{ $high }}' golang:1.12-alpine
```
{% endraw %}

<details>
<summary>Result</summary>

```
Critical: 0, High: 2
```
</details>

For other features of sprig, see the official [sprig][sprig] documentation.

### Load templates from a file
You can load templates from a file prefixing the template path with an @.

```
$ trivy image --format template --template "@/path/to/template" golang:1.12-alpine
```

### Default Templates
#### XML
In the following example using the template `junit.tpl` XML can be generated.
```
$ trivy image --format template --template "@contrib/junit.tpl" -o junit-report.xml  golang:1.12-alpine
```

#### SARIF
In the following example using the template `sarif.tpl` [Sarif][sarif] can be generated.
```
$ trivy image --format template --template "@contrib/sarif.tpl" -o report.sarif  golang:1.12-alpine
```
This SARIF format can be uploaded to GitHub code scanning results, and there is a [Trivy GitHub Action][action] for automating this process.

Trivy also supports an [ASFF template for reporting findings to AWS Security Hub][asff]

#### HTML

```
$ trivy image --format template --template "@contrib/html.tpl" -o report.html golang:1.12-alpine
```

[action]: https://github.com/aquasecurity/trivy-action
[asff]: https://github.com/aquasecurity/trivy/tree/main/docs/integration/security-hub.md
[sarif]: https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/managing-results-from-code-scanning
[sprig]: http://masterminds.github.io/sprig/
