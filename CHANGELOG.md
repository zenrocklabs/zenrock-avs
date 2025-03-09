# [1.6.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.5.0...v1.6.0) (2025-03-09)


### Features

* quorum creator + ejector setter utils ([1324a95](https://github.com/zenrocklabs/zenrock-avs/commit/1324a950b5c2274ce116f9d2c63a38cc1b093d91))

# [1.5.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.18...v1.5.0) (2025-03-07)


### Features

* add strategy remover ([7c7315e](https://github.com/zenrocklabs/zenrock-avs/commit/7c7315eb2edc6f667ae908f45adfcc1f22f099e8))
* gas estimation for strategy adder ([24e2ff2](https://github.com/zenrocklabs/zenrock-avs/commit/24e2ff2e37ad4cd9195268cd683cb2d6a333c624))

## [1.4.18](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.17...v1.4.18) (2025-02-06)


### Bug Fixes

* pass sidecar oracle to operator ([aeacdd5](https://github.com/zenrocklabs/zenrock-avs/commit/aeacdd5ade2881f710431aaff9541549713eddb9))

## [1.4.17](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.16...v1.4.17) (2025-01-28)


### Bug Fixes

* remove duplicate config field ([ac9e7b7](https://github.com/zenrocklabs/zenrock-avs/commit/ac9e7b7506cb7fce8948ab35ec39b504493b8e3e))

## [1.4.16](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.15...v1.4.16) (2025-01-23)


### Bug Fixes

* implement deregister CLI cmd ([181a708](https://github.com/zenrocklabs/zenrock-avs/commit/181a70896ef5d3dec1872b9bc9036ccfa8ca1519))

## [1.4.15](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.14...v1.4.15) (2025-01-23)


### Bug Fixes

* registration CLI ([fe2f555](https://github.com/zenrocklabs/zenrock-avs/commit/fe2f555d2f8de0467984f497406bed33e206dcf4))

## [1.4.14](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.13...v1.4.14) (2025-01-06)


### Bug Fixes

* move registration code to CLI, don't require EL ECDSA key unless registering ([#7](https://github.com/zenrocklabs/zenrock-avs/issues/7)) ([ec71533](https://github.com/zenrocklabs/zenrock-avs/commit/ec715332f6cf53980bcc2a7b04d80f4dbcfec504))

## [1.4.13](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.12...v1.4.13) (2024-12-24)


### Bug Fixes

* bump `eigensdk-go` ver +`zrchain` ver ([d804ddf](https://github.com/zenrocklabs/zenrock-avs/commit/d804ddfa49365368ba0df25fb209f31d31d424f9))

## [1.4.12](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.11...v1.4.12) (2024-12-17)


### Bug Fixes

* send delegated validators not in active set for removal via tasks ([#6](https://github.com/zenrocklabs/zenrock-avs/issues/6)) ([bee5d1e](https://github.com/zenrocklabs/zenrock-avs/commit/bee5d1e480f38958975e11e256c6952fb10b12aa))

## [1.4.11](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.10...v1.4.11) (2024-11-22)


### Bug Fixes

* add proxy for TaskManager ([#5](https://github.com/zenrocklabs/zenrock-avs/issues/5)) ([7f6608c](https://github.com/zenrocklabs/zenrock-avs/commit/7f6608cd8cf212607a1a5278c4fbca40f984dd20))

## [1.4.10](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.9...v1.4.10) (2024-11-22)


### Bug Fixes

* add response delay to aggregator ([21ab304](https://github.com/zenrocklabs/zenrock-avs/commit/21ab304cdce4ed76d2695b0801588b9ede200205))

## [1.4.9](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.8...v1.4.9) (2024-11-21)


### Bug Fixes

* set task cadence to 15 mins ([15fa163](https://github.com/zenrocklabs/zenrock-avs/commit/15fa16327653199e3cdf7026abc6b11bd3ab52d7))

## [1.4.8](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.7...v1.4.8) (2024-11-21)


### Bug Fixes

* increased cadence for tasks (testnet) ([069984b](https://github.com/zenrocklabs/zenrock-avs/commit/069984bb05804299f14c3faef6f2803e3435f9dd))

## [1.4.7](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.6...v1.4.7) (2024-11-21)


### Bug Fixes

* revert taskmanager changes + aggregator instantly sends task ([2fb2d72](https://github.com/zenrocklabs/zenrock-avs/commit/2fb2d72127c5ac2946020077737725d97fa96336))

## [1.4.6](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.5...v1.4.6) (2024-11-21)


### Bug Fixes

* refactor contracts ([#4](https://github.com/zenrocklabs/zenrock-avs/issues/4)) ([b9cd0f6](https://github.com/zenrocklabs/zenrock-avs/commit/b9cd0f6215d5e8ec913833222af543b2312a2f97))

## [1.4.5](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.4...v1.4.5) (2024-11-20)


### Bug Fixes

* update deployment addrs + strategy-adder ([ed492d3](https://github.com/zenrocklabs/zenrock-avs/commit/ed492d333a87323477b0674772a8f20484d05f03))

## [1.4.4](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.3...v1.4.4) (2024-11-20)


### Bug Fixes

* remove stake check on registration ([3e24c62](https://github.com/zenrocklabs/zenrock-avs/commit/3e24c62315f3f7c70e62528368a0d1eea57ce244))

## [1.4.3](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.2...v1.4.3) (2024-11-20)


### Bug Fixes

* use forked `eigensdk-go` ([d8b3622](https://github.com/zenrocklabs/zenrock-avs/commit/d8b36220c40a0e5bea325139b1b40c67784d3b0a))

## [1.4.2](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.1...v1.4.2) (2024-11-20)


### Bug Fixes

* bump `zrchain` ver ([58c815b](https://github.com/zenrocklabs/zenrock-avs/commit/58c815babf102b9770aadc617be92fb67c37324a))
* uncomment inactive set query tasks ([57d0c85](https://github.com/zenrocklabs/zenrock-avs/commit/57d0c85c4b05628de1f53ed0f6e4476fde0bda60))

## [1.4.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.4.0...v1.4.1) (2024-11-20)


### Bug Fixes

* temporarily comment out inactive set query code ([5234ad6](https://github.com/zenrocklabs/zenrock-avs/commit/5234ad6bb671b32c96917a52c1b9b2156403ac91))

# [1.4.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.6...v1.4.0) (2024-11-20)


### Features

* update contracts for mainnet ([aa3ee1c](https://github.com/zenrocklabs/zenrock-avs/commit/aa3ee1c02d1ab6a1af7c63be37b84e467ca9a451))

## [1.3.6](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.5...v1.3.6) (2024-11-18)


### Bug Fixes

* generate random salt for registration ([f9c407b](https://github.com/zenrocklabs/zenrock-avs/commit/f9c407b0f0e7adb988881b3cf03a2a4db4eb7917))

## [1.3.5](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.4...v1.3.5) (2024-11-15)


### Bug Fixes

* make zrchain addr configurable ([077b0ce](https://github.com/zenrocklabs/zenrock-avs/commit/077b0ce3a13e51bc407320d937cbc3dd4a203a9b))

## [1.3.4](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.3...v1.3.4) (2024-11-12)


### Bug Fixes

* set registration error logs to debug level ([ecc9380](https://github.com/zenrocklabs/zenrock-avs/commit/ecc938037091080d179560e4e0eb0271020a76d0))

## [1.3.3](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.2...v1.3.3) (2024-11-11)


### Bug Fixes

* aggregator log text ([374b8e8](https://github.com/zenrocklabs/zenrock-avs/commit/374b8e83f5a0a13cf45098f136e8d43d77a6803e))

## [1.3.2](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.1...v1.3.2) (2024-11-11)


### Bug Fixes

* increase task cadence to 15min ([09ce915](https://github.com/zenrocklabs/zenrock-avs/commit/09ce915d1707cc19eeebcc4e458ff92b9b65dd4f))

## [1.3.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.0...v1.3.1) (2024-11-11)


### Bug Fixes

* aggregator + registration flow ([7942310](https://github.com/zenrocklabs/zenrock-avs/commit/7942310301a9e3a62a4a2c4009b119369eecb3fe))

# [1.3.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.0...v1.3.0) (2024-11-07)


### Bug Fixes

* `registration.go` ([5d2efa5](https://github.com/zenrocklabs/zenrock-avs/commit/5d2efa5ecf4cac9617165fb59cfe6bf5e49b266b))
* remove ERC20Mock-related code ([423e2d0](https://github.com/zenrocklabs/zenrock-avs/commit/423e2d0163e7778925d456f648535f23ba5e8a0c))
* remove ERC20Mock-related code ([bf1aa0a](https://github.com/zenrocklabs/zenrock-avs/commit/bf1aa0ab20767898801ffd047e61c7818bf32ee2))


### Features

* force new tag version ([7e7b416](https://github.com/zenrocklabs/zenrock-avs/commit/7e7b41621a1dfd19fdbfdfdc420e356d2fc9cfa2))

## [1.2.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.0...v1.2.1) (2024-11-07)


### Bug Fixes

* `registration.go` ([5d2efa5](https://github.com/zenrocklabs/zenrock-avs/commit/5d2efa5ecf4cac9617165fb59cfe6bf5e49b266b))
* remove ERC20Mock-related code ([423e2d0](https://github.com/zenrocklabs/zenrock-avs/commit/423e2d0163e7778925d456f648535f23ba5e8a0c))
* remove ERC20Mock-related code ([bf1aa0a](https://github.com/zenrocklabs/zenrock-avs/commit/bf1aa0ab20767898801ffd047e61c7818bf32ee2))

## [1.2.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.0...v1.2.1) (2024-11-06)


### Bug Fixes

* remove ERC20Mock-related code ([423e2d0](https://github.com/zenrocklabs/zenrock-avs/commit/423e2d0163e7778925d456f648535f23ba5e8a0c))
* remove ERC20Mock-related code ([bf1aa0a](https://github.com/zenrocklabs/zenrock-avs/commit/bf1aa0ab20767898801ffd047e61c7818bf32ee2))

## [1.2.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.0...v1.2.1) (2024-11-06)


### Bug Fixes

* Dockerfile ([073bbb8](https://github.com/zenrocklabs/zenrock-avs/commit/073bbb8f5ca9eaeee5e5c6825c14ece0fe0e0aa1))

# [1.2.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.4...v1.2.0) (2024-11-06)


### Bug Fixes

* new task manager contract for testnet deployment ([05f9ea3](https://github.com/zenrocklabs/zenrock-avs/commit/05f9ea3e5d713cc1167d2c0da56d2a7e98f0c743))
* reduce size of task manager contract ([f70aa5c](https://github.com/zenrocklabs/zenrock-avs/commit/f70aa5c7953bafad6f9ea1b2b68840a0d13d9632))
* task manager ([d112e1b](https://github.com/zenrocklabs/zenrock-avs/commit/d112e1bb2379189c36e89f74a99eb8ff970bb1b7))


### Features

* improve task manager contract ([fcac103](https://github.com/zenrocklabs/zenrock-avs/commit/fcac103d323ac389dec4edc47e77ec60d69f1b93))

## [1.1.4](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.3...v1.1.4) (2024-11-01)


### Bug Fixes

* use `go-client` from `zrchain/v5` ([eaa4e33](https://github.com/zenrocklabs/zenrock-avs/commit/eaa4e33a2a327a0010a90f47c41cf6df775dfc1c))

## [1.1.3](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.2...v1.1.3) (2024-11-01)


### Bug Fixes

* temporarily comment out go-client call code ([dcdd361](https://github.com/zenrocklabs/zenrock-avs/commit/dcdd36143489dfe71021f885eb7a065e6fa9e584))

## [1.1.2](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.1...v1.1.2) (2024-11-01)


### Bug Fixes

* bump `go-client` version ([2f43dc3](https://github.com/zenrocklabs/zenrock-avs/commit/2f43dc381e2207bc85e0d9b87dc0ef19278e1cfc))

## [1.1.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.0...v1.1.1) (2024-11-01)


### Bug Fixes

* uncomment code `registration.go` + fix `go.mod` ([795c0b3](https://github.com/zenrocklabs/zenrock-avs/commit/795c0b301b143f499042204fe26d580b03a7844b))

# [1.1.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.0.1...v1.1.0) (2024-10-31)


### Features

* bridge active set via tasks ([72a88a5](https://github.com/zenrocklabs/zenrock-avs/commit/72a88a50d8e7c14e0730e532666cd728fd6a65d0))

## [1.0.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.0.0...v1.0.1) (2024-10-21)


### Bug Fixes

* temporarily disable code `registration.go` ([45f883f](https://github.com/zenrocklabs/zenrock-avs/commit/45f883f21140b1360b4b13e14a3bc9f06e0f2aaf))

# 1.0.0 (2024-10-21)


### Bug Fixes

* semantic release branch ([28a6331](https://github.com/zenrocklabs/zenrock-avs/commit/28a633113e3f5bb59216ccbfa41cf92439ed5e18))


### Features

* initial version ([fe0e597](https://github.com/zenrocklabs/zenrock-avs/commit/fe0e597dea31de419937044078b209684deab3e5))


### Reverts

* Revert "revert to eigenlayer-middleware testnet-holesky branch b/c dev has payments stuff" ([abe3340](https://github.com/zenrocklabs/zenrock-avs/commit/abe3340e64ccc94b3e8959608cc9d687b3c86307))

# [1.4.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.5...v1.4.0) (2024-10-16)


### Features

* overhauled contract, avs now fully working end-to-end ([22c2a97](https://github.com/zenrocklabs/zenrock-avs/commit/22c2a97e4026327d6a69dd3ec91eed5e6d06ee07))
* strategy adder ([6fcd496](https://github.com/zenrocklabs/zenrock-avs/commit/6fcd49602856c4ef31588c89e8fef316eb331b8c))
* ZRServiceManager overhaul, strategy adder, AVS working e2e ([7cf1dc7](https://github.com/zenrocklabs/zenrock-avs/commit/7cf1dc77242d9ff4fe129618f6017d6436be8e60))

## [1.3.5](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.4...v1.3.5) (2024-09-06)


### Bug Fixes

* comment out log message ([58cb03f](https://github.com/zenrocklabs/zenrock-avs/commit/58cb03fe4b910f00fa264dc1ac600b88cd5677ae))

## [1.3.4](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.3...v1.3.4) (2024-09-06)


### Bug Fixes

* task cadence to 1000 seconds ([b59eb2d](https://github.com/zenrocklabs/zenrock-avs/commit/b59eb2de1d8684f20f9e4517b8c4b20960b59480))

## [1.3.3](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.2...v1.3.3) (2024-09-06)

## [1.3.2](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.1...v1.3.2) (2024-09-06)


### Bug Fixes

* strategy deposit cli + temp disable faulty delegate call ([08a3015](https://github.com/zenrocklabs/zenrock-avs/commit/08a3015cc35d33e51160c5b9cf8c3f024c9c4881))

## [1.3.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.3.0...v1.3.1) (2024-09-04)

# [1.3.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.1...v1.3.0) (2024-09-03)


### Features

* update deployer script to use WETH strategy + update configs ([c35098e](https://github.com/zenrocklabs/zenrock-avs/commit/c35098e6a12101aa24a70e0e6aea9c75a6c48986))

## [1.2.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.2.0...v1.2.1) (2024-09-03)

# [1.2.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.1.0...v1.2.0) (2024-08-30)


### Features

* create AVS delegation during `registerOperatorOnStartup` ([0a3bbd6](https://github.com/zenrocklabs/zenrock-avs/commit/0a3bbd6f7834d84176df6747a04d3fea7e5cd5d1))

# [1.1.0](https://github.com/zenrocklabs/zenrock-avs/compare/v1.0.2...v1.1.0) (2024-08-29)


### Features

* automatically register operator if not enrolled ([c4ef24a](https://github.com/zenrocklabs/zenrock-avs/commit/c4ef24a676ca7b0bbb0a406ba4ac7dff72dd0738))

## [1.0.2](https://github.com/zenrocklabs/zenrock-avs/compare/v1.0.1...v1.0.2) (2024-08-23)


### Bug Fixes

* update config + misc fixes ([fe598ff](https://github.com/zenrocklabs/zenrock-avs/commit/fe598ff2734101775a354253806fef574ccd9fd1))

## [1.0.1](https://github.com/zenrocklabs/zenrock-avs/compare/v1.0.0...v1.0.1) (2024-08-22)


### Bug Fixes

* update field naming in EL deployment json ([5fccc75](https://github.com/zenrocklabs/zenrock-avs/commit/5fccc75ad1d6d599cd1421f57b2142a093bb4a60))
* use holesky deployment json from `eigenlayer-contracts` repo ([bf13d70](https://github.com/zenrocklabs/zenrock-avs/commit/bf13d704629ae4de2eb3eaa1612bb9c9db1451b8))

# 1.0.0 (2024-08-22)


### Bug Fixes

* CI pipelines ([98e92ea](https://github.com/zenrocklabs/zenrock-avs/commit/98e92ea8ed24da0a402ae3c598b0199b95e57011))


### Features

* add CI pipelines ([ceac379](https://github.com/zenrocklabs/zenrock-avs/commit/ceac379c20694fb341cdc162cd072cbad0ec16eb))
* zenrock impl ([d87cf85](https://github.com/zenrocklabs/zenrock-avs/commit/d87cf85ca56bad83026118e60e0e5c473c4d9151))


### Reverts

* Revert "revert to eigenlayer-middleware testnet-holesky branch b/c dev has payments stuff" ([abe3340](https://github.com/zenrocklabs/zenrock-avs/commit/abe3340e64ccc94b3e8959608cc9d687b3c86307))
