### To Do 

---

### Issue
https://github.com/TrustWallet/wallet-core/issues/3

### Course

##### Address derivation
`func Test_writeReadLedgerInfo(t *testing.T)` in types_test.go
HRP is defined in `address.go`

##### Testnet
* set up a local full node 
  https://medium.com/forbole/a-step-by-step-guide-to-join-cosmos-hub-testnet-e591a3d2cb41
* connect to testnet full node with gaiacli
  https://cosmos.network/docs/gaia/delegator-guide-cli.html#accessing-the-cosmos-hub-network
* a working testnet full node for gaia-13002
  http://bity-testnet.cosmos-validators.com:26657
* a working testnet full node for gaia-13003
  node = "https://gaia.validator.network:443"
* a working testnet and faucet 
  https://hubble.figment.network/chains/gaia-13002
 
key06 address cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02
key09 address cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573

* gaiacli create a tx, --generate-only for offline tx
```
gaiacli tx send cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573 1muon --fees=1photino --from=key06 --chain-id=gaia-13003 --generate-only > unsignedSendTx.json
```

* gaicacli submit a tx against mainnet (chain-id=cosmoshub-1)
* gaicacli submit a tx against testnet (chain-id=gaia-13002)
```
 gaiacli tx send cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573 1muon --fees=200muon --from=key06 --chain-id=gaia-13003
```

* gaiacli sign a tx
```
gaiacli tx sign --from=key06 --chain-id=gaia-13003 unsignedSendTx.json > signedSendTx.json
```

* gaiacli query an account
```
gaiacli query account cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02
gaiacli query account cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573
```

##### Accounts
`gaiacli query account <your_account_address>`

##### TX signing
* `func TestTxSigningForTW(t *testing.T)` in stdtx_test.go for tx signing
* `func TestMsgSendGetSignBytes(t *testing.T)` for MsgSend encoding

* `sig, err := priv1.Sign(signBytes)` in app4_test.go
* https://media.readthedocs.org/pdf/cosmos-sdk/docs-refactor/cosmos-sdk.pdf

* mainnet node 
```
http://rpc.hub.certus.one:26657
http://cosmos-node.sparkpool.com:26657
```

* gaiacli config in file `/Users/nanjiang/.gaiacli/config/config.toml`
```
chain-id = "cosmoshub-1"
node = "cosmos-node.sparkpool.com:26657"
```

* Cosmos mainnet explorer, validator list, wallet
`https://lunie.io/#/staking/validators/`

* mainnet validators
`https://ztake.org/`
`https://support.sparkpool.com/hc/en-us/articles/360018949214-Cosmos-Delegator-Guide-CLI-`

* testnet explorer, validator list, wallet
`testnet.lunie.io`

* mainnet validator list
`https://forum.cosmos.network/t/validator-candidates/127`

* Gaiacli
** Send Tx using
   `https://cosmos.network/docs/gaia/delegator-guide-cli.html#sending-transactions`

##### Staking
* query account on testnet
  `gaiacli query account cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02 --chain-id=gaia-13003`
  `gaiacli query account cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573`

* get validator list on testnet
  `gaiacli query staking validators --chain-id=gaia-13003`

* delegate to a validator on testnet 
```
gaiacli tx staking delegate cosmosvaloper1zkupr83hrzkn3up5elktzcq3tuft8nxsmwdqgp 10muon --from key06 --chain-id=gaia-13003 --node https://gaia.validator.network:443 --gas auto --gas-adjustment 1.2 --gas-prices 0.01muon
```

##### Build
`dep ensure -v`
`make tools install`
`make distclean tools vendor-deps install`

##### Config
`gaiacli config chain-id cosmoshub-1`

/Users/nanjiang/.gaiacli/config/config.toml

```
chain-id = "gaia-13003"
node = "https://gaia.validator.network:443"
```

##### Keys
key password is `password`

