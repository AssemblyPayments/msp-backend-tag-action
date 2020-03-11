
## version-json-tagging-action

## Usage

Create a file called something like `tag-master-on-merge.yml` with the following content under your repo's `github/workflows`. And it'll it will create an annotated tag with name from `version.json` when PR is merged to `master`

```
name: Tag master on merge
on:
  push:
    branches:
      - master

jobs:
  build:
    runs-on: ubuntu-latest
    steps:

    - name: Checkout
      uses: actions/checkout@v2
      with:
        fetch-depth: '0'

    - name: Tag master
      uses: mx51/version-json-tagging-action@master
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```

More [info](https://help.github.com/en/actions/building-actions) on building github actions
