package pwned

//Template for showing the list of breaches to the user
const breachesTmpl = `

Hello!

The email account has been compromised in a security breach related to the following Companies:


Affected Companies:

	Name
{{range .}}
	-{{.Name}}  Breach Date -> {{.BreachDate}}
{{end}}

TIP: You can get detail information using: gopwned company <name>

`

//Template for showing the information about a company's security breach
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
