package script

type ScriptSig struct {
	Signature  []byte
	PubkeyHash []byte
}

type ScriptPubKey struct {
	PubkeyHash []byte
}

func P2PKH(scriptSig ScriptSig, scriptPubKey ScriptPubKey) {

}
