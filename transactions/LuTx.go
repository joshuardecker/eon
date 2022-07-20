package transactions

type LuTx struct {
	inScripts []txScript

	outScripts []txScript
}

type txScript struct {
	scriptStr string
}
