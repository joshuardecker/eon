package transactions

// This struct are the tx's on the Luncheon Network.
type LuTx struct {
	inScripts []scriptStr

	outScripts []scriptStr
}

// This struct is the transaction script or scripts
type scriptStr struct {
	ScriptStr string
}

// This function adds a scriptStr to the tx scriptStrs.
// First input is the scriptStr thats being added.
// The second is a bool.
// Enter true to add the scriptStr to the inScripts on the tx.
// Enter false to add the scriptStr to the outScripts on the tx.
// Returns nothing.
func (l LuTx) AddScriptStr(scriptstr string, scriptType bool) {

	tScript := new(scriptStr)
	tScript.ScriptStr = scriptstr

	if scriptType {

		l.inScripts = append(l.inScripts, *tScript)
	}

	if !scriptType {

		l.outScripts = append(l.outScripts, *tScript)
	}
}
