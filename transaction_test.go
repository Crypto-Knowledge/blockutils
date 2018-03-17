package blockutils

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

var digibytetxcoinbase = "010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff1803527a5f04ea66ab5a08540000fd29000000052f6d70682f00000000020000000000000000266a24aa21a9ed735a4c6d92c7bc860c0558bf0b49feb40e553dffe846613bd6d6bac983473d2cf934eb8b120000001976a914510fffca0668d410aea742e95a2fefa7952f695e88ac0120000000000000000000000000000000000000000000000000000000000000000000000000"
var digibytetx = "0100000002be92100bc9f1b6e6e11637d3bbc841bea9cffcc0a5d710ef83e36c438d5dcd78020000006b483045022100eb4671f9bbcbcc937855ef8aad774ff81cd4aedc65f79fedf2a9c88c9cd566c6022034039dd992ab0be0db95a1d7b615bb2c39e7b16515c74c9f021dd39ac0ffe213012102f24f8135e2f62f81d6c4ff172fd2681a3e03cf7485510a2871ca2c41b5aa9733ffffffff89491ae9534c2c5b7f000352588ff7778999b5ee0d19cad1bc0396e3fdf48c9c000000006a47304402200d5fff4b02e1b89e7a5067c6f8383d08b56c14ea54b6cb6257601dd150b11a07022025c75524788cc1de76146d2849f4cb821739f47459ed8fc8057cf9b703f399f2012102629fe53bdbf029c7d3be5dd64758229f0f754529981d70788d916e48c9e9af6cffffffff025a4ccf805e0000001976a914b788297cf734149f6225228c50ff905917aa8f4088ac51e68b050a0000001976a914d00455c4000530f93bf53e32615a7dee6da2a03b88ac4d7a5f00"

