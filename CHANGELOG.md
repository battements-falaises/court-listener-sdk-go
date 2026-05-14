# Changelog

## 0.1.0 (2026-05-14)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/battements-falaises/court-listener-sdk-go/compare/v0.0.1...v0.1.0)

### Features

* **client:** optimize json encoder for internal types ([6bc3589](https://github.com/battements-falaises/court-listener-sdk-go/commit/6bc358932e1ff195febd5f4d4e716e875e4d2b18))
* **go:** add default http client with timeout ([babe381](https://github.com/battements-falaises/court-listener-sdk-go/commit/babe38153437298de2d0421a2ef60429bf08a306))
* **internal:** support comma format in multipart form encoding ([64d016a](https://github.com/battements-falaises/court-listener-sdk-go/commit/64d016a0a49303f1286e7bd99658d5c722e9a2cf))
* support setting headers via env ([76757c6](https://github.com/battements-falaises/court-listener-sdk-go/commit/76757c6aea2a6c01d725693230cd4a1313a69a65))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([d4abcfb](https://github.com/battements-falaises/court-listener-sdk-go/commit/d4abcfbb9583c5cfad8168e302fb07a47b0db956))
* prevent duplicate ? in query params ([d603611](https://github.com/battements-falaises/court-listener-sdk-go/commit/d603611ad434f9a17e48aadb4105ec8d388b63b4))


### Chores

* avoid embedding reflect.Type for dead code elimination ([6bed87d](https://github.com/battements-falaises/court-listener-sdk-go/commit/6bed87d9cf0f72c2dae28bc15ad036f05ea1c5f6))
* **ci:** skip lint on metadata-only changes ([b7dfb9d](https://github.com/battements-falaises/court-listener-sdk-go/commit/b7dfb9dba93c945dd5d9590e1505d9ac27b24ba3))
* **ci:** skip uploading artifacts on stainless-internal branches ([ad3b947](https://github.com/battements-falaises/court-listener-sdk-go/commit/ad3b9473f2885824844b3e9366a87b454c98ad10))
* **ci:** support opting out of skipping builds on metadata-only commits ([183d99b](https://github.com/battements-falaises/court-listener-sdk-go/commit/183d99b6a67885d6ef4b228b1b7ac9896fda3045))
* **client:** fix multipart serialisation of Default() fields ([52c9e94](https://github.com/battements-falaises/court-listener-sdk-go/commit/52c9e94f06118bac3cc13d05332b7bf8409fb20b))
* configure new SDK language ([9fb1ed2](https://github.com/battements-falaises/court-listener-sdk-go/commit/9fb1ed2c997d219412ebe5cb4739c4db3e83ebd6))
* configure new SDK language ([f21df1f](https://github.com/battements-falaises/court-listener-sdk-go/commit/f21df1f094b2c5720ed76896cb944b4a0f7bf92d))
* configure new SDK language ([4d211e1](https://github.com/battements-falaises/court-listener-sdk-go/commit/4d211e10f2fa3a86b2866a5babcb6957a87723a2))
* **internal:** codegen related update ([dc994e3](https://github.com/battements-falaises/court-listener-sdk-go/commit/dc994e328824f13bab43bb21befcd0524f8adcd8))
* **internal:** minor cleanup ([2bad12a](https://github.com/battements-falaises/court-listener-sdk-go/commit/2bad12a8a433b3bc87086efe609cf71c33957634))
* **internal:** more robust bootstrap script ([f043e78](https://github.com/battements-falaises/court-listener-sdk-go/commit/f043e78aa738671e2638378c9e7300c8335baa39))
* **internal:** support default value struct tag ([78dd3e9](https://github.com/battements-falaises/court-listener-sdk-go/commit/78dd3e9532ff7a30938b505afbf0a668c8aeb2a1))
* **internal:** tweak CI branches ([8025ae5](https://github.com/battements-falaises/court-listener-sdk-go/commit/8025ae57a3dd029db8732402d458c8a5d98125d5))
* **internal:** update gitignore ([ee6cc60](https://github.com/battements-falaises/court-listener-sdk-go/commit/ee6cc600d35e3caa9da1dbfadf7fd9b3c5831d83))
* **internal:** use explicit returns ([081202e](https://github.com/battements-falaises/court-listener-sdk-go/commit/081202e75495cbe8c90081a61c65c2c7c5f059fa))
* **internal:** use explicit returns in more places ([5939145](https://github.com/battements-falaises/court-listener-sdk-go/commit/593914594802af398770deaba34b044c15f62594))
* redact api-key headers in debug logs ([7f2a90f](https://github.com/battements-falaises/court-listener-sdk-go/commit/7f2a90ff9e2ae7066361de73dbfdc9298dd26e0e))
* remove unnecessary error check for url parsing ([112a5b0](https://github.com/battements-falaises/court-listener-sdk-go/commit/112a5b09d219c7dfd1edbef8ebb08a00a5572ecc))
* **test:** do not count install time for mock server timeout ([fe85cf8](https://github.com/battements-falaises/court-listener-sdk-go/commit/fe85cf83ebfcafcb9d61195a66022d80e5025e14))
* **tests:** bump steady to v0.19.4 ([2e2fe23](https://github.com/battements-falaises/court-listener-sdk-go/commit/2e2fe23c3ae21997f221ac8fbb3249066a1ad577))
* **tests:** bump steady to v0.19.5 ([376ea36](https://github.com/battements-falaises/court-listener-sdk-go/commit/376ea36873a8472090db2ae03b9cf39cea0cebcf))
* **tests:** bump steady to v0.19.6 ([392d4b1](https://github.com/battements-falaises/court-listener-sdk-go/commit/392d4b13fcd775d447635910446e380f11dea280))
* **tests:** bump steady to v0.19.7 ([758f5ab](https://github.com/battements-falaises/court-listener-sdk-go/commit/758f5abecb83d7d71b93c74c1064adac891a95c3))
* **tests:** bump steady to v0.20.1 ([f15f49b](https://github.com/battements-falaises/court-listener-sdk-go/commit/f15f49b35658ce64e33172bf1f8190d250bb00a3))
* **tests:** bump steady to v0.20.2 ([0e0abea](https://github.com/battements-falaises/court-listener-sdk-go/commit/0e0abea2d0cc38b5dd44b399b1f779b81b4a8bf2))
* **tests:** bump steady to v0.22.1 ([1621f65](https://github.com/battements-falaises/court-listener-sdk-go/commit/1621f654796444ddd100c90ec2ed9f0fcbb97ab0))
* **tests:** change mock server to steady ([2fdfba1](https://github.com/battements-falaises/court-listener-sdk-go/commit/2fdfba117ee8ef3490eba82095173388e784e7fb))
* update docs for api:"required" ([7db458e](https://github.com/battements-falaises/court-listener-sdk-go/commit/7db458eea8ab708ae3d0a402e9546c671dcedfec))
* update SDK settings ([0c0ca80](https://github.com/battements-falaises/court-listener-sdk-go/commit/0c0ca80f6cde8acb2d319fc9e215e92fa1e2687f))
