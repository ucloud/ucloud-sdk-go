## 0.21.21 (2021-10-27)

ENHANCEMENTS:

- Update all APIs of `PathX` (#299)

## 0.21.20 (2021-10-27)

ENHANCEMENTS:

- Update all APIs of `PathX` (#296)

## 0.21.19 (2021-10-25)

ENHANCEMENTS:

- Update all APIs of `UFS` (#294 )

## 0.21.18 (2021-09-06)

ENHANCEMENTS:

- Update all APIs of `ISMS` (#292)

## 0.21.17 (2021-09-03)

ENHANCEMENTS:

- Update all APIs of `Cube` (#290)

## 0.21.16 (2021-08-31)

ENHANCEMENTS:

- Update all APIs of `UCDN` (#288)
- Update all APIs of `ISMS` (#287)

## 0.21.15 (2021-08-30)

ENHANCEMENTS:

- Update all APIs of `USMS` (#284)

## 0.21.14 (2021-08-19)

ENHANCEMENTS:

- Add `UPHost` support for the metadata server

## 0.21.13 (2021-08-17)

ENHANCEMENTS:

- Update all APIs of `Cube` (#280)

## 0.21.12 (2021-08-12)

ENHANCEMENTS:

- Update all APIs of `UMem` (#277)
- Update all APIs of `PathX` (#276)

## 0.21.11 (2021-07-27)

ENHANCEMENTS:

- Update all APIs of `UDTS` (#274)

## 0.21.10 (2021-07-22)

ENHANCEMENTS:

- Update all APIs of `UHost` (#272)
- Update all APIs of `USMS` (#271)

## 0.21.9 (2021-07-08)

ENHANCEMENTS:

- Update all APIs of `UCDN` (#268)
- Update all APIs of `UK8S` (#266)

## 0.21.8 (2021-07-05)

ENHANCEMENTS:

- Update all APIs of `UPHost` (#263)

## 0.21.7 (2021-07-02)

ENHANCEMENTS:

- Update all APIs of `UMem` (#261)
- Update all APIs of `UCDN` (#260)

## 0.21.6 (2021-06-11)

ENHANCEMENTS:

- Update all APIs of `VPC2.0` (#258)

## 0.21.5 (2021-06-03)

BUG FIXES:

- delete duplicated field `Action` of the object `CreateCubePodResponse` about api `CreateCubePod`.[#256]
- delete duplicated field `Action` of the object `GetCubePriceResponse` about api `CreateCubePod`.[#256]

## 0.21.4 (2021-05-14)

ENHANCEMENTS:

- Update all APIs of `UAccount` (#254)
- Update all APIs of `Cube` (#253)
- Update all APIs of `UK8S` (#252)

## 0.21.3 (2021-05-08)

ENHANCEMENTS:

- Update all APIs of `UNet` (#250)
- Update all APIs of `UFile`(US3) (#249)
- Update all APIs of `UMem` (#248)
- Update all APIs of `ULB` (#247 )
- Update all APIs of `UDB` (#246)

## 0.21.2 (2021-04-30)

ENHANCEMENTS:

- Update all APIs of `UDPN` (#244)
- Update all APIs of `UEC` (#241)
- Update all APIs of `UBill` (#240)
- Update all APIs of `VPC2.0` (#239)
- Update all APIs of `USMS` (#238)
- Update all APIs of `IPSecVPN` (#237)
- Update all APIs of `UNet` (#236)

## 0.21.1 (2021-04-28)

ENHANCEMENTS:

- Update all APIs of `UFS` (#234)
- Update all APIs of `PathX` (#232)
- Update all APIs of `UDDB` (#231)
- Update all APIs of `UCDN` (#230)
- Update all APIs of `UHost` (#229)
- Update all APIs of `UDisk` (#227)

## 0.21.0 (2021-04-22)

FEATURES:

- add `UGN` apis to be consistent with official document(#225  )
- add `URocketMQ` apis to be consistent with official document(#224  )
- add `ISMS ` apis to be consistent with official document(#223   )

ENHANCEMENTS:

- add u-timestamp-ms request header for tracing e2e latency

## 0.20.2 (2021-04-14)

## 0.20.1 (2021-03-11)

ENHANCEMENTS:

- Update all APIs of `PathX` (#217 )

## 0.20.0 (2021-03-05)

FEATURES:

- Add request encoder to support **json-only** API (#214)

ENHANCEMENTS:

- Add `UpdateBackendBatch` to private package (#214)

## 0.19.2 (2021-03-04)

ENHANCEMENTS:

- Update all APIs of `UCDN` (#212)

## 0.19.1 (2021-02-23)

ENHANCEMENTS:

- Update all APIs of `Cube` (#209)

## 0.19.0 (2021-02-03)

FEATURES:

- Add product `UEC`

## 0.18.1 (2021-01-27)

ENHANCEMENTS:

- Update all APIs of `UHost` (#204 )

## 0.18.0 (2020-11-26)

FEATURES:

- Add product `UK8S`
- Add product `Cube`

## 0.17.5 (2020-09-25)

ENHANCEMENTS:

- Update all APIs of `UMem` (#195)
- add client error about `NullCredentialError` and `NullConfigError`

## 0.17.4 (2020-09-18)

BUG FIXES:

- fix the type of FirewallSet from object to array about ULB response model ULBSimpleSet (#192  )
- fix the type of FirewallSet from object to array about ULB response model ULBSet (#192  )

## 0.17.3 (2020-09-03)

BUG FIXES:

- fix the type of `EnableLog` from `bool` to `int` about ULB response model `ULBSimpleSet ` (#190 )

## 0.17.2 (2020-09-03)

ENHANCEMENTS:

- Update all APIs of `ULB` (#187 )

BUG FIXES:

- Delete parameter `EnableHTTP2 ` from `ULB` apis: `CreateVServer` and `DescribeVServer`(#187 )

## 0.17.1 (2020-08-20)

ENHANCEMENTS:

- Add request field `RdmaClusterId ` of  `CreateUDisk` api about UDIsk product(#185)

## 0.17.0 (2020-08-13)

FEATURES:

- Add `UCDN` apis to be consistent with [official document](https://docs.ucloud.cn/api/ucdn-api/README) (#177)

ENHANCEMENTS:

- Update all APIs of `UDisk` (#181)
- Update all APIs of `VPC` (#180)
- Update all APIs of `UHost` (#179)
- Update all APIs of `ULB` (#178)

## 0.16.6 (2020-07-30)

ENHANCEMENTS:

- Update all APIs of `UHost` (#175)
- Update all APIs of `UDisk` (#174)

## 0.16.5 (2020-07-24)

ENHANCEMENTS:

- update response logging about request uuid (#172)
- Add `ErrResponseBodyError` and `ErrEmptyResponseBodyError` as `ServerError`  (#171)
- Update external `LoadSTSConfig` about `Region` and `Zone` to be compatible with `UPHost` and another. (#171)
- refine the public and private keys example order

## 0.16.4 (2020-07-08)

ENHANCEMENTS:

- Add request uuid into response logging (#166)
- Add `CreateAttachUDisk` api (#165)
- Add `CreateNIC` api (#164)
- Add `DescribeNIC` api (#164)
- Add `AttachNIC` api (#164)
- Add `DetachNIC` api (#164)
- Add `DeleteNIC` api (#164)
- Update `DescribeVMInstance` api (#164)
- Update `CreateVMInstance` api (#164)
- Update `DescribeEIP` api (#164)
- Update `AllocateEIP` api (#164)
- Update `DescribeDisk` api (#164)

## 0.16.3 (2020-06-12)

ENHANCEMENTS:

- update external `LoadSTSConfig` to set default `Region`, `Zone`, `BaseUrl` for `ConfigProvider`

## 0.16.2 (2020-06-05)

ENHANCEMENTS:

- Add `ModifyUHostIP ` to `UHost` api
- Add `CreateCertificate` api (#156)
- Add `DescribeCertificate` api (#156)
- Add `DeleteCertificate` api (#156)
- Add `DescribeOPLogs` api (#156)
- Update `DescribeVMInstance` api (#156)
- Update `CreateVS` api (#156)
- Update `UHost` apis to be consistent with official document

## 0.16.1 (2020-05-21)

ENHANCEMENTS:

- Update `UCloudStack` apis to be consistent with official document
- Deprecated the api about `LoginByPassword`

## 0.16.0 (2020-04-23)

FEATURES:

- add UMedia apis to be consistent with [official document](https://docs.ucloud.cn/api/umedia-api/README)(#147  )

## 0.15.0 (2020-04-09)

FEATURES:

- add `UDTS` apis  to be consistent with official document(#142 )

## 0.14.3 (2020-03-30)

ENHANCEMENTS:

- Update `ULB` apis to be consistent with [official document](https://docs.ucloud.cn/api/ulb-api/README)(#137 )

## 0.14.2 (2020-03-30)

BUG FIXES:

- Fix the model `SubnetInfo` to `VPCSubnetInfoSet` of  `DescribeSubnetResponse.DataSet` (#135)
- Fix the model `SubnetResource ` to `ResourceInfo` of  `DescribeSubnetResourceResponse.DataSet` (#135)

## 0.14.1 (2020-03-26)

ENHANCEMENTS:

- Update `VPC` apis to be consistent with official document

## 0.14.0 (2020-03-05)

FEATURES:

- add the method about Generic Invoke of the Client
- add utest framework

ENHANCEMENTS:

- Update `UHost` apis to be consistent with official document

## 0.13.2 (2019-12-27)

ENHANCEMENTS:

- Update `ResizeAttachedDisk`

## 0.13.1 (2019-12-26)

ENHANCEMENTS:

- update `ucloudstack` apis

## 0.13.0 (2019-11-22)

ENHANCEMENTS:

- Add ucloud stack api about `vm/disk/eip/user` (#101)
