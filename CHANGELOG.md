# Changelog

## 0.1.3 (2026-03-25)

Full Changelog: [v0.1.2...v0.1.3](https://github.com/greenflash-ai/cli/compare/v0.1.2...v0.1.3)

### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([57c8237](https://github.com/greenflash-ai/cli/commit/57c8237091aead50b55c99affc5395b509fac868))
* cli no longer hangs when stdin is attached to a pipe with empty input ([c7d8caa](https://github.com/greenflash-ai/cli/commit/c7d8caa92d667f30525fa91999167de53b80c72c))
* improve linking behavior when developing on a branch not in the Go SDK ([13b88fe](https://github.com/greenflash-ai/cli/commit/13b88fe84abcae8a4498fb711d383ac1698c573d))


### Chores

* **ci:** skip lint on metadata-only changes ([eb3c77f](https://github.com/greenflash-ai/cli/commit/eb3c77f9493976f139c7d57c7606a31d6889f740))
* **internal:** update gitignore ([8b5e563](https://github.com/greenflash-ai/cli/commit/8b5e563fab07db371ada13432beb03c6042eeeb9))
* **tests:** bump steady to v0.19.4 ([2dad7f4](https://github.com/greenflash-ai/cli/commit/2dad7f4deb7c6f1222aa30b25cc56c18d3e584e1))
* **tests:** bump steady to v0.19.5 ([4897bbb](https://github.com/greenflash-ai/cli/commit/4897bbbbb6936f6d9d8d652223158e04dfcfd5ec))
* **tests:** bump steady to v0.19.6 ([c40e4a6](https://github.com/greenflash-ai/cli/commit/c40e4a66f68454c5f1eeac63010d0bd983a3cfcc))
* **tests:** bump steady to v0.19.7 ([916eecf](https://github.com/greenflash-ai/cli/commit/916eecfe193632a6f2843aa0994b59b448038981))


### Refactors

* **tests:** switch from prism to steady ([6ac7aa6](https://github.com/greenflash-ai/cli/commit/6ac7aa6660aa258959f43419221e0c25ee7eff27))

## 0.1.2 (2026-03-17)

Full Changelog: [v0.1.1...v0.1.2](https://github.com/greenflash-ai/cli/compare/v0.1.1...v0.1.2)

### Bug Fixes

* better support passing client args in any position ([2b2b450](https://github.com/greenflash-ai/cli/commit/2b2b45031c6dfa0652a4d244a7d97b3e3510b740))
* improved workflow for developing on branches ([151049a](https://github.com/greenflash-ai/cli/commit/151049a019c5b98feee448398341bb7169fecd7d))
* no longer require an API key when building on production repos ([47940a0](https://github.com/greenflash-ai/cli/commit/47940a030b91d40f1bc3d55f4cbd2e9b488917e4))


### Chores

* **internal:** tweak CI branches ([46c12f5](https://github.com/greenflash-ai/cli/commit/46c12f54efd11b4ab0a6c7c71806b0b7c6a7b1b9))

## 0.1.1 (2026-03-14)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/greenflash-ai/cli/compare/v0.1.0...v0.1.1)

### Bug Fixes

* only set client options when the corresponding CLI flag or env var is explicitly set ([fb3ea7d](https://github.com/greenflash-ai/cli/commit/fb3ea7d9ecc61fbf5c5f2f8eca1094e8c4c07dff))


### Chores

* **internal:** codegen related update ([dc63e7b](https://github.com/greenflash-ai/cli/commit/dc63e7bf61a160f164e2f9d25abe14b9122d05b7))

## 0.1.0 (2026-03-12)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/greenflash-ai/cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([280cc7b](https://github.com/greenflash-ai/cli/commit/280cc7ba3784071794637635979c550315adadc5))


### Chores

* configure new SDK language ([d1bf6a7](https://github.com/greenflash-ai/cli/commit/d1bf6a71d68cefaceadd15acd25e8422c2cbf9e9))
* update SDK settings ([90ccdb2](https://github.com/greenflash-ai/cli/commit/90ccdb2e643ec6348fe707c2581d71fe6aef64dd))