var ltcsegwittx = "02000000000101b539b9e41717be24d14c06cd72aed10a1d9593a860067850116e458d96b56d660000000017160014336d166ab51b21b3ef2f0c885b7004bd3ad38b3dfeffffff0200c2eb0b000000001976a914f6a3510afba93284b4a1969bcf411a225423acd188ac4924fe020000000017a9148a4275e9d10794c5d54d0b2ef9d33cb028258c5a870247304402202a91f2110e7a06b926bb8166fbffac12552326c6099ff1f077f2f8e9a5ac74be02202d19aad053f65d30d89b99205696c8c18bebaca1a188c4f0886a0542b01d3dcc01210271f262fee7b7aba93564d0ed468018f3ccca489ef9c87032a8c9db2dc820f7a0ba671400"
var ltc7vinsegwittx = "01000000000107ee3c061fe7e47aa415bc1081897b32e49af6598594a24c1d1b2bb7e00604a507000000001716001447c9e793a02bd8385b4111ef3e5f2a679232a705ffffffff4c730d6d90731449ff15cc930339f0328a81b8f3d0d8c9890d66db0e708a093501000000171600145d06566495653f42b8f849e1578958bb36696ba9ffffffff8966062f275134c0cf7e4f0f5b3582c7c6d4fb77ed96f45183c5cba5e9ea0c4f0000000017160014fd65dd5839468aa5b3ddce60e0d7d80b5ecdf4efffffffff7a8fe4b6203ad7358c2e8f5ee3f87a66c8cfe1dfe870403a3d49ed877b8c224f000000001716001495f4e71fbb90275814c9545429dbd5c8062992ddffffffff04f6d410c581a51345e0ee5de21cecaa609d934a1a6d9bbf2be7fcdb7303a052010000001716001434619178c272b4d5961fe60680b747b994ed0a66ffffffff0e7e6334d9ba1880dbb897102ed0101bf48bbb356749803b0ac7a4494ff5aa6a00000000171600145eef28cba09def35dd712be66cdf3ff2af5e2b6affffffff35ce7af9655cb6628167d62fc7fe3485bd2e81495577d529a3112421f2378791000000001716001439583f25323a5350db60f9dfe4cf45c074be44d6ffffffff0140787d01000000001976a9142516366665de65a2e13eed5695849dce6d9f023988ac02473044022077003d553bd6c50dcb0cb21983507bd47cbe3e2e2592d2104d7549ebe3840488022030d444bde565580ce780c607f193b4d86b5a770978e42b8f4c43d0c6558c268d012103d7dfa3bb835b90dd38ed749006884405a67d4fa0786e09759a7b33806eddf0a202483045022100d073e3710639d36c0ffe44fcd44445d948a91886dbc5ef5baee6413e396e08a2022072345d2af23bba25a642c409ef7b0ffe0fa46012d367f044448f3c893df6910f012102988800ebe1d9153c52a7b9c6f0db0c35825cf8292e83e184f64e7521dd22b03b02483045022100cb17e5827810f537b0a0f1f75ed422de9357f357911d699ec136bcc0fe7f3893022018ce673961f94c6aebb1e059a4e7bead1794efaee198a3927a6b964fbdb6f730012103ac701d3f53635a120f71ed5016fd722ab46633d1529504dd8ffa8938a1c95dc00247304402200d66f27be606cdbe1dcfb0aa2ed31ce3b88d9f731bcad9f3e7cbd45dcc474ce60220047f5957d01a8d5e1a33d50aee70ef334dcb1b2255d2dddf15aeeef81ccb36f20121026570d59029ae96cf76bf11b3746583f23b2a6c0807946db5e8b71e46f503d60902483045022100e9ddfc7e319987a612c88e43efcc6398a58727fda783a013348b66188c7dc786022047b782d99185dacaf7120a37c84070c20c1adc3b75eadb68f0f5f5f02fac0d42012102de96c2efc8199edd8eb8819deca092cfb581bb02752acd5ad50e60503880ead0024730440220175c48b5da5c35345851ad3bc2685b9b09d02b3abe15720f90346e39f3faf15802204f45f31c940365fa57e999ad50cad5970c07f877fa2e44b90e8ebb576481a18f0121027d27027ac5033ae07a8c88f53905cb5c93c63c831515ca5e132dfd763d8a528b024830450221008ca24611625ab8390fc7b9a5ad017acf5c7c602acc72eaf06cc213a72c6e4bb3022050baac8a1298dcab8e9a489b8a3cf3f84a6e525da288a7cfa82642887b90b3100121036bd0e251c8d657e4b5d9f7ecfb562b5df95c15fda445c5902aa64e28c13ce3a400000000"

func TestDigiByteTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(digibytetx)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}
	spew.Dump(tx)

	if tx.Version != 1 {
		t.Errorf("Invalid tx version. Expect %d, got %d", 1, tx.Version)
	}

	if tx.Hash.String() != "d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef" {
		t.Errorf("TX ID did not match for digibyte tx, expected %s, got %s", "d0e075c1e5c52854a5b5386e89bd6436c767a2570901d38537703baef3a313ef", tx.Hash.String())
	}

	if len(tx.Vin) != 2 {
		t.Errorf("Invalid input count. Expected %d, found %d", 2, len(tx.Vin))
	}

	if len(tx.Vout) != 2 {
		t.Errorf("Invalid output count. Expected %d, found %d", 2, len(tx.Vin))
	}

	if tx.Locktime != 6257229 {
		t.Errorf("Incorrect lock time. Expected %d, found %d", 6257229, tx.Locktime)
	}
}

func TestDigiByteCoinbaseTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(digibytetxcoinbase)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}
	spew.Dump(tx)
}

func TestLitecoinSegwitTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(ltcsegwittx)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}
	spew.Dump(tx)
}

func Test7VinLitecoinSegwitTx(t *testing.T) {
	tx, err := NewTransactionFromHexString(ltc7vinsegwittx)
	if err != nil {
		t.Errorf("Could not parse tx hex; %s", err)
	}
	spew.Dump(tx)
}
