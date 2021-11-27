# Supported OS

The unfixed/unfixable vulnerabilities mean that the patch has not yet been provided on their distribution. Trivy doesn't support self-compiled packages/binaries, but official packages provided by vendors such as Red Hat and Debian.

| OS                           | Supported Versions                       | Target Packages               | Detection of unfixed vulnerabilities |
| ---------------------------- | ---------------------------------------- | ----------------------------- | :----------------------------------: |
| Alpine Linux                 | 2.2 - 2.7, 3.0 - 3.13                    | Installed by apk              |                  NO                  |
| Red Hat Universal Base Image | 7, 8                                     | Installed by yum/rpm          |                 YES                  |
| Red Hat Enterprise Linux     | 6, 7, 8                                  | Installed by yum/rpm          |                 YES                  |
| CentOS                       | 6, 7                                     | Installed by yum/rpm          |                 YES                  |
| Oracle Linux                 | 5, 6, 7, 8                               | Installed by yum/rpm          |                  NO                  |
| Amazon Linux                 | 1, 2                                     | Installed by yum/rpm          |                  NO                  |
| openSUSE Leap                | 42, 15                                   | Installed by zypper/rpm       |                  NO                  |
| SUSE Enterprise Linux        | 11, 12, 15                               | Installed by zypper/rpm       |                  NO                  |
| Photon OS                    | 1.0, 2.0, 3.0                            | Installed by tdnf/yum/rpm     |                  NO                  |
| Debian GNU/Linux             | wheezy, jessie, stretch, buster          | Installed by apt/apt-get/dpkg |                 YES                  |
| Ubuntu                       | 12.04, 14.04, 16.04, 18.04, 18.10, 19.04 | Installed by apt/apt-get/dpkg |                 YES                  |
| Distroless*                  | Any                                      | Installed by apt/apt-get/dpkg |                 YES                  |

*Distroless: https://github.com/GoogleContainerTools/distroless
