FROM 120.77.97.206:8077/library/go-zero-alpine:v1

WORKDIR /app

COPY etc /app/etc
COPY {{.ExeFile}}-{{.BaseImage}} /app/{{.ExeFile}}

{{if .HasPort}}
EXPOSE {{.Port}}
{{end}}

CMD ["./{{.ExeFile}}"{{.Argument}}]
