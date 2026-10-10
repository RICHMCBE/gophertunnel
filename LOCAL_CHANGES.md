# 로컬 변경

기준 버전: github.com/sandertv/gophertunnel v1.62.0. 원본 MIT LICENSE를 유지한다.

`minecraft/resource.Pack.WithDownloadURL`을 추가하여 이미 컴파일된 팩에 자체 HTTP 배포 주소를 연결한다. 원본 데이터·UUID·버전·체크섬·암호화 정보는 보존하며 원본 객체를 변경하지 않는다. 외부 URL을 시작 시 다시 다운로드하는 과정 없이 HTTP와 RakNet에 동일 데이터를 제공한다.

버전 업데이트 시 이 변경과 `internal/game/packcdn_protocol_test.go`를 함께 확인한다.
