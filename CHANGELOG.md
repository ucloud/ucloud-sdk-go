## 0.16.5 (2020-07-24)

ENHANCEMENTS:

- update response logging about request uuid (#72)
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

