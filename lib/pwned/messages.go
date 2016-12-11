package pwned

const breachTmpl = `

Hello!

The email account has been compromised in a security breach related to the following domain(s):


Affected websites:

{{range .}}
	- {{.Domain}}
{{end}}

TIP: You can get detail information using: gopwned verify user@example.com --domain <domain-name>

`
