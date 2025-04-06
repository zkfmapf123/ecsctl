# ecsctl

## Description & Use

- ecs CLI Tool ( Like eksctl ...)

## Install

```sh
    brew tap zkfmapf123/homebrew-tap
    brew search ecsctl

    brew install zkfmapf123/tap/ecsctl

```

## Diff kubernetes Object

| eksctl        | ecsctl        | command       | implementation |
| ------------- | ------------- | ------------- | -------------- |
| api-resources | api-resources | api-resources | O              |
| namespace     | cluster       | cl            | X              |
| pods          | services      | s             | X              |
| containers    | containers    | c             | X              |
| svc           | alb           | al            | X              |
| exec          | exec          | exec          | O              |

## implementation

- ecsns (configuration profile, cluster, svc)

## Update History

| Todo                        | UpdatedAt | Version |
| --------------------------- | --------- | ------- |
| Add command : api-resources | 2024.9.20 | none    |
| Add command : get           | 2024.9.20 | none    |
| Add command : exec          | 2024.9.21 | none    |
