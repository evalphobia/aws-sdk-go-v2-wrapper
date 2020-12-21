aws-sdk-go-v2-wrapper
----

[![License: MIT][401]][402] [![GoDoc][101]][102] [![Release][103]][104] [![Build Status][201]][202] [![Coveralls Coverage][203]][204] [![Codecov Coverage][205]][206]
[![Go Report Card][301]][302] [![Code Climate][303]][304] [![BCH compliance][305]][306] [![CodeFactor][307]][308] [![codebeat][309]][310] [![Scrutinizer Code Quality][311]][312] [![FOSSA Status][403]][404]


<!-- Basic -->

[101]: https://godoc.org/github.com/evalphobia/aws-sdk-go-v2-wrapper?status.svg
[102]: https://godoc.org/github.com/evalphobia/aws-sdk-go-v2-wrapper
[103]: https://img.shields.io/github/release/evalphobia/aws-sdk-go-v2-wrapper.svg
[104]: https://github.com/evalphobia/aws-sdk-go-v2-wrapper/releases/latest
[105]: https://img.shields.io/github/downloads/evalphobia/aws-sdk-go-v2-wrapper/total.svg?maxAge=1800
[106]: https://github.com/evalphobia/aws-sdk-go-v2-wrapper/releases
[107]: https://img.shields.io/github/stars/evalphobia/aws-sdk-go-v2-wrapper.svg
[108]: https://github.com/evalphobia/aws-sdk-go-v2-wrapper/stargazers


<!-- Testing -->

[201]: https://github.com/evalphobia/aws-sdk-go-v2-wrapper/workflows/test/badge.svg
[202]: https://github.com/evalphobia/aws-sdk-go-v2-wrapper/actions?query=workflow%3Atest
[203]: https://coveralls.io/repos/evalphobia/aws-sdk-go-v2-wrapper/badge.svg?branch=master&service=github
[204]: https://coveralls.io/github/evalphobia/aws-sdk-go-v2-wrapper?branch=master
[205]: https://codecov.io/gh/evalphobia/aws-sdk-go-v2-wrapper/branch/master/graph/badge.svg
[206]: https://codecov.io/gh/evalphobia/aws-sdk-go-v2-wrapper


<!-- Code Quality -->

[301]: https://goreportcard.com/badge/github.com/evalphobia/aws-sdk-go-v2-wrapper
[302]: https://goreportcard.com/report/github.com/evalphobia/aws-sdk-go-v2-wrapper
[303]: https://codeclimate.com/github/evalphobia/aws-sdk-go-v2-wrapper/badges/gpa.svg
[304]: https://codeclimate.com/github/evalphobia/aws-sdk-go-v2-wrapper
[305]: https://bettercodehub.com/edge/badge/evalphobia/aws-sdk-go-v2-wrapper?branch=master
[306]: https://bettercodehub.com/
[307]: https://www.codefactor.io/repository/github/evalphobia/aws-sdk-go-v2-wrapper/badge
[308]: https://www.codefactor.io/repository/github/evalphobia/aws-sdk-go-v2-wrapper
[309]: https://codebeat.co/badges/142f5ca7-da37-474f-9264-f708ade08b5c
[310]: https://codebeat.co/projects/github-com-evalphobia-aws-sdk-go-v2-wrapper-master
[311]: https://scrutinizer-ci.com/g/evalphobia/aws-sdk-go-v2-wrapper/badges/quality-score.png?b=master
[312]: https://scrutinizer-ci.com/g/evalphobia/aws-sdk-go-v2-wrapper/?branch=master

<!-- License -->
[401]: https://img.shields.io/badge/License-MIT-blue.svg
[402]: LICENSE.md
[403]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Faws-sdk-go-v2-wrapper.svg?type=shield
[404]: https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Faws-sdk-go-v2-wrapper?ref=badge_shield


Simple wrapper for [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)


## What's for?

The primary motivation is for my personal uses.
AWS SDK provides solid and rubust APIs. But sometimes it's very painful for me, too much parameters and many pointer values.
(I don't want to check `nil` every lines.)

I don't know it's useful for others.
It might helpful to just see these codes to know how to use [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) SDK besides official example test code.

(* Target version: [v0.24.0](https://github.com/aws/aws-sdk-go-v2/tree/v0.24.0) on 2020-07-21)


## APIs

aws-sdk-go-v2-wrapper provides three types of APIs.

| Type | Description | NamingRule |
|:--|:--|:--|
| Wrapper API | This type APIs wraps `Raw` API to avoid pointer values. | - |
| Raw API| This type APIs just call methods of [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2). You can use it when the APIs are not implemented in this repo (or Wrapper APIs are buggy :sweat_smile: ). | `Raw...` |
| X API | This type APIs provides easy way to call frequently used functions. | `X...` |

## Services

| Name |
|:--|
| [athena](/athena) |
| [cloudwatchlogs](/cloudwatchlogs) |
| [dynamodb](/dynamodb) |
| [kms](/kms) |
| [pinpointemail](/pinpointemail) |
| [s3](/s3) |
| [ssm](/ssm) |

## Environment variables

| Name | Description |
|:--|:--|
| `AWS_REGION` | Default AWS region for clients. |
| `AWS_ENDPOINT` | Default endpoints for clients. |
| `AWS_ATHENA_ENDPOINT` | Custom endpoint for Athena client. |
| `AWS_CLOUDWATCH_LOGS_ENDPOINT` | Custom endpoint for Cloudwatch Logs client. |
| `AWS_DYNAMODB_ENDPOINT` | Custom endpoint for DynamoDB client. |
| `AWS_KMS_ENDPOINT` | Custom endpoint for KMS client. |
| `AWS_PINPOINT_EMAIL_ENDPOINT` | Custom endpoint for PinpointEmail client. |
| `AWS_S3_ENDPOINT` | Custom endpoint for S3 client. |
| `AWS_SSM_ENDPOINT` | Custom endpoint for SSM client. |

Besides above, supported some values from [official env_config](https://github.com/aws/aws-sdk-go-v2/blob/master/aws/external/env_config.go)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Faws-sdk-go-v2-wrapper.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Faws-sdk-go-v2-wrapper?ref=badge_large)
