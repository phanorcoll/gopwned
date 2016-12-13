package pwned

const breachesTmpl = `

Hello!

The email account has been compromised in a security breach related to the following Companies:


Affected Companies:

	Name
{{range .}}
	-{{.Name}} [{{.BreachDate}}]
{{end}}

TIP: You can get detail information using: gopwned company <name>

`
const breachTmpl = `

Information about the breach
-----------------------------
{{.BreachDate}}
{{.Name}}
{{.Title}}
{{.Domain}}
{{.AddedDate}}
{{.Description}}

`
