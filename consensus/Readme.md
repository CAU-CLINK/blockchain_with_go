# Consensus

## POW
1. import Version
2. import PrevblockHash
3. import MerkleRootHash
4. import Timestamp
5. set bits   
if current blockheight is multiple of 2016, readjust difficulty  
else, import last block's difficulty
6. set Nonce to 0
7. set current block's data
8. ++Nonce till it satisfies difficulty