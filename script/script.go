package script

type ScriptSig struct {
	signature  []byte
	PubkeyHash []byte
}

type ScriptPubKey struct {
	PubkeyHash []byte
}

func P2PKH(scriptSig ScriptSig, scriptPubKey ScriptPubKey) {

}
