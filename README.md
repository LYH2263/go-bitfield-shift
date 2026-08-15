# go-bitfield-shift

位移越界

internal/bitfield/bits.go 用 1<<bit 未按 word 取模

```bash
go test ./... -count=1
```