jnj:cosmos-sdk nanjiang$ `gaiacli keys add key06`
override the existing name key06 [Y/n]: Y
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
kb.CreateAccount(name, mnemonic, bip39Passphrase, encryptPassword, account, index)
0
0
*** persistDerivedKey()
fullHdPath
44'/118'/0'/0/0
masterPriv
[64 162 28 243 173 217 42 162 183 45 125 240 180 252 21 235 139 134 33 24 150 171 48 109 56 209 205 149 125 20 227 46]
40a21cf3add92aa2b72d7df0b4fc15eb8b86211896ab306d38d1cd957d14e32e
derivedPriv
[128 232 30 162 105 230 106 10 5 177 18 54 223 121 25 251 127 190 237 186 135 69 45 102 116 137 215 64 58 2 240 5]
80e81ea269e66a0a05b11236df7919fb7fbeedba87452d667489d7403a02f005

*** Bech32KeyOutput()
info.GetPubKey() 
PubKeySecp256k1{0257286EC3F37D33557BBBAA000B27744AC9023AA9967CAE75A181D1FF91FA9DC5}
info.GetAddress()
BC2DA90C84049370D1B7C528BC164BC588833F21
NAME:   TYPE:   ADDRESS:                    PUBKEY:
key06   local   cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02   cosmospub1addwnpepqftjsmkr7d7nx4tmhw4qqze8w39vjq364xt8etn45xqarlu3l2wu2n7pgrq

**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

attract term foster morning tail foam excite copper disease measure cheese camera rug enroll cause flip sword waste try local purchase between idea thank

TX history
https://hubble.figment.network/chains/gaia-13002/transactions/DA9704152B75F4899EBEA590245C9AB60F65308672C881EC42899F473064981E

---

jnj:cosmos-sdk nanjiang$ gaiacli keys add key09
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
kb.CreateAccount(name, mnemonic, bip39Passphrase, encryptPassword, account, index)
0
0
*** persistDerivedKey()
fullHdPath
44'/118'/0'/0/0
masterPriv
[165 122 175 162 138 136 243 255 227 152 173 43 161 208 39 106 52 171 199 96 92 182 184 194 253 180 60 44 242 161 130 249]
a57aafa28a88f3ffe398ad2ba1d0276a34abc7605cb6b8c2fdb43c2cf2a182f9
derivedPriv
[18 78 105 194 194 218 204 118 96 15 128 106 49 51 60 16 11 65 177 212 55 78 153 245 57 228 17 86 194 121 44 12]
124e69c2c2dacc76600f806a31333c100b41b1d4374e99f539e41156c2792c0c

*** Bech32KeyOutput()
info.GetPubKey() 
PubKeySecp256k1{03F7707B1E983FE6231D46E586DCCAB76501302481C9656859C830A42A2C1BD414}
info.GetAddress()
12E8FE8B81ECC1F4F774EA6EC8DF267138B9F2D9
NAME:   TYPE:   ADDRESS:                    PUBKEY:
key09   local   cosmos1zt50azupanqlfam5afhv3hexwyutnukeh4c573   cosmospub1addwnpepq0mhq7c7nql7vgcagmjcdhx2kajszvpys8yk26zeeqc2g23vr02pglhlrpp

**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

extend artwork victory tourist into claim legal speed noodle try knock dutch quantum grief season aspect taste suffer search maple deny reopen inform audit

---

#### References

##### C++
* print bytes from string
```
inline unsigned int to_uint(char ch)
{
    // EDIT: multi-cast fix as per David Hammen's comment
    return static_cast<unsigned int>(static_cast<unsigned char>(ch));
}

inline void printStringInBytes(std::string str) {
    std::cout << "[";
    for (char ch : str) {
        std::cout << to_uint(ch) << ' '; 
    }
    std::cout << "]" << std::endl;
}
```

##### Git
```
git fetch upstream
git pull upstream master
```

```
git merge --squash branch-name
git commit -m "???"
```

```
git rebase -i HEAD~x
```